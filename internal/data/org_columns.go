package data

import (
	"context"
	"sort"

	"github.com/star-table/go-common/pkg/middleware/meta"

	"gorm.io/gorm/clause"

	"github.com/star-table/go-table/internal/biz/bo/covert"

	"github.com/star-table/go-table/internal/biz/bo"

	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-table/internal/biz"
	"github.com/star-table/go-table/internal/data/po"
)

type orgColumnsRepo struct {
	data            *Data
	orgColumnsCache *orgColumnsCache
	log             *log.Helper
}

// NewOrgColumnsRepo .
func NewOrgColumnsRepo(data *Data, orgColumnsCache *orgColumnsCache, logger log.Logger) biz.OrgColumnsRepo {
	return &orgColumnsRepo{
		data:            data,
		orgColumnsCache: orgColumnsCache,
		log:             log.NewHelper(logger),
	}
}

func (o *orgColumnsRepo) GetColumns(ctx context.Context, orgId int64, columnIds []string) ([]*bo.OrgColumn, error) {
	columns, err := o.orgColumnsCache.getColumns(ctx, orgId, columnIds)
	if err != nil && errors.Cause(err) != redis.Nil {
		return nil, err
	}

	if len(columns) == 0 {
		columns = make([]*po.OrgColumn, 0, 15)
		err = o.data.mysqlLcGo.Where("org_id = ?", orgId).Find(&columns).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if len(columns) == 0 {
			return []*bo.OrgColumn{}, nil
		}

		_ = o.orgColumnsCache.setColumns(ctx, orgId, columns)

		columns = o.getLimitColumns(columns, columnIds)
	}

	sort.Slice(columns, func(i, j int) bool {
		return columns[i].ID > columns[j].ID
	})

	return covert.OrgColumnCovert.ToColumns(columns)
}

func (o *orgColumnsRepo) getLimitColumns(columns []*po.OrgColumn, columnIds []string) []*po.OrgColumn {
	// 如果只拿一部分，由于还是要缓存所有的列，所以拿出来后再过滤
	if len(columnIds) > 0 {
		newColumns := make([]*po.OrgColumn, 0, len(columnIds))
		for _, id := range columnIds {
			for i, column := range columns {
				if column.ColumnId == id {
					newColumns = append(newColumns, columns[i])
					break
				}
			}
		}
		return newColumns
	}

	return columns
}

// CreateColumns 创建列
func (o *orgColumnsRepo) CreateColumns(ctx context.Context, orgId int64, columns []*po.OrgColumn) error {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	for _, column := range columns {
		column.ID = o.data.snowFlake.Generate().Int64()
		column.Creator = ch.UserId
	}

	err := o.data.mysqlLcGo.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := o.data.mysqlLcGo.WithContext(ctx).Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{"column_type", "schema"}),
		}).Create(columns).Error
		if err != nil {
			return err
		}

		return o.orgColumnsCache.deleteColumns(ctx, orgId)
	})

	return errors.WithStack(err)
}

func (o *orgColumnsRepo) DeleteColumn(ctx context.Context, orgId int64, columnId string) error {
	err := o.data.mysqlLcGo.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		orgColumn := &po.OrgColumn{OrgId: orgId, ColumnId: columnId}
		err := tx.Where(orgColumn).Delete(orgColumn).Error
		if err != nil {
			return err
		}

		return o.orgColumnsCache.deleteColumns(ctx, orgId)
	})

	return errors.WithStack(err)
}
