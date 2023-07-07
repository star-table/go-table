package data

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"sort"

	"github.com/spf13/cast"

	"github.com/star-table/go-common/utils/unsafe"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/star-table/go-table/internal/data/consts/cache"
	"github.com/star-table/go-table/internal/data/po"
)

type tableCache struct {
	redisBase
	data *Data
	log  *log.Helper
}

// NewTableCache .
func NewTableCache(data *Data, logger log.Logger) *tableCache {
	return &tableCache{
		redisBase: redisBase{
			data: data,
		},
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t *tableCache) getTables(ctx context.Context, appId int64) ([]*po.Table, error) {
	key := cache.GetAppTablesKey(appId)
	result := make([]*po.Table, 0, 20)

	err := t.getHashValues(ctx, key, nil, func() interface{} {
		temp := &po.Table{}
		result = append(result, temp)
		return temp
	})
	result = t.sortTables(result)

	return result, err
}

// 因为出来是map，根据id拍一下序
func (t *tableCache) sortTables(tables []*po.Table) []*po.Table {
	sort.Slice(tables, func(i, j int) bool {
		return tables[i].ID < tables[j].ID
	})

	return tables
}

func (t *tableCache) getTablesByApps(ctx context.Context, appIds []int64) (map[int64][]*po.Table, []int64, error) {
	result := make(map[int64][]*po.Table, len(appIds))
	keys := make([]string, 0, len(appIds))
	for _, id := range appIds {
		keys = append(keys, cache.GetAppTablesKey(id))
	}
	_, notFounds, err := t.getHashValuesByKeys(ctx, keys, nil, appIds, func(k int64) interface{} {
		temp := &po.Table{}
		result[k] = append(result[k], temp)
		return temp
	})

	for i := range result {
		result[i] = t.sortTables(result[i])
	}

	return result, notFounds, errors.WithStack(err)
}

func (t *tableCache) setSummeryTable(ctx context.Context, orgId int64, tb *po.Table) error {
	key := cache.GetSummeryTableKey(orgId)
	return t.setObject(ctx, key, tb)
}

func (t *tableCache) getSummeryTable(ctx context.Context, orgId int64) (*po.Table, error) {
	key := cache.GetSummeryTableKey(orgId)
	tb := &po.Table{}
	err := t.getObject(ctx, key, tb)

	return tb, err
}

func (t *tableCache) setTables(ctx context.Context, appId int64, tables []*po.Table) error {
	key := cache.GetAppTablesKey(appId)
	values := make(map[interface{}]interface{}, len(tables))
	for _, tb := range tables {
		values[cast.ToString(tb.ID)] = tb
	}
	return t.setHashValues(ctx, key, values, false)
}

func (t *tableCache) setTablesBatch(ctx context.Context, tablesMap map[int64][]*po.Table) error {
	hashValuesMap := make(map[string]map[interface{}]interface{}, len(tablesMap))
	for appId, tables := range tablesMap {
		key := cache.GetAppTablesKey(appId)
		hashValues := make(map[interface{}]interface{}, len(tables))
		for _, tb := range tables {
			hashValues[cast.ToString(tb.ID)] = tb
		}
		hashValuesMap[key] = hashValues
	}

	return t.setHashValuesPipelined(ctx, hashValuesMap)
}

func (t *tableCache) setTable(ctx context.Context, appId int64, table *po.Table) error {
	key := cache.GetAppTablesKey(appId)
	return t.setHashValue(ctx, key, cast.ToString(table.ID), table, true)
}

func (t *tableCache) deleteTable(ctx context.Context, appId, tableId int64) error {
	return t.data.redisCli.HDel(ctx, cache.GetAppTablesKey(appId), cast.ToString(tableId)).Err()
}

func (t *tableCache) deleteTables(ctx context.Context, appId int64) error {
	return t.data.redisCli.Del(ctx, cache.GetAppTablesKey(appId)).Err()
}

// getColumns 取所有列
func (t *tableCache) getColumns(ctx context.Context, tableId int64, columnIds []string) ([]*po.TableColumn, error) {
	key := cache.GetTableSchemasKey(tableId)
	result := make([]*po.TableColumn, 0, 20)

	err := t.getHashValues(ctx, key, columnIds, func() interface{} {
		temp := &po.TableColumn{}
		result = append(result, temp)
		return temp
	})

	return result, errors.WithStack(err)
}

func (t *tableCache) getColumnsByTableIds(ctx context.Context, tableIds []int64, columnIds []string) (
	map[int64][]*po.TableColumn, []int64, error) {

	result := make(map[int64][]*po.TableColumn, len(tableIds))
	keys := make([]string, 0, len(tableIds))
	for _, id := range tableIds {
		keys = append(keys, cache.GetTableSchemasKey(id))
	}
	_, notFounds, err := t.getHashValuesByKeys(ctx, keys, columnIds, tableIds, func(k int64) interface{} {
		temp := &po.TableColumn{}
		result[k] = append(result[k], temp)
		return temp
	})

	return result, notFounds, errors.WithStack(err)
}

// getColumnsDescriptionByTableIds 获取字段描述
func (t *tableCache) getColumnsDescriptionByTableIds(ctx context.Context, tableIds []int64, columnIds []string) (
	map[int64]map[string]string, []int64, error) {

	keys := make([]string, 0, len(tableIds))
	for _, id := range tableIds {
		keys = append(keys, cache.GetColumnDescriptionKey(id))
	}
	hashValuesMap, notFounds, err := t.getHashValuesByKeys(ctx, keys, columnIds, tableIds, nil)

	return hashValuesMap, notFounds, errors.WithStack(err)
}

// getColumn 取一列
func (t *tableCache) getColumn(ctx context.Context, tableId int64, columnId string) (*po.TableColumn, error) {
	s, err := t.data.redisCli.HGet(ctx, cache.GetTableSchemasKey(tableId), columnId).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, err
		}
		return nil, errors.WithStack(err)
	}

	tc := &po.TableColumn{}
	err = json.Unmarshal(unsafe.StringBytes(s), tc)
	if err != nil {
		return nil, errors.Wrapf(err, "[getColumns] unmarshal str:%v error:%v", s, err)
	}

	return tc, nil
}

// setColumns 存所有列的配置信息
func (t *tableCache) setColumns(ctx context.Context, tableId int64, columns []*po.TableColumn) error {
	key := cache.GetTableSchemasKey(tableId)
	values := make(map[interface{}]interface{}, len(columns))
	for _, c := range columns {
		values[c.ColumnId] = c
	}

	return t.setHashValues(ctx, key, values, false)
}

func (t *tableCache) setColumnsBatch(ctx context.Context, columnsMap map[int64][]*po.TableColumn) error {
	hashValuesMap := make(map[string]map[interface{}]interface{}, len(columnsMap))
	for tableId, columns := range columnsMap {
		key := cache.GetTableSchemasKey(tableId)
		hashValues := make(map[interface{}]interface{}, len(columns))
		for _, c := range columns {
			hashValues[c.ColumnId] = c
		}
		hashValuesMap[key] = hashValues
	}

	return t.setHashValuesPipelined(ctx, hashValuesMap)
}

func (t *tableCache) setColumnsDescriptionBatch(ctx context.Context, descMap map[int64]map[string]string) error {
	hashValuesMap := make(map[string]map[interface{}]interface{}, len(descMap))
	for tableId, m := range descMap {
		key := cache.GetColumnDescriptionKey(tableId)
		hashValues := make(map[interface{}]interface{}, len(m))
		for s, s2 := range m {
			hashValues[s] = s2
		}
		hashValuesMap[key] = hashValues
	}

	return t.setHashValuesPipelined(ctx, hashValuesMap)
}

func (t *tableCache) setColumnDescription(ctx context.Context, tableId int64, columnId, description string) error {
	key := cache.GetColumnDescriptionKey(tableId)
	return t.data.redisCli.HSet(ctx, key, columnId, description).Err()
}

func (t *tableCache) deleteTableDescription(ctx context.Context, tableId int64) error {
	key := cache.GetColumnDescriptionKey(tableId)
	return t.data.redisCli.Del(ctx, key).Err()
}

// deleteColumns 删除所有列数据
func (t *tableCache) deleteColumns(ctx context.Context, tableId int64) error {
	return t.data.redisCli.Del(ctx, cache.GetTableSchemasKey(tableId)).Err()
}

// deleteColumn 删除一列
func (t *tableCache) deleteColumn(ctx context.Context, tableId int64, columnId string) error {
	return t.data.redisCli.HDel(ctx, cache.GetTableSchemasKey(tableId), columnId).Err()
}

// setColumn 存一列的配置信息
func (t *tableCache) setColumn(ctx context.Context, c *po.TableColumn) error {
	key := cache.GetTableSchemasKey(c.TableId)
	return t.setHashValue(ctx, key, c.ColumnId, c, true)
}
