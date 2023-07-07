package biz

import (
	"context"

	"github.com/star-table/go-table/internal/data/facade/vo/datacentervo"
)

type DatacenterRepo interface {
	UpdateBatchRowByRowIds(ctx context.Context, orgId, tableId int64, values map[int64]interface{}) error
	UpdateTableRowByRowId(ctx context.Context, orgId, tableId, rowId int64, values interface{}) error
	GetTableRowData(ctx context.Context, orgId, tableId int64, condition *datacentervo.LessCondsData, limit, offset int, columns []string, orders []datacentervo.Order) ([]map[string]interface{}, error)
	GetTableRowDataByRowId(ctx context.Context, orgId, tableId, rowId int64) (map[string]interface{}, error)
	GetTableRowListByRowIds(ctx context.Context, orgId, tableId int64, rowIds []int64) ([]map[string]interface{}, error)
	WrapperTableName(orgId, tableId int64) string
	JsonColumn(str string) string
	CreateTables(ctx context.Context, tables []*datacentervo.CreateTableReq) (*datacentervo.CreateTableResp, error)
	DatacenterExecutor(ctx context.Context, dsId, dbId int64, req *datacentervo.Executor) (*datacentervo.ExecutorResp, error)
	DatacenterExecutors(ctx context.Context, dsId, dbId int64, req []*datacentervo.Executor) (*datacentervo.ExecutorsResp, error)
}
