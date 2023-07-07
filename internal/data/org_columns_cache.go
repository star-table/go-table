package data

import (
	"context"

	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-table/internal/data/po"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/star-table/go-table/internal/data/consts/cache"
)

type orgColumnsCache struct {
	redisBase
	data *Data
	log  *log.Helper
}

// NewOrgColumnsCache .
func NewOrgColumnsCache(data *Data, logger log.Logger) *orgColumnsCache {
	return &orgColumnsCache{
		redisBase: redisBase{
			data: data,
		},
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o *orgColumnsCache) getColumns(ctx context.Context, orgId int64, columnIds []string) ([]*po.OrgColumn, error) {
	key := cache.GetOrgColumnsKey(orgId)
	result := make([]*po.OrgColumn, 0, 20)

	err := o.getHashValues(ctx, key, columnIds, func() interface{} {
		temp := &po.OrgColumn{}
		result = append(result, temp)
		return temp
	})

	return result, errors.WithStack(err)
}

func (o *orgColumnsCache) deleteColumn(ctx context.Context, orgId int64, columnId string) error {
	return o.data.redisCli.HDel(ctx, cache.GetOrgColumnsKey(orgId), columnId).Err()
}

func (o *orgColumnsCache) deleteColumns(ctx context.Context, orgId int64) error {
	return o.data.redisCli.Del(ctx, cache.GetOrgColumnsKey(orgId)).Err()
}

func (o *orgColumnsCache) setColumns(ctx context.Context, orgId int64, columns []*po.OrgColumn) error {
	key := cache.GetOrgColumnsKey(orgId)
	values := make(map[interface{}]interface{}, len(columns))
	for _, c := range columns {
		values[c.ColumnId] = c
	}

	return o.setHashValues(ctx, key, values, false)
}
