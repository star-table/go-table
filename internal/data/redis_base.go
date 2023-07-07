package data

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-common/utils/unsafe"
	"github.com/star-table/go-table/internal/data/consts/cache"
)

// 一个闭包函数，用于生成新的类，并添加到返回值里面
type addInterface func() interface{}
type addInterfaceByKey func(k int64) interface{}

type redisBase struct {
	data *Data
}

func (t *redisBase) setObject(ctx context.Context, key string, obj interface{}) error {
	bts, err := json.Marshal(obj)
	if err != nil {
		return errors.WithStack(err)
	}
	return t.data.redisCli.Set(ctx, key, unsafe.BytesString(bts), cache.DefaultDuration).Err()
}

func (t *redisBase) getObject(ctx context.Context, key string, obj interface{}) error {
	s, err := t.data.redisCli.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return errors.WithStack(json.Unmarshal(unsafe.StringBytes(s), obj))
}

// getHashValuesByKeys 目前只针对场景函数，没有特别的抽象化
func (t *redisBase) getHashValuesByKeys(ctx context.Context, keys []string, hKeys []string, ids []int64,
	addFunc addInterfaceByKey) (map[int64]map[string]string, []int64, error) {
	var (
		allCmds       []*redis.StringStringMapCmd
		partCmds      []*redis.SliceCmd
		notFounds     []int64
		hashValuesMap map[int64]map[string]string
	)
	if len(hKeys) > 0 {
		partCmds = make([]*redis.SliceCmd, 0, len(keys))
	} else {
		allCmds = make([]*redis.StringStringMapCmd, 0, len(keys))
	}

	_, err := t.data.redisCli.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, key := range keys {
			if len(hKeys) > 0 {
				cmd := pipe.HMGet(ctx, key, hKeys...)
				partCmds = append(partCmds, cmd)
			} else {
				cmd := pipe.HGetAll(ctx, key)
				allCmds = append(allCmds, cmd)
			}
		}
		return nil
	})
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	if len(hKeys) > 0 {
		hashValuesMap, notFounds, err = t.getValuesBySliceCmd(ctx, partCmds, keys, hKeys, ids, addFunc)
	} else {
		hashValuesMap, notFounds, err = t.getValuesByMapCmd(allCmds, ids, addFunc)
	}

	return hashValuesMap, notFounds, err
}

func (t *redisBase) getValuesBySliceCmd(ctx context.Context, partCmds []*redis.SliceCmd, keys, hKeys []string,
	ids []int64, addFunc addInterfaceByKey) (map[int64]map[string]string, []int64, error) {

	notFounds := make([]int64, 0, 1)
	result := make(map[int64]map[string]string, len(keys))
	for i, cmd := range partCmds {
		list, err := cmd.Result()
		if err != nil && err != redis.Nil {
			return nil, nil, errors.WithStack(err)
		}
		isAllNil := true
		for _, s := range list {
			if s != nil {
				isAllNil = false
			}
		}
		if isAllNil {
			isExist, err := t.data.redisCli.Exists(ctx, keys[i]).Result()
			if err != nil {
				return nil, nil, errors.WithStack(err)
			}
			if isExist == 0 {
				notFounds = append(notFounds, ids[i])
			}
		} else {
			tempM := make(map[string]string, len(list))
			for k, s := range list {
				if s != nil {
					str := cast.ToString(s)
					tempM[hKeys[k]] = str
					// 只有需要解析的时候才进入，有些其实只是收集字符串，并没有结构化
					if addFunc != nil {
						temp := addFunc(ids[i])
						err = json.Unmarshal(unsafe.StringBytes(str), temp)
						if err != nil {
							return nil, nil, errors.Wrapf(err, "[getHashValuesByKeys] unmarshal str:%v error:%v", s, err)
						}
					}
				}
			}
			result[ids[i]] = tempM
		}
	}

	return result, notFounds, nil
}

func (t *redisBase) getValuesByMapCmd(allCmds []*redis.StringStringMapCmd, ids []int64, addFunc addInterfaceByKey) (
	map[int64]map[string]string, []int64, error) {

	notFounds := make([]int64, 0, 1)
	result := make(map[int64]map[string]string, len(ids))
	for i, cmd := range allCmds {
		m, err := cmd.Result()
		if err != nil && err != redis.Nil {
			return nil, nil, errors.WithStack(err)
		}

		if len(m) == 0 {
			notFounds = append(notFounds, ids[i])
		} else {
			result[ids[i]] = m
			if addFunc != nil {
				for _, s := range m {
					temp := addFunc(ids[i])
					err = json.Unmarshal(unsafe.StringBytes(s), temp)
					if err != nil {
						return nil, nil, errors.Wrapf(err, "[getHashValuesByKeys] unmarshal str:%v error:%v", s, err)
					}
				}
			}
		}
	}

	return result, notFounds, nil
}

func (t *redisBase) getHashValues(ctx context.Context, key string, hKeys []string, nf addInterface) error {
	var (
		m   map[string]string
		err error
	)
	if len(hKeys) == 0 {
		m, err = t.data.redisCli.HGetAll(ctx, key).Result()
	} else {
		var list []interface{}
		list, err = t.data.redisCli.HMGet(ctx, key, hKeys...).Result()
		if err == nil {
			m = make(map[string]string, len(list))
			for i, id := range hKeys {
				str := cast.ToString(list[i])
				if str != "" {
					m[id] = cast.ToString(list[i])
				}
			}
		}
	}

	if err != nil && err != redis.Nil {
		return errors.WithStack(err)
	}
	if len(m) == 0 {
		return redis.Nil
	}

	for _, s := range m {
		temp := nf()
		err = json.Unmarshal(unsafe.StringBytes(s), temp)
		if err != nil {
			return errors.Wrapf(err, "[getHashValues] unmarshal str:%v error:%v", s, err)
		}
	}

	return nil
}

func (t *redisBase) setHashValue(ctx context.Context, key, hKey string, v interface{}, justSetByExist bool) error {
	// 如果设置了只有存在才设置，检查下是否存在，不存在才设置
	if justSetByExist {
		count, err := t.data.redisCli.HLen(ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				return nil
			}
			return errors.WithStack(err)
		}
		if count == 0 {
			return nil
		}
	}

	bts, _ := json.Marshal(v)
	err := t.data.redisCli.HSet(ctx, key, hKey, string(bts)).Err()
	if err != nil {
		return errors.WithStack(err)
	}

	err = t.data.redisCli.Expire(ctx, key, cache.DefaultDuration).Err()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (t *redisBase) setHashValues(ctx context.Context, key string, hashValues map[interface{}]interface{}, justSetByExist bool) error {
	// 如果设置了只有存在才设置，检查下是否存在，不存在才设置
	if justSetByExist {
		count, err := t.data.redisCli.HLen(ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				return nil
			}
			return errors.WithStack(err)
		}
		if count == 0 {
			return nil
		}
	}

	values := make([]interface{}, 0, len(hashValues)*2)
	for k, c := range hashValues {
		if s, ok := c.(string); ok {
			values = append(values, k, s)
		} else {
			bts, _ := json.Marshal(c)
			values = append(values, k, unsafe.BytesString(bts))
		}
	}
	err := t.data.redisCli.HMSet(ctx, key, values...).Err()
	if err != nil {
		return errors.WithStack(err)
	}

	err = t.data.redisCli.Expire(ctx, key, cache.DefaultDuration).Err()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// setHashValuesPipelined 批量设置hash
func (t *redisBase) setHashValuesPipelined(ctx context.Context, hashValuesMap map[string]map[interface{}]interface{}) error {
	_, err := t.data.redisCli.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for key, hashValues := range hashValuesMap {
			if len(hashValues) > 0 {
				values := make([]interface{}, 0, len(hashValues)*2)
				for k, c := range hashValues {
					if s, ok := c.(string); ok {
						values = append(values, k, s)
					} else {
						bts, _ := json.Marshal(c)
						values = append(values, k, unsafe.BytesString(bts))
					}
				}

				pipe.HMSet(ctx, key, values...)
				pipe.Expire(ctx, key, cache.DefaultDuration)
			}
		}
		return nil
	})

	return err
}
