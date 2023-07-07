package data

import (
	"context"
	"fmt"

	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-table/internal/data/consts"
	comomPb "github.com/star-table/interface/golang/common/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/star-table/go-table/internal/biz"
	"github.com/star-table/go-table/internal/conf"
	"github.com/star-table/go-table/internal/data/facade/vo/datacentervo"
)

type datacenterRepo struct {
	dc  *http.Client
	log *log.Helper
}

func NewDatacenterRepo(conf *conf.Data, r registry.Discovery, logger log.Logger) biz.DatacenterRepo {
	conn, err := getHttpConn(conf.Facade.DatacenterServer, r, logger)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &datacenterRepo{
		dc:  conn,
		log: log.NewHelper(logger),
	}
}

func (d *datacenterRepo) DatacenterQuery(ctx context.Context, req *datacentervo.QueryReq) (*datacentervo.QueryResp, error) {
	respVo := datacentervo.QueryResp{}
	path := fmt.Sprintf("/datacenter/inner/api/v1/%d/%d/query", consts.DataSource, consts.DataBase)
	err := d.dc.Invoke(ctx, "POST", path, req, &respVo)
	if err != nil {
		return nil, err
	}

	return &respVo, nil
}

func (d *datacenterRepo) DatacenterExecutor(ctx context.Context, dsId, dbId int64, req *datacentervo.Executor) (*datacentervo.ExecutorResp, error) {
	respVo := datacentervo.ExecutorResp{}
	path := fmt.Sprintf("/datacenter/inner/api/v1/%d/%d/execute", dsId, dbId)
	err := d.dc.Invoke(ctx, "POST", path, req, &respVo)
	if err != nil {
		return nil, err
	}

	if respVo.Failure() {
		return nil, errors.Wrapf(comomPb.ErrorResourceNotExist("execute error"), "[DatacenterExecutors] get app failed, req:%v, error:%v", req, respVo.Error())
	}

	return &respVo, err
}

func (d *datacenterRepo) DatacenterExecutors(ctx context.Context, dsId, dbId int64, req []*datacentervo.Executor) (*datacentervo.ExecutorsResp, error) {
	respVo := datacentervo.ExecutorsResp{}
	path := fmt.Sprintf("/datacenter/inner/api/v1/%d/%d/execute-batch", dsId, dbId)
	err := d.dc.Invoke(ctx, "POST", path, req, &respVo)
	if err != nil {
		return nil, err
	}
	if respVo.Failure() {
		return nil, errors.Wrapf(comomPb.ErrorResourceNotExist("execute error"), "[DatacenterExecutors] get app failed, req:%v, error:%v", req, respVo.Error())
	}

	return &respVo, err
}

//func (d *datacenterRepo) DatacenterBatchExecutor(ctx context.Context, req []*datacentervo.Executor) (*datacentervo.ExecutorBatchResp, error) {
//	respVo := &datacentervo.ExecutorBatchResp{}
//	path := fmt.Sprintf("/datacenter/inner/api/v1/%d/%d/execute-batch", consts.DataSource, consts.DataBase)
//	err := d.dc.Invoke(ctx, "POST", path, req, respVo)
//	if err != nil {
//		return nil, err
//	}
//	return respVo, nil
//}

func (d *datacenterRepo) GetTableRowData(ctx context.Context, orgId, tableId int64, condition *datacentervo.LessCondsData, limit, offset int, columns []string, orders []datacentervo.Order) ([]map[string]interface{}, error) {
	queryResp, err := d.DatacenterQuery(ctx, &datacentervo.QueryReq{
		From: []datacentervo.Table{
			{
				Type:   consts.TableType,
				Schema: d.WrapperTableName(orgId, tableId),
			},
		},
		Condition: *condition,
		Limit:     limit,
		Offset:    offset,
		Columns:   columns,
		Orders:    orders,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return queryResp.Data, nil
}

func (d *datacenterRepo) GetTableRowDataByRowId(ctx context.Context, orgId, tableId, rowId int64) (map[string]interface{}, error) {
	queryResp, err := d.DatacenterQuery(ctx, &datacentervo.QueryReq{
		From: []datacentervo.Table{
			{
				Type:   consts.TableType,
				Schema: d.WrapperTableName(orgId, tableId),
			},
		},
		Condition: datacentervo.LessCondsData{
			Type: datacentervo.ConditionAnd,
			Conds: []*datacentervo.LessCondsData{
				{
					Type:   datacentervo.ConditionEqual,
					Column: d.JsonColumn(consts.ColumnIdIssueId),
					Value:  rowId,
				},
				{
					Type:   datacentervo.ConditionEqual,
					Column: d.JsonColumn(consts.ColumnIdDelFlag),
					Value:  consts.DeleteFlagNotDel,
				},
			},
		},
		Limit: 1,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(queryResp.Data) < 1 {
		return nil, errors.Wrapf(comomPb.ErrorResourceNotExist("row not exist"), "get rowdata failed, orgId:%v, tableId:%v, rowId:%v, err:%v", orgId, tableId, rowId, queryResp.Error())
	}

	return queryResp.Data[0], nil
}

func (d *datacenterRepo) GetTableRowListByRowIds(ctx context.Context, orgId, tableId int64, rowIds []int64) ([]map[string]interface{}, error) {
	var valueRowIds []interface{}
	for _, rId := range rowIds {
		valueRowIds = append(valueRowIds, rId)
	}
	rowDataList, err := d.DatacenterQuery(ctx, &datacentervo.QueryReq{
		From: []datacentervo.Table{
			{
				Type:   consts.TableType,
				Schema: d.WrapperTableName(orgId, tableId),
			},
		},
		Condition: datacentervo.LessCondsData{
			Type: datacentervo.ConditionAnd,
			Conds: []*datacentervo.LessCondsData{
				{
					Type:   datacentervo.ConditionEqual,
					Column: d.JsonColumn(consts.ColumnIdDelFlag),
					Value:  consts.DeleteFlagNotDel,
				},
				{
					Type:   datacentervo.ConditionIn,
					Column: d.JsonColumn(consts.ColumnIdIssueId),
					Values: valueRowIds,
				},
			},
		},
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(rowDataList.Data) < 1 {
		return nil, errors.Wrapf(comomPb.ErrorResourceNotExist("row not exist"), "get rowdata failed, orgId:%v, tableId:%v, rowIds:%v, err:%v", orgId, tableId, rowIds, rowDataList.Error())
	}
	return rowDataList.Data, nil
}

func (d *datacenterRepo) UpdateBatchRowByRowIds(ctx context.Context, orgId, tableId int64, values map[int64]interface{}) error {
	executors := []*datacentervo.Executor{}
	table := &datacentervo.Table{
		Type:   consts.TableType,
		Schema: d.WrapperTableName(orgId, tableId),
	}

	for rowId, v := range values {
		executors = append(executors, &datacentervo.Executor{
			Type:  consts.ExecutorUpdate,
			Table: table,
			Condition: &datacentervo.LessCondsData{
				Type: datacentervo.ConditionAnd,
				Conds: []*datacentervo.LessCondsData{
					{
						Type:   datacentervo.ConditionEqual,
						Column: d.JsonColumn(consts.ColumnIdIssueId),
						Value:  rowId,
					},
					{
						Type:   datacentervo.ConditionEqual,
						Column: d.JsonColumn(consts.ColumnIdDelFlag),
						Value:  consts.DeleteFlagNotDel,
					},
				},
			},
			Sets: []datacentervo.Set{
				{
					Column: consts.ColumnData,
					Type:   consts.JsonbFields,
					Value:  v,
				},
			},
		})
	}

	_, err := d.DatacenterExecutors(ctx, consts.DataSource, consts.DataBase, executors)
	if err != nil {
		return err
	}
	return nil
}

func (d *datacenterRepo) UpdateTableRowByRowId(ctx context.Context, orgId, tableId, rowId int64, values interface{}) error {
	lessCond := datacentervo.LessCondsData{
		Type: datacentervo.ConditionAnd,
		Conds: []*datacentervo.LessCondsData{
			{
				Type:   datacentervo.ConditionEqual,
				Column: d.JsonColumn(consts.ColumnIdIssueId),
				Value:  rowId,
			},
			{
				Type:   datacentervo.ConditionEqual,
				Column: d.JsonColumn(consts.ColumnIdDelFlag),
				Value:  consts.DeleteFlagNotDel,
			},
		},
	}
	table := &datacentervo.Table{
		Type:   consts.TableType,
		Schema: d.WrapperTableName(orgId, tableId),
	}
	setsValue := []datacentervo.Set{
		{
			Column: consts.ColumnData,
			Type:   consts.JsonbFields,
			Value:  values,
		},
	}
	_, err := d.DatacenterExecutor(ctx, consts.DataSource, consts.DataBase, &datacentervo.Executor{
		Type:      consts.ExecutorUpdate,
		Table:     table,
		Sets:      setsValue,
		Condition: &lessCond,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *datacenterRepo) CreateTables(ctx context.Context, tables []*datacentervo.CreateTableReq) (*datacentervo.CreateTableResp, error) {
	respVo := &datacentervo.CreateTableResp{}
	path := fmt.Sprintf("/datacenter/inner/api/v1/%d/%d/tables-batch", consts.DataSource, consts.DataBase)
	err := d.dc.Invoke(ctx, "POST", path, tables, respVo)
	if err != nil {
		return nil, err
	}

	if respVo.Failure() {
		return nil, errors.Wrapf(err, "[CreateTables], req:%v, error:%v", tables, respVo.Error())
	}

	return respVo, nil
}

func (d *datacenterRepo) WrapperTableName(orgId, tableId int64) string {
	return "lc_data"
	//return fmt.Sprintf("_form_%d_%d", orgId, tableId)
}

func (d *datacenterRepo) JsonColumn(str string) string {
	return fmt.Sprintf("\"data\"::jsonb -> '%s'", str)
}
