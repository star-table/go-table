package data

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/star-table/go-common/pkg/middleware/meta"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-table/internal/biz"
	"github.com/star-table/go-table/internal/data/consts"
	"github.com/star-table/go-table/internal/data/po"
	commonPb "github.com/star-table/interface/golang/common/v1"
	"gorm.io/gorm"
)

type tableRepo struct {
	data       *Data
	tableCache *tableCache
	log        *log.Helper
}

// NewTableRepo .
func NewTableRepo(data *Data, tableCache *tableCache, logger log.Logger) biz.TableRepo {
	return &tableRepo{
		data:       data,
		tableCache: tableCache,
		log:        log.NewHelper(logger),
	}
}

// StartTransactionLcGo 开启事务操作
func (t *tableRepo) StartTransactionLcGo(ctx context.Context, fc func(tx *gorm.DB) error) error {
	return t.data.mysqlLcGo.WithContext(ctx).Transaction(fc)
}

// GetSummeryTableId 获取汇总表的tableId
func (t *tableRepo) GetSummeryTableId(ctx context.Context, orgId int64) (int64, error) {
	tb, err := t.tableCache.getSummeryTable(ctx, orgId)
	if err != nil && err != redis.Nil {
		return 0, err
	}
	if tb != nil && tb.ID > 0 {
		return tb.ID, nil
	}

	tb = &po.Table{}
	err = t.data.mysqlLcGo.WithContext(ctx).Select(po.TableSimpleSelect).Where(&po.Table{OrgId: orgId, SummeryFlag: consts.FlagYes}).Take(&tb).Error
	if err != nil {
		return 0, errors.Ignore(commonPb.ErrorResourceNotExist(fmt.Sprintf("can not find summery table with orgId:%v", orgId)))
	}

	_ = t.tableCache.setSummeryTable(ctx, orgId, tb)

	return tb.ID, nil
}

// GetTables 获取表列表
func (t *tableRepo) GetTables(ctx context.Context, appId int64, tx ...*gorm.DB) ([]*po.Table, error) {
	db := t.data.mysqlLcGo
	if len(tx) >= 1 {
		db = tx[0]
	}

	tables, err := t.tableCache.getTables(ctx, appId)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if len(tables) > 0 {
		return tables, nil
	}

	tables = make([]*po.Table, 0, 5)
	err = db.WithContext(ctx).Select(po.TableSimpleSelect).Where(&po.Table{AppId: appId, DelFlag: consts.DeleteFlagNotDel}).Order("id asc").Find(&tables).Error

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(tables) == 0 {
		return nil, errors.Ignore(commonPb.ErrorResourceNotExist(fmt.Sprintf("can not find tables with appId:%v", appId)))
	}

	err = t.tableCache.setTables(ctx, appId, tables)

	return tables, errors.WithStack(err)
}

func (t *tableRepo) GetTable(ctx context.Context, tableId int64) (*po.Table, error) {
	tb := &po.Table{}
	err := t.data.mysqlLcGo.WithContext(ctx).Select(po.TableSimpleSelect).Where(&po.Table{Base: po.Base{ID: tableId}, DelFlag: consts.DeleteFlagNotDel}).Order("id asc").Take(tb).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Ignore(commonPb.ErrorResourceNotExist("table not exists"))
	}

	return tb, errors.WithStack(err)
}

// GetTablesByApps 获取表列表，根据app列表
func (t *tableRepo) GetTablesByApps(ctx context.Context, appIds []int64) (map[int64][]*po.Table, error) {
	tablesMap, notInCacheIds, err := t.tableCache.getTablesByApps(ctx, appIds)
	if err != nil {
		return nil, err
	}

	if len(notInCacheIds) > 0 {
		tables := make([]*po.Table, 0, 5)
		err = t.data.mysqlLcGo.WithContext(ctx).Select(po.TableSimpleSelect).Where("app_id in(?) and del_flag = ?", notInCacheIds, consts.DeleteFlagNotDel).
			Order("id asc").Find(&tables).Error

		if err != nil {
			return nil, errors.WithStack(err)
		}

		needSetCacheAppIdsMap := make(map[int64][]*po.Table, 2)
		for _, tb := range tables {
			tablesMap[tb.AppId] = append(tablesMap[tb.AppId], tb)
			needSetCacheAppIdsMap[tb.AppId] = append(needSetCacheAppIdsMap[tb.AppId], tb)
		}

		_ = t.tableCache.setTablesBatch(ctx, needSetCacheAppIdsMap)
	}

	return tablesMap, nil
}

// GetTablesByOrgId 获取表列表
func (t *tableRepo) GetTablesByOrgId(ctx context.Context, orgId int64) ([]*po.Table, error) {
	tables := make([]*po.Table, 0, 5)
	err := t.data.mysqlLcGo.WithContext(ctx).Select("id,app_id,name").Where(&po.Table{OrgId: orgId, DelFlag: consts.DeleteFlagNotDel}).Find(&tables).Error

	return tables, errors.WithStack(err)
}

// CreateTable 创建表和表头数据
func (t *tableRepo) CreateTable(ctx context.Context, tb *po.Table, columns []*po.TableColumn, tx *gorm.DB) error {
	tb.ID = t.data.snowFlake.Generate().Int64()
	err := tx.Create(tb).Error
	if err != nil {
		return err
	}
	for _, tc := range columns {
		tc.ID = t.data.snowFlake.Generate().Int64()
		tc.TableId = tb.ID
	}

	if len(columns) > 0 {
		err = tx.Save(columns).Error
		if err != nil {
			return err
		}
	}

	return t.tableCache.deleteTables(ctx, tb.AppId)
}

// CopyTables 拷贝表以及列数据
func (t *tableRepo) CopyTables(ctx context.Context, srcAppId, destAppId int64,
	tableIds []int64, oldToNewPermission map[int64]int64) (map[int64]int64, error) {

	tables, columns, oldToNewId, err := t.getCopyTables(ctx, srcAppId, destAppId, tableIds, oldToNewPermission)
	if err != nil {
		return nil, err
	}

	err = t.data.mysqlLcGo.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(tables).Error
		if err != nil {
			return err
		}
		err = tx.CreateInBatches(columns, 100).Error
		if err != nil {
			return err
		}

		_ = t.tableCache.deleteTables(ctx, destAppId)

		return nil
	})

	return oldToNewId, errors.WithStack(err)
}

// getCopyTables 获取要copy的表以及列数据，这个时候不用缓存会比较安全，所以直接从数据库获取
func (t *tableRepo) getCopyTables(ctx context.Context, srcAppId, destAppId int64,
	tableIds []int64, oldToNewPermission map[int64]int64) ([]*po.Table, []*po.TableColumn, map[int64]int64, error) {

	ch := meta.GetCommonHeaderFromCtx(ctx)
	oldToNewId := make(map[int64]int64, 3)
	tables := make([]*po.Table, 0, len(tableIds))
	db := t.data.mysqlLcGo.WithContext(ctx)
	// 如果不传tableIds证明copy整个app下面的表
	if len(tableIds) == 0 {
		db = db.Where(&po.Table{AppId: srcAppId, DelFlag: int32(consts.DeleteFlagNotDel)})
	} else {
		db = db.Where("id in(?)", tableIds)
	}
	err := db.Order("id asc").Find(&tables).Error
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}

	if len(tableIds) == 0 {
		for _, tb := range tables {
			tableIds = append(tableIds, tb.ID)
		}
	}

	columns := make([]*po.TableColumn, 0, 15*len(tableIds))
	// 不复制表关联表头、引用表头，要不会有问题， 如果后期要做，如果是同一个app的表，可以复制表头，后面再实现
	err = t.data.mysqlLcGo.Where("table_id in(?)", tableIds).Find(&columns).Error
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}

	now := time.Now()
	for _, tb := range tables {
		newId := t.data.snowFlake.Generate().Int64()
		oldToNewId[tb.ID] = newId
		tb.ID = newId
		tb.OrgId = ch.OrgId
		tb.AppId = destAppId
		tb.UpdateAt = now
		tb.CreateAt = now
		tb.Creator = ch.UserId
		tb.Updater = 0
		tb.AutoScheduleFlag = 1
	}

	for _, column := range columns {
		column.ID = t.data.snowFlake.Generate().Int64()
		column.OrgId = ch.OrgId
		column.TableId = oldToNewId[column.TableId]
		column.CreateAt = now
		column.UpdateAt = now
		// 这个操作是从应用模板过来。。真是nb的操作，不怕替换错了吗
		column.Schema = strings.ReplaceAll(column.Schema, cast.ToString(srcAppId), cast.ToString(destAppId))
		for oldId, newId := range oldToNewId {
			column.Schema = strings.ReplaceAll(column.Schema, cast.ToString(oldId), cast.ToString(newId))
		}

		for o, n := range oldToNewPermission {
			column.Schema = strings.ReplaceAll(column.Schema, cast.ToString(o), cast.ToString(n))
		}
	}

	return tables, columns, oldToNewId, nil
}

// UpdateTableName 更新表名字
func (t *tableRepo) UpdateTableName(ctx context.Context, appId, tableId int64, name string, userId int64) error {
	tb := &po.Table{Base: po.Base{ID: tableId}}
	db := t.data.mysqlLcGo.WithContext(ctx).Where(tb).Updates(&po.Table{Name: name, Updater: userId})
	if db.Error != nil {
		return errors.WithStack(db.Error)
	}
	if db.RowsAffected == 0 {
		return nil
	}
	_ = t.tableCache.deleteTables(ctx, appId)
	return nil
}

// DeleteTable 删除表
func (t *tableRepo) DeleteTable(ctx context.Context, appId, tableId int64, tx *gorm.DB) error {
	tables, err := t.GetTables(ctx, appId, tx)
	if err != nil {
		return err
	}
	if len(tables) == 1 {
		return errors.Ignore(commonPb.ErrorCanNotDeleteLastTable("can not delete last table"))
	}

	err = tx.Where(&po.Table{Base: po.Base{ID: tableId}}).Updates(&po.Table{DelFlag: consts.DeleteFlagDel}).Error
	if err != nil {
		return errors.WithStack(err)
	}

	err = tx.Where(&po.TableColumn{TableId: tableId}).Updates(&po.TableColumn{DelFlag: consts.DeleteFlagDel}).Error
	if err != nil {
		return errors.WithStack(err)
	}

	err = t.tableCache.deleteTables(ctx, appId)
	if err != nil {
		return errors.WithStack(err)
	}

	_ = t.tableCache.deleteColumns(ctx, tableId)
	return nil
}

// SetAutoSchedule 设置自动排期开关
func (t *tableRepo) SetAutoSchedule(ctx context.Context, appId, tableId int64, autoScheduleFlag int32, userId int64) error {
	tb := &po.Table{Base: po.Base{ID: tableId}}
	db := t.data.mysqlLcGo.WithContext(ctx).Where(tb).Updates(&po.Table{AutoScheduleFlag: autoScheduleFlag, Updater: userId})
	if db.Error != nil {
		return errors.WithStack(db.Error)
	}
	if db.RowsAffected == 0 {
		return nil
	}
	_ = t.tableCache.deleteTables(ctx, appId)
	return nil
}
