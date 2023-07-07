package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cast"
	"github.com/star-table/go-common/pkg/encoding"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-common/pkg/middleware/meta"
	"github.com/star-table/go-table/internal/biz/bo"
	"github.com/star-table/go-table/internal/biz/bo/covert"
	"github.com/star-table/go-table/internal/biz/utils"
	"github.com/star-table/go-table/internal/data/consts"
	"github.com/star-table/go-table/internal/data/po"
	commonPb "github.com/star-table/interface/golang/common/v1"
	msgPb "github.com/star-table/interface/golang/msg/v1"
	tablePb "github.com/star-table/interface/golang/table/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/gorm"
)

func (t *TableUseCase) ReadTableSchemas(ctx context.Context, req *tablePb.ReadTableSchemasRequest) (*tablePb.ReadTableSchemasReply, error) {
	columns, err := t.tableRepo.GetColumnsByTables(ctx, req.TableIds, req.ColumnIds, req.IsNeedDescription)
	if err != nil {
		return nil, err
	}

	ch := meta.GetCommonHeaderFromCtx(ctx)
	summaryTableId, err := t.tableRepo.GetSummeryTableId(ctx, ch.OrgId)
	if err != nil {
		return nil, err
	}

	reply := &tablePb.ReadTableSchemasReply{Tables: make([]*tablePb.TableSchema, 0, len(columns))}
	for _, column := range columns {
		if req.IsNeedCommonColumn {
			column.Columns = t.mergeCommonColumns(column.Columns)
		} else {
			// 如果不需要通用表头，则将系统隐藏的字段给删除
			column.Columns = t.deleteSysHiddenColumns(column.Columns)
		}
		column.Columns = t.deleteNoNeedColumns(column.Columns)

		if column.TableId != summaryTableId {
			column, err = t.mergeUserSummaryColumn(ctx, column, summaryTableId, req.IsNeedCommonColumn)
			if err != nil {
				return nil, err
			}
		}

		reply.Tables = append(reply.Tables, &tablePb.TableSchema{
			TableId: column.TableId,
			Columns: column.Columns,
		})
	}

	// 需要描述的地方只有一个，就是前端请求，这个时候返回关联的列数据
	if req.IsNeedDescription && len(req.TableIds) == 1 && len(columns) == 1 {
		err = t.addRefColumn(ctx, columns[0].Columns)
		if err != nil {
			return nil, err
		}
	}

	return reply, nil
}

func (t *TableUseCase) addRefColumn(ctx context.Context, columns []*tablePb.Column) error {
	boColumns, err := t.tableRepo.GetRefColumns(ctx, columns)
	if err != nil {
		return err
	}

	refSettingPropsMap := make(map[string][]*structpb.Struct)
	for _, column := range columns {
		tableId, columnId := t.tableRepo.GetColumnRefTableInfo(column, columns)
		if columnId != "" {
			columnType := column.Field.GetType().String()
			key := fmt.Sprintf("%v_%v", tableId, columnId)
			refSettingPropsMap[key] = append(refSettingPropsMap[key], column.Field.Props.GetFields()[columnType].GetStructValue())
		}
	}

	for _, boColumn := range boColumns {
		refProps := refSettingPropsMap[fmt.Sprintf("%v_%v", boColumn.TableId, boColumn.Schema.Name)]
		if refProps != nil {
			value, err := utils.ModelToProtoStruct(boColumn.Schema)
			if err != nil {
				return err
			}
			for _, prop := range refProps {
				prop.GetFields()["column"] = value
			}
		}
	}

	return nil
}

func (t *TableUseCase) ReadTableSchemasByAppId(ctx context.Context, req *tablePb.ReadTableSchemasByAppIdRequest) (*tablePb.ReadTableSchemasByAppIdReply, error) {
	columns, err := t.tableRepo.GetAppTableColumns(ctx, req.AppId, req.ColumnIds)
	if err != nil {
		return nil, err
	}
	reply := &tablePb.ReadTableSchemasByAppIdReply{Tables: make([]*tablePb.TableSchema, 0, len(columns))}
	for _, column := range columns {
		if req.IsNeedCommonColumn {
			column.Columns = t.mergeCommonColumns(column.Columns)
		}
		column.Columns = t.deleteNoNeedColumns(column.Columns)
		reply.Tables = append(reply.Tables, &tablePb.TableSchema{
			TableId: column.TableId,
			Columns: column.Columns,
		})
	}

	return reply, nil
}

// 合并通用头
func (t *TableUseCase) mergeCommonColumns(columns []*tablePb.Column) []*tablePb.Column {
	if len(columns) == 0 {
		return columns
	}

	// 合并通用表头
	newColumns := make([]*tablePb.Column, 0, len(columns))
	newColumns = append(newColumns, columns...)
	for _, commonColumn := range consts.CommonColumns {
		isIn := false
		for _, column := range columns {
			if column.Name == commonColumn.Name {
				isIn = true
				break
			}
		}
		if !isIn {
			newColumns = append(newColumns, commonColumn)
		}
	}
	return newColumns
}

// 合并用户创建的表头
func (t *TableUseCase) mergeUserSummaryColumn(ctx context.Context, boColumns *bo.Columns, summaryTableId int64, isNeedDescription bool) (*bo.Columns, error) {
	tableInfo, err := t.tableRepo.GetTable(ctx, boColumns.TableId)
	if err != nil {
		return nil, err
	}
	appSummaryTableId, err := t.getAppSummaryTableId(ctx, tableInfo.AppId)
	if err != nil {
		return nil, err
	}

	tableIds := []int64{summaryTableId}
	if appSummaryTableId != 0 && appSummaryTableId != boColumns.TableId && tableInfo.SummeryFlag != consts.SummaryFlagFolder {
		tableIds = append(tableIds, appSummaryTableId)
	}

	summaryColumns, err := t.tableRepo.GetColumnsByTables(ctx, tableIds, nil, isNeedDescription)
	if err != nil {
		return nil, err
	}

	for _, column := range summaryColumns {
		for _, c := range column.Columns {
			if c.SummaryFlag == consts.SummaryFlagApp || c.SummaryFlag == consts.SummaryFlagAll {
				boColumns.Columns = append(boColumns.Columns, c)
			}
		}
	}

	return boColumns, nil
}

func (t *TableUseCase) deleteNoNeedColumns(columns []*tablePb.Column) []*tablePb.Column {
	newColumns := make([]*tablePb.Column, 0, len(columns))
	for _, column := range columns {
		if _, ok := consts.NoNeedColumnIdsMap[column.Name]; ok {
			continue
		}
		newColumns = append(newColumns, column)
	}

	return newColumns
}

func (t *TableUseCase) deleteSysHiddenColumns(columns []*tablePb.Column) []*tablePb.Column {
	newColumns := make([]*tablePb.Column, 0, len(columns))
	for _, column := range columns {
		if column.IsSysHiden && column.Field.Type.String() == tablePb.ColumnType_singleRelating.String() {
			continue
		}
		newColumns = append(newColumns, column)
	}

	return newColumns
}

func (t *TableUseCase) getColumnsByTables(ctx context.Context, tables []*po.Table, columnIds []string) ([]*bo.Columns, error) {
	tableIds := make([]int64, 0, len(tables))
	for _, table := range tables {
		tableIds = append(tableIds, table.ID)
	}
	return t.tableRepo.GetColumnsByTables(ctx, tableIds, columnIds, false)
}

// ReadOrgTableSchemas 读取组织内的所有表的某些的字段
func (t *TableUseCase) ReadOrgTableSchemas(ctx context.Context, req *tablePb.ReadOrgTableSchemasRequest) (*tablePb.ReadOrgTableSchemasReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	tables, err := t.tableRepo.GetTablesByOrgId(ctx, ch.OrgId)
	if err != nil {
		return nil, err
	}

	tablesMap := make(map[int64]*po.Table, len(tables))
	for _, table := range tables {
		tablesMap[table.ID] = table
	}

	columns, err := t.getColumnsByTables(ctx, tables, req.ColumnIds)
	if err != nil {
		return nil, err
	}
	reply := &tablePb.ReadOrgTableSchemasReply{Tables: make([]*tablePb.TableSchema, 0, len(columns))}
	for _, column := range columns {
		reply.Tables = append(reply.Tables, &tablePb.TableSchema{
			AppId:   tablesMap[column.TableId].AppId,
			TableId: column.TableId,
			Name:    tablesMap[column.TableId].Name,
			Columns: column.Columns,
		})
	}

	return reply, nil
}

func (t *TableUseCase) CreateColumn(ctx context.Context, req *tablePb.CreateColumnRequest) (*tablePb.CreateColumnReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)

	relateTableId := cast.ToInt64(t.tableRepo.GetColumnPropsStringValue(req.Column, consts.RelateTableId))
	// 关联比较特殊，不是只能一列，下面会改列名
	if relateTableId == 0 {
		oldColumn, err := t.tableRepo.GetColumn(ctx, req.TableId, req.Column.Name, true)
		if err != nil && errors.Cause(err) != gorm.ErrRecordNotFound {
			return nil, err
		}

		if oldColumn != nil {
			return &tablePb.CreateColumnReply{AppId: req.AppId, TableId: req.TableId}, nil
		}
	}

	// 如果创建的时候来源是组织字段，用组织字段覆盖，防止被改了
	if req.SourceOrgColumnId != "" {
		isCreated, err := t.resetRequestOrgColumn(ctx, ch.OrgId, req)
		if err != nil {
			return nil, err
		}
		// 如果已经创建过，直接返回成功
		if isCreated {
			return &tablePb.CreateColumnReply{AppId: req.AppId, TableId: req.TableId}, nil
		}

		req.Column, err = t.resetColumnWithOrg(ctx, ch, req)
		if err != nil {
			return nil, err
		}
	} else {
		// 如果是关联普通表，则需要改掉id
		if strings.Contains(req.Column.Name, consts.UserDefineColumnPrefix) || relateTableId > 0 {
			req.Column.Name = t.tableRepo.GetNewColumnId(consts.ColumnCategoryNormal)
		}
	}

	// 设置列的类型
	summaryFlag, err := t.getTableSummaryFlag(ctx, req.TableId)
	if err != nil {
		return nil, err
	}
	req.Column.SummaryFlag = summaryFlag

	poColumn, err := covert.ColumnCovert.ToPoColumn(req.Column, ch.OrgId, req.TableId)
	if err != nil {
		return nil, err
	}
	poColumn.SourceOrgColumnId = req.SourceOrgColumnId

	err = t.tableRepo.StartTransactionLcGo(ctx, func(tx *gorm.DB) error {
		err = t.tableRepo.CreateColumn(ctx, poColumn, tx)
		if err != nil {
			return err
		}

		// 如果是表关联逻辑，则在关联表里创建同一列关联表头
		if (req.Column.Field.Type.String() == tablePb.ColumnType_relating.String() || req.Column.Field.Type.String() == tablePb.ColumnType_singleRelating.String()) &&
			relateTableId != req.TableId {

			return t.createRelateColumn(ctx, ch.OrgId, req.TableId, req.Column, tx)
		}

		return nil
	})

	if err == nil {
		column := make(map[string]interface{})
		err = ProtoJsonCopy(req.Column, &column)
		if err == nil {
			t.reportEvent(ctx, ch.OrgId, req.AppId, req.TableId, ch.UserId, msgPb.EventType_TableColumnRefresh, column, nil)
		}
	}
	return &tablePb.CreateColumnReply{AppId: req.AppId, TableId: req.TableId, Column: req.Column}, err
}

func (t *TableUseCase) getTableSummaryFlag(ctx context.Context, tableId int64) (int32, error) {
	tableInfo, err := t.tableRepo.GetTable(ctx, tableId)
	if err != nil {
		return 0, err
	}

	return tableInfo.SummeryFlag, nil
}

func (t *TableUseCase) getAppSummaryTableId(ctx context.Context, appId int64) (int64, error) {
	tables, err := t.tableRepo.GetTables(ctx, appId)
	if err != nil {
		return 0, err
	}

	for _, table := range tables {
		if table.SummeryFlag == consts.SummaryFlagApp {
			return table.ID, nil
		}
	}

	return 0, nil
}

// 创建关联列
func (t *TableUseCase) createRelateColumn(ctx context.Context, orgId, currentTableId int64, column *tablePb.Column, tx *gorm.DB) error {
	relateTableId := cast.ToInt64(t.tableRepo.GetColumnPropsStringValue(column, consts.RelateTableId))
	// 不是关联的普通表，不需要给关联表创建新的列
	if relateTableId <= 0 {
		return nil
	}

	relateInfo, err := t.getRelateColumnInfo(ctx, relateTableId, currentTableId)
	if err != nil {
		return err
	}

	// 修改名字，还有相关关联属性
	newColumn := &tablePb.Column{}
	err = utils.CopyProto(column, newColumn)
	if err != nil {
		return err
	}

	newColumn.Label = relateInfo.Name
	newColumn.AliasTitle = ""
	// 如果是单向关联，则这个表头在外部不显示，在内部需要使用
	newColumn.IsSysHiden = newColumn.Field.Type.String() == tablePb.ColumnType_singleRelating.String()
	propFields := newColumn.Field.GetProps().Fields[newColumn.Field.Type.String()]
	propFields.GetStructValue().Fields[consts.RelateTableId] = structpb.NewStringValue(cast.ToString(currentTableId))
	propFields.GetStructValue().Fields[consts.RelateAppId] = structpb.NewStringValue(cast.ToString(relateInfo.AppId))
	relatePoColumn, err := covert.ColumnCovert.ToPoColumn(newColumn, orgId, relateTableId)
	if err != nil {
		return err
	}

	return t.tableRepo.CreateColumn(ctx, relatePoColumn, tx)
}

// 如果是org的表头创建的话，重置下，不需要客户端传入的值，如果已经创建过，直接返回
func (t *TableUseCase) resetRequestOrgColumn(ctx context.Context, orgId int64, req *tablePb.CreateColumnRequest) (bool, error) {
	hadUse, err := t.tableRepo.CheckOrgColumnIdHadUseInTable(ctx, req.TableId, req.SourceOrgColumnId)
	if err != nil {
		return false, err
	}
	// 如果这个字段已经在使用了，则不建立新的了，直接返回，目前一个团队字段在一个表格里面只能出现一次
	if hadUse {
		return true, nil
	}

	cs, err := t.orgRepo.GetColumns(ctx, orgId, []string{req.SourceOrgColumnId})
	if err != nil {
		return false, err
	}
	if len(cs) == 0 {
		return false, errors.Ignore(commonPb.ErrorResourceNotExist("can not find org column"))
	}
	req.Column = cs[0].Schema
	req.Column.IsOrg = true

	return false, nil
}

// resetColumn 重置下orgColumn相关信息，只保留别名和默认值，其他都是用库里面的org字段，免的被修改了
func (t *TableUseCase) resetColumnWithOrg(ctx context.Context, ch *meta.CommonHeader, req *tablePb.CreateColumnRequest) (*tablePb.Column, error) {
	cs, err := t.orgRepo.GetColumns(ctx, ch.OrgId, []string{req.SourceOrgColumnId})
	if err != nil {
		return nil, err
	}
	if len(cs) == 0 {
		return nil, errors.Ignore(commonPb.ErrorResourceNotExist("can not find org column"))
	}

	t.setOrgColumnCanUpdateProps(cs[0].Schema, req.Column)

	return req.Column, nil
}

// 将能修改的org字段值赋值，其他用库里的值
func (t *TableUseCase) setOrgColumnCanUpdateProps(orgColumn *tablePb.Column, reqColumn *tablePb.Column) {
	defaultValue := reqColumn.Field.Props.Fields["default"]
	if defaultValue != nil {
		defaultValueBts, err := defaultValue.MarshalJSON()
		if err == nil {
			v := &structpb.Value{}
			err = v.UnmarshalJSON(defaultValueBts)
			if err != nil {
				orgColumn.Field.Props.Fields["default"] = v
			}
		}
	}
	orgColumn.AliasTitle = reqColumn.AliasTitle

	reqColumn = orgColumn
	reqColumn.IsOrg = true
}

// CopyColumn 拷贝列
func (t *TableUseCase) CopyColumn(ctx context.Context, req *tablePb.CopyColumnRequest) (*tablePb.CopyColumnReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)

	// 获取原来的列
	oldColumn, err := t.tableRepo.GetColumn(ctx, req.TableId, req.SrcColumnId, true)
	if err != nil {
		return nil, err
	}

	// 如果不允许复制，直接返回
	if _, ok := consts.CanNotCopyAndOrgColumns[req.SrcColumnId]; ok {
		return &tablePb.CopyColumnReply{CreateColumnId: oldColumn.ColumnId}, nil
	}

	// 更新列名
	oldColumn.ColumnId = t.tableRepo.GetNewColumnId(consts.ColumnCategoryNormal)
	oldColumn.Schema.Name = oldColumn.ColumnId
	// copy后的列名和别名如果有值不能和之前的一样
	newLabel := oldColumn.Schema.Label + oldColumn.ColumnId
	oldColumn.Schema.Label = newLabel
	oldColumn.Schema.AliasTitle = newLabel

	poColumn, err := covert.ColumnCovert.ToPoColumn(oldColumn.Schema, ch.OrgId, req.TableId)
	if err != nil {
		return nil, err
	}

	summeryTableId, err := t.tableRepo.GetSummeryTableId(ctx, ch.OrgId)
	if err != nil {
		return nil, err
	}

	// 存到数据，而且拷贝列值
	err = t.tableRepo.StartTransactionLcGo(ctx, func(tx *gorm.DB) error {
		err := t.tableRepo.CreateColumn(ctx, poColumn, tx)
		if err != nil {
			return err
		}

		tableId := req.TableId
		if tableId == summeryTableId {
			tableId = 0
		}
		err = t.rows.copyColumnData(ctx, ch.OrgId, tableId, req.SrcColumnId, oldColumn.ColumnId)
		if err != nil {
			return err
		}

		// 如果原来的列打开了协作者模式，则还要复制一份协作者数据
		if t.tableRepo.CheckColumnCollaboratorSwitchOn(oldColumn.Schema) {
			err = t.rows.CopyColumnCollaborator(ctx, ch.OrgId, req.AppId, req.TableId, req.SrcColumnId, oldColumn.Schema.Name)
		}

		return err
	})

	if err == nil {
		column := make(map[string]interface{})
		err = ProtoJsonCopy(oldColumn.Schema, &column)
		if err == nil {
			t.reportEvent(ctx, ch.OrgId, req.AppId, req.TableId, ch.UserId, msgPb.EventType_TableColumnCopyed, column, nil)
		}
	}

	return &tablePb.CopyColumnReply{CreateColumnId: oldColumn.ColumnId}, err
}

func (t *TableUseCase) UpdateColumn(ctx context.Context, req *tablePb.UpdateColumnRequest) (*tablePb.UpdateColumnReply, error) {
	// 更新column的时候不需要存这个参数，防止客户端传上来，然后存起来，会非常大
	ch := meta.GetCommonHeaderFromCtx(ctx)
	req.Column.Description = ""

	allColumns, err := t.tableRepo.GetColumns(ctx, req.TableId)
	if err != nil {
		return nil, err
	}

	if err = t.checkColumnNameConflit(ctx, allColumns, req); err != nil {
		return nil, err
	}

	oldColumn, err := t.tableRepo.GetColumn(ctx, req.TableId, req.Column.Name, true)
	if err != nil {
		return nil, err
	}

	// 校验appId是否正常，现在客户端会传一个不是这个tableId的app过来，导致问题
	table, err := t.tableRepo.GetTable(ctx, req.TableId)
	if err != nil {
		return nil, err
	}
	if table.AppId != req.AppId {
		return nil, commonPb.ErrorParamsNotCorrect(fmt.Sprintf("AppId:%v and TableId:%v is not match", req.AppId, req.TableId))
	}

	// 如果是组织字段跟系统字段转换
	if oldColumn.Schema.IsOrg != req.Column.IsOrg {
		// 如果不允许操作的，直接返回
		if _, ok := consts.CanNotCopyAndOrgColumns[req.Column.Name]; ok {
			return &tablePb.UpdateColumnReply{AppId: req.AppId, TableId: req.TableId}, nil
		}

		newColumnId, err := t.changeOrgColumn(ctx, ch.OrgId, req.AppId, req.TableId, oldColumn.Schema, req.Column)
		return &tablePb.UpdateColumnReply{AppId: req.AppId, TableId: req.TableId, Column: &tablePb.Column{Name: newColumnId}}, err
	}

	// 正常字段修改
	err = t.tableRepo.StartTransactionLcGo(ctx, func(tx *gorm.DB) error {
		// 如果是组织字段，只能改为别名，其他不允许更改
		if oldColumn.Schema.IsOrg {
			t.setOrgColumnCanUpdateProps(oldColumn.Schema, req.Column)
		}

		// 如果两个类型一致，则证明没有变化类型，如果是修改groupSelect的话，有可能要同步修改数据的分类
		if oldColumn.ColumnType == req.Column.Field.Type.String() {
			switch oldColumn.ColumnType {
			case tablePb.ColumnType_groupSelect.String():
				if req.Column.Name == consts.IssueStatusField {
					err = t.updateGroupSelectRows(ctx, oldColumn, req)
				}
			case tablePb.ColumnType_relating.String(), tablePb.ColumnType_singleRelating.String():
				// 有可能会修改column的id，如果从汇总表和普通表直接转换的话
				req.Column, err = t.checkChangeRelateTableId(ctx, ch.OrgId, req.TableId, oldColumn.Schema, req.Column, tx)
			//case tablePb.ColumnType_singleRelating.String():
			//	err = t.checkChangeSingleRelateTableId(ctx, req.TableId, oldColumn.Schema, req.Column)
			default:
			}
		}

		// 更新下column
		poColumn, err := covert.ColumnCovert.ToPoColumn(req.Column, ch.OrgId, req.TableId)
		if err != nil {
			return err
		}
		err = t.tableRepo.UpdateColumnWithOldColumnId(ctx, poColumn, oldColumn.ColumnId, tx)
		if err != nil {
			return err
		}

		err = t.checkCollaborator(ctx, ch.OrgId, req.AppId, req.TableId, oldColumn.Schema, req.Column)
		if err != nil {
			return err
		}

		return err
	})

	if err == nil {
		column := make(map[string]interface{})
		old := make(map[string]interface{})
		err = ProtoJsonCopy(req.Column, &column)
		err2 := ProtoJsonCopy(oldColumn.Schema, &old)
		if err == nil && err2 == nil {
			t.reportEvent(ctx, ch.OrgId, req.AppId, req.TableId, ch.UserId, msgPb.EventType_TableColumnRefresh, column, old)
		}
	}

	return &tablePb.UpdateColumnReply{AppId: req.AppId, TableId: req.TableId, Column: req.Column}, err
}

func (t *TableUseCase) getRelateColumnInfo(ctx context.Context, relateTableId, currentTableId int64) (*bo.RelateColumnInfo, error) {
	tableInfo, err := t.tableRepo.GetTable(ctx, currentTableId)
	if err != nil {
		return nil, err
	}
	columns, err := t.tableRepo.GetColumns(ctx, relateTableId)
	if err != nil {
		return nil, err
	}
	columnNamesMap := make(map[string]struct{}, len(columns))
	for _, column := range columns {
		columnNamesMap[column.Schema.Label] = struct{}{}
	}

	tempTableName := tableInfo.Name
	for i := 1; i < len(columns)+2; i++ {
		if _, ok := columnNamesMap[tempTableName]; !ok {
			break
		}
		tempTableName = tableInfo.Name + cast.ToString(i)
	}

	return &bo.RelateColumnInfo{
		Name:  tempTableName,
		AppId: tableInfo.AppId,
	}, nil
}

func (t *TableUseCase) checkColumnNameConflit(ctx context.Context, allColumns []*bo.Column, req *tablePb.UpdateColumnRequest) error {
	allNames := map[string]struct{}{}
	for _, column := range allColumns {
		if column.Schema.Name != req.Column.Name {
			if column.Schema.Label != "" {
				allNames[column.Schema.Label] = struct{}{}
			}
			if column.Schema.AliasTitle != "" {
				allNames[column.Schema.AliasTitle] = struct{}{}
			}
		}
	}

	if req.Column.Label != "" {
		if _, ok := allNames[req.Column.Label]; ok {
			return commonPb.ErrorParamsNotCorrect(fmt.Sprintf("Column Name: %v is not unique", req.Column.Label))
		}
	}
	if req.Column.AliasTitle != "" {
		if _, ok := allNames[req.Column.AliasTitle]; ok {
			return commonPb.ErrorParamsNotCorrect(fmt.Sprintf("Column Name: %v is not unique", req.Column.AliasTitle))
		}
	}
	return nil
}

// 检查下是否关联字段改变了绑定的表，如果改变了要做老关联表字段删除处理，新的表增加列操作
// 这里有三种可能：
// 1、从汇总表变成普通表；需要改变表头名字变为固定名字relating，
// 2、从一个普通表变成汇总表；
// 3、从一个表变为另一个表
func (t *TableUseCase) checkChangeRelateTableId(ctx context.Context, orgId, currentTableId int64, oldColumn, newColumn *tablePb.Column, tx *gorm.DB) (*tablePb.Column, error) {
	oldRelateTableId := cast.ToInt64(t.tableRepo.GetColumnPropsStringValue(oldColumn, consts.RelateTableId))
	newRelateTableId := cast.ToInt64(t.tableRepo.GetColumnPropsStringValue(newColumn, consts.RelateTableId))

	if oldRelateTableId != newRelateTableId {
		if oldRelateTableId == 0 {
			newColumn.Name = t.tableRepo.GetNewColumnId(consts.ColumnCategoryNormal)
			if newRelateTableId != currentTableId {
				err := t.createRelateColumn(ctx, orgId, currentTableId, newColumn, tx)
				if err != nil {
					return nil, err
				}
			}

			// 新关联的表创建表头
			return newColumn, nil
		} else if newRelateTableId == 0 {
			// 删除老关联表表头以及记录
			if oldRelateTableId != currentTableId {
				err := t.deleteColumnAndData(ctx, orgId, oldRelateTableId, oldColumn.Name, tx)
				if err != nil {
					return newColumn, err
				}
			}
			// 删除本表数据
			err := t.rows.deleteColumnData(ctx, orgId, currentTableId, oldColumn.Name)
			if err != nil {
				return newColumn, err
			}
			newColumn.Name = consts.ColumnIdRelating
		} else {
			if oldRelateTableId != currentTableId {
				// 删除老关联表表头以及记录
				err := t.deleteColumnAndData(ctx, orgId, oldRelateTableId, oldColumn.Name, tx)
				if err != nil {
					return newColumn, err
				}
			}

			// 删除本表数据
			err := t.rows.deleteColumnData(ctx, orgId, currentTableId, newColumn.Name)
			if err != nil {
				return newColumn, err
			}
			if newRelateTableId != currentTableId {
				// 新关联的表创建表头
				return newColumn, t.createRelateColumn(ctx, orgId, currentTableId, newColumn, tx)
			}
		}
	}

	return newColumn, nil
}

// checkChangeSingleRelateTableId 单向关联的时候切换表只删除数据
func (t *TableUseCase) checkChangeSingleRelateTableId(ctx context.Context, currentTableId int64, oldColumn, newColumn *tablePb.Column) error {
	var (
		oldRelateTableId = int64(0)
		newRelateTableId = int64(0)
	)
	oldProps := oldColumn.Field.Props.Fields[oldColumn.Field.Type.String()]
	if oldProps != nil && oldProps.GetStructValue() != nil && oldProps.GetStructValue().Fields[consts.RelateTableId] != nil {
		oldRelateTableId = cast.ToInt64(oldProps.GetStructValue().Fields[consts.RelateTableId].GetStringValue())
	}
	newProps := newColumn.Field.Props.Fields[oldColumn.Field.Type.String()]
	if newProps != nil && newProps.GetStructValue() != nil && newProps.GetStructValue().Fields[consts.RelateTableId] != nil {
		newRelateTableId = cast.ToInt64(newProps.GetStructValue().Fields[consts.RelateTableId].GetStringValue())
	}

	ch := meta.GetCommonHeaderFromCtx(ctx)
	if oldRelateTableId != newRelateTableId && oldRelateTableId != 0 && newRelateTableId != 0 {
		// 删除本表数据
		err := t.rows.deleteColumnData(ctx, ch.OrgId, currentTableId, newColumn.Name)
		if err != nil {
			return err
		}
		// 删除关联表数据
		err = t.rows.deleteColumnData(ctx, ch.OrgId, oldRelateTableId, newColumn.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

// changeOrgColumn 系统字段和团队字段互转
func (t *TableUseCase) changeOrgColumn(ctx context.Context, orgId, appId, tableId int64, oldColumn, newColumn *tablePb.Column) (string, error) {
	if oldColumn.IsOrg == true {
		// 团队字段转为系统字段，换个名字，一个团队字段只能出现一次，所以要改下名字
		return t.changeOrgToSys(ctx, orgId, appId, tableId, oldColumn, newColumn)
	} else {
		// 系统字段转为团队字段，操作为新增一个系统字段，并更新下
		return t.changeSysToOrg(ctx, orgId, appId, tableId, oldColumn, newColumn)
	}
}

// changeOrgToSys 团队字段变为系统字段，这个时候只需要把关联的sourceOrgColumnId置空，和更新下schema

func (t *TableUseCase) changeOrgToSys(ctx context.Context, orgId, appId, tableId int64, oldColumn, newColumn *tablePb.Column) (string, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)

	oldColumnId := newColumn.Name
	newColumn.Name = t.tableRepo.GetNewColumnId(consts.ColumnCategoryNormal)
	poColumn, err := covert.ColumnCovert.ToPoColumn(newColumn, orgId, tableId)
	if err != nil {
		return "", err
	}

	err = t.tableRepo.StartTransactionLcGo(ctx, func(tx *gorm.DB) error {
		// 这个时候置空下sourceOrgColumnId，并更新下schema
		err = t.tableRepo.UpdateColumnAndResetOrgColumnId(ctx, poColumn, oldColumnId, tx)
		if err != nil {
			return err
		}

		return t.rows.moveColumnData(ctx, tableId, oldColumnId, newColumn.Name, t.tableRepo.CheckColumnCollaboratorSwitchOn(newColumn))

	})

	if err == nil {
		newC := make(map[string]interface{})
		oldC := make(map[string]interface{})
		err1 := ProtoJsonCopy(newColumn, &newC)
		err2 := ProtoJsonCopy(oldColumn, &oldC)
		if err1 == nil && err2 == nil {
			t.reportEvent(ctx, ch.OrgId, appId, tableId, ch.UserId, msgPb.EventType_TableOrgColumnRefresh, newC, oldC)
		}
	}

	return newColumn.Name, err
}

// changeSysToOrg 系统字段转团队字段
func (t *TableUseCase) changeSysToOrg(ctx context.Context, orgId, appId, tableId int64, oldColumn, newColumn *tablePb.Column) (string, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)

	newColumn.Description = oldColumn.Description
	poOrgColumn, err := covert.OrgColumnCovert.ToPoColumn(newColumn, orgId)
	if err != nil {
		return newColumn.Name, err
	}
	// 创建新的团队字段
	err = t.orgRepo.CreateColumns(ctx, orgId, []*po.OrgColumn{poOrgColumn})
	if err != nil {
		return newColumn.Name, err
	}

	poColumn, err := covert.ColumnCovert.ToPoColumn(newColumn, orgId, tableId)
	if err != nil {
		return newColumn.Name, err
	}
	// 绑定团队字段
	poColumn.SourceOrgColumnId = newColumn.Name
	err = t.tableRepo.UpdateColumn(ctx, poColumn)

	if err == nil {
		newC := make(map[string]interface{})
		oldC := make(map[string]interface{})
		err1 := ProtoJsonCopy(newColumn, &newC)
		err2 := ProtoJsonCopy(oldColumn, &oldC)
		if err1 == nil && err2 == nil {
			t.reportEvent(ctx, ch.OrgId, appId, tableId, ch.UserId, msgPb.EventType_TableOrgColumnRefresh, newC, oldC)
		}
	}

	return newColumn.Name, err
}

func (t *TableUseCase) checkCollaborator(ctx context.Context, orgId, appId, tableId int64, oldColumn, newColumn *tablePb.Column) error {
	oldOpen := t.tableRepo.CheckColumnCollaboratorSwitchOn(oldColumn)
	newOpen := t.tableRepo.CheckColumnCollaboratorSwitchOn(newColumn)
	if oldOpen != newOpen {
		if oldOpen == true {
			return t.rows.SwitchColumnCollaboratorOff(ctx, orgId, appId, tableId, oldColumn.Name)
		} else {
			return t.rows.SwitchColumnCollaboratorOn(ctx, orgId, appId, tableId, oldColumn.Name)
		}
	}

	return nil
}

func (t *TableUseCase) UpdateColumnDescription(ctx context.Context, req *tablePb.UpdateColumnDescriptionRequest) (*tablePb.UpdateColumnDescriptionReply, error) {
	err := t.tableRepo.UpdateColumnDescription(ctx, req.TableId, req.ColumnId, req.Description)

	return &tablePb.UpdateColumnDescriptionReply{AppId: req.AppId, TableId: req.TableId}, err
}

// updateGroupSelectRows 更新groupSelect分类的同时更新数据分类
func (t *TableUseCase) updateGroupSelectRows(ctx context.Context, oldColumn *bo.Column, req *tablePb.UpdateColumnRequest) error {
	_, changeType := t.checkGroupSelectOptions(oldColumn.Schema, req.Column)
	if len(changeType) > 0 {
		return t.rows.updateRowsGroupSelectType(ctx, req.TableId, req.Column.Name, changeType)
	}

	return nil
}

func (t *TableUseCase) checkGroupSelectOptions(oldColumn, newColumn *tablePb.Column) ([]int64, map[int64][]interface{}) {
	// 新老选项进行对比，将删除的和变化父类的选项拿出来，这个是要刷数据的
	oldOptions := t.getGroupSelectOptions(oldColumn)
	newOptions := t.getGroupSelectOptions(newColumn)
	deletes := make([]int64, 0, 1)
	changeType := make(map[int64][]interface{}, 1)
	for _, os := range oldOptions {
		if newOptions[os.Id] == nil {
			deletes = append(deletes, os.Id)
		} else {
			parentId := newOptions[os.Id].ParentId
			if parentId != os.ParentId {
				changeType[parentId] = append(changeType[parentId], os.Id)
			}
		}
	}

	return deletes, changeType
}

func (t *TableUseCase) getGroupSelectOptions(column *tablePb.Column) map[int64]*bo.GroupSelectOption {
	bgs := make(map[int64]*bo.GroupSelectOption, 3)
	groupSelect := column.Field.Props.Fields["groupSelect"]
	if groupSelect != nil {
		for _, value := range groupSelect.GetStructValue().Fields["options"].GetListValue().GetValues() {
			id := cast.ToInt64(value.GetStructValue().Fields["id"].GetNumberValue())
			parentId := cast.ToInt64(value.GetStructValue().Fields["parentId"].GetNumberValue())
			if id != 0 && parentId != 0 {
				bgs[id] = &bo.GroupSelectOption{
					Id:       id,
					ParentId: parentId,
				}
			}
		}
	}

	return bgs
}

func (t *TableUseCase) DeleteColumn(ctx context.Context, req *tablePb.DeleteColumnRequest) (*tablePb.DeleteColumnReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)

	oldColumn, err := t.tableRepo.GetColumn(ctx, req.TableId, req.ColumnId, false)
	if err != nil {
		return nil, err
	}

	summeryTableId, err := t.tableRepo.GetSummeryTableId(ctx, ch.OrgId)
	if err != nil {
		return nil, err
	}
	err = t.tableRepo.StartTransactionLcGo(ctx, func(tx *gorm.DB) error {
		err = t.deleteColumnAndData(ctx, ch.OrgId, req.TableId, req.ColumnId, tx)
		if err != nil {
			return err
		}

		// 如果是关联字段，则删除对应的表的关联字段，关联字段是成对出现的
		if oldColumn.ColumnType == tablePb.ColumnType_relating.String() || oldColumn.ColumnType == tablePb.ColumnType_singleRelating.String() {
			err = t.deleteRelateColumn(ctx, ch.OrgId, oldColumn, tx)
			if err != nil {
				return err
			}
		}
		tableId := req.TableId
		if req.TableId == summeryTableId {
			tableId = 0
		}
		err = t.rows.deleteColumnData(ctx, ch.OrgId, tableId, req.ColumnId)
		if err != nil {
			return err
		}

		if t.tableRepo.CheckColumnCollaboratorSwitchOn(oldColumn.Schema) {
			err = t.rows.SwitchColumnCollaboratorOff(ctx, ch.OrgId, req.AppId, req.TableId, req.ColumnId)
			if err != nil {
				return err
			}
		}

		return errors.WithStack(err)
	})

	if err == nil {
		t.reportEvent(ctx, ch.OrgId, req.AppId, req.TableId, ch.UserId, msgPb.EventType_TableColumnDeleted, req.ColumnId, nil)
	}
	return &tablePb.DeleteColumnReply{AppId: req.AppId, TableId: req.TableId, ColumnId: req.ColumnId}, err
}

func (t *TableUseCase) deleteColumnAndData(ctx context.Context, orgId, tableId int64, columnId string, tx *gorm.DB) error {
	err := t.tableRepo.DeleteColumn(ctx, tableId, columnId, tx)
	if err != nil {
		return err
	}

	return t.rows.deleteColumnData(ctx, orgId, tableId, columnId)
}

func (t *TableUseCase) deleteRelateColumn(ctx context.Context, orgId int64, column *bo.Column, tx *gorm.DB) error {
	relateTableId := cast.ToInt64(t.tableRepo.GetColumnPropsStringValue(column.Schema, consts.RelateTableId))
	if relateTableId <= 0 {
		return nil
	}

	return t.deleteColumnAndData(ctx, orgId, relateTableId, column.ColumnId, tx)
}

// GetMergeColumns 创建表的时候，需要拿汇总表的表头和formbase里面的基础字段进行合并
func (t *TableUseCase) getMergeColumns(ctx context.Context, orgId int64, orgColumnIds, notNeedSummeryColumnIds []string) ([]*tablePb.Column, error) {
	summeryTableId, err := t.tableRepo.GetSummeryTableId(ctx, orgId)
	if err != nil {
		return nil, err
	}
	summeryColumns, err := t.tableRepo.GetColumns(ctx, summeryTableId)
	if err != nil {
		return nil, err
	}

	// 获取基础表头
	orgColumns, err := t.getOrgColumns(ctx, orgId, orgColumnIds)
	if err != nil {
		return nil, err
	}

	// 合并两个表头
	mergeColumns := make([]*tablePb.Column, 0, len(summeryColumns)+len(orgColumns)+1)
	for _, sc := range summeryColumns {
		if sc.Schema.SummaryFlag == consts.SummaryFlagAll {
			continue
		}
		isIn := false
		for _, id := range notNeedSummeryColumnIds {
			if id == sc.ColumnId {
				isIn = true
			}
		}
		if !isIn {
			mergeColumns = append(mergeColumns, sc.Schema)
		}
	}
	mergeColumns = append(mergeColumns, orgColumns...)

	return mergeColumns, nil
}

// getOrgColumns 获取组织字段
func (t *TableUseCase) getOrgColumns(ctx context.Context, orgId int64, orgColumnIds []string) ([]*tablePb.Column, error) {
	if len(orgColumnIds) == 0 {
		return []*tablePb.Column{}, nil
	}

	boOrgColumns, err := t.orgRepo.GetColumns(ctx, orgId, orgColumnIds)
	if err != nil {
		return nil, err
	}
	cs := make([]*tablePb.Column, 0, len(boOrgColumns))

	for _, oc := range boOrgColumns {
		cs = append(cs, oc.Schema)
	}

	return cs, nil
}

func ProtoJsonCopy(src proto.Message, dst interface{}) error {
	jsonStr, err := encoding.GetJsonCodec().Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonStr, dst)
	if err != nil {
		return err
	}
	return nil
}
