package data

import (
	"context"
	"fmt"
	"math/rand"
	"sort"

	"github.com/spf13/cast"

	"google.golang.org/protobuf/types/known/structpb"

	"github.com/go-redis/redis/v8"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-common/pkg/middleware/meta"
	"github.com/star-table/go-table/internal/biz/bo"
	"github.com/star-table/go-table/internal/biz/bo/covert"
	"github.com/star-table/go-table/internal/biz/utils"
	"github.com/star-table/go-table/internal/data/consts"
	"github.com/star-table/go-table/internal/data/po"
	pb "github.com/star-table/interface/golang/table/v1"
	"gorm.io/gorm"
)

func (t *tableRepo) GetColumnsMap(ctx context.Context, tableId int64) (map[string]*bo.Column, error) {
	columns, err := t.GetColumns(ctx, tableId)
	if err != nil {
		return nil, err
	}

	m := make(map[string]*bo.Column)
	for _, c := range columns {
		m[c.ColumnId] = c
	}
	return m, nil
}

func (t *tableRepo) GetColumns(ctx context.Context, tableId int64) ([]*bo.Column, error) {
	columns, err := t.getColumns(ctx, tableId, nil)
	if err != nil {
		return nil, err
	}

	columns = t.sortColumns(columns)

	return covert.ColumnCovert.ToColumns(columns)
}

func (t *tableRepo) GetRefColumns(ctx context.Context, columns []*pb.Column) (map[string]*bo.Column, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	tableIds := make([]int64, 0, 1)
	tableIdsMap := make(map[int64]struct{})
	columnIds := make([]string, 0, 1)
	columnToRefColumnMap := make(map[string][]string)
	getKeyFunc := func(tableId int64, columnId string) string {
		return cast.ToString(tableId) + columnId
	}

	summeryTableId, err := t.GetSummeryTableId(ctx, ch.OrgId)
	if err != nil {
		return nil, err
	}
	tableIds = append(tableIds, summeryTableId)

	for _, column := range columns {
		tableId, columnId := t.GetColumnRefTableInfo(column, columns)
		if columnId != "" {
			if _, ok := tableIdsMap[tableId]; !ok && tableId != 0 {
				tableIds = append(tableIds, tableId)
			}
			tableIdsMap[tableId] = struct{}{}
			columnIds = append(columnIds, columnId)
			key := getKeyFunc(tableId, columnId)
			columnToRefColumnMap[key] = append(columnToRefColumnMap[key], column.Name)
		}
	}

	if len(columnToRefColumnMap) > 0 {
		boColumns, err := t.GetColumnsByTables(ctx, tableIds, columnIds, false)
		if err != nil {
			return nil, err
		}
		result := make(map[string]*bo.Column, len(boColumns))
		for _, tableColumns := range boColumns {
			for _, column := range tableColumns.Columns {
				if tableColumns.TableId == summeryTableId {
					tableColumns.TableId = 0
				}
				key := getKeyFunc(tableColumns.TableId, column.Name)
				for _, s := range columnToRefColumnMap[key] {
					result[s] = &bo.Column{
						TableId:    tableColumns.TableId,
						ColumnId:   column.Name,
						ColumnType: column.Field.Type.String(),
						Schema:     column,
					}
				}
			}
		}

		return result, nil
	}

	return map[string]*bo.Column{}, nil
}

func (t *tableRepo) GetColumnRefTableInfo(column *pb.Column, columns []*pb.Column) (int64, string) {
	columnType := column.Field.GetType().String()
	if column.Field.Props != nil {
		props := column.Field.Props.GetFields()[columnType]
		if props != nil && props.GetStructValue() != nil {
			switch columnType {
			case pb.ColumnType_conditionRef.String():
				return t.getConditionRefTableInfo(props)
			case pb.ColumnType_reference.String():
				return t.getReferenceTableInfo(props, columns)
			}
		}
	}

	return 0, ""
}

func (t *tableRepo) GetColumnPropsStringValue(column *pb.Column, key string) string {
	columnType := column.Field.GetType().String()
	if column.Field.Props != nil {
		props := column.Field.Props.GetFields()[columnType]
		if props != nil && props.GetStructValue() != nil && props.GetStructValue().Fields[key] != nil {
			return props.GetStructValue().Fields[key].GetStringValue()
		}
	}

	return ""
}

func (t *tableRepo) SetColumnPropsValues(column *pb.Column, values map[string]interface{}) error {
	columnType := column.Field.GetType().String()
	if column.Field.Props != nil {
		props := column.Field.Props.GetFields()[columnType]
		if props != nil && props.GetStructValue() != nil {
			for key, value := range values {
				temp, err := structpb.NewValue(value)
				if err != nil {
					return errors.WithStack(err)
				}
				props.GetStructValue().Fields[key] = temp
			}
		}
	}

	return nil
}

func (t *tableRepo) getConditionRefTableInfo(props *structpb.Value) (int64, string) {
	m := &pb.ConditionRefSetting{}
	err := utils.ProtoStructToModel(props.GetStructValue(), m)
	if err != nil {
		t.log.Errorf("[addRefColumn] ProtoStructToModel err:%v", err)
	}
	if m.TableId > 0 && m.ColumnId != "" {
		return m.TableId, m.ColumnId
	}

	return 0, ""
}

func (t *tableRepo) getReferenceTableInfo(props *structpb.Value, columns []*pb.Column) (int64, string) {
	referenceColumnId := props.GetStructValue().Fields[consts.ReferenceColumnId]
	relateColumnId := props.GetStructValue().Fields[consts.RelateColumnId]
	if referenceColumnId == nil || relateColumnId == nil {
		return 0, ""
	}

	tableId := int64(0)
	for _, column := range columns {
		if column.Name == relateColumnId.GetStringValue() {
			tableId = cast.ToInt64(t.GetColumnPropsStringValue(column, consts.RelateTableId))
			break
		}
	}

	return tableId, referenceColumnId.GetStringValue()
}

func (t *tableRepo) GetColumnsByTables(ctx context.Context, tableIds []int64, columnIds []string, isNeedDescription bool) ([]*bo.Columns, error) {
	boColumns := make([]*bo.Columns, 0, len(tableIds))
	columnsMap, err := t.getColumnsMapByTables(ctx, tableIds, columnIds)
	if err != nil {
		return nil, err
	}

	var descMap map[int64]map[string]string
	if isNeedDescription && len(tableIds) == 1 {
		descMap, err = t.getColumnsDescription(ctx, tableIds, columnIds)
	}

	for id, columns := range columnsMap {
		pbColumns, err := covert.ColumnCovert.ToPbColumns(columns)
		if err != nil {
			return nil, err
		}
		// 将描述放入表头
		if len(descMap) > 0 {
			for _, pbColumn := range pbColumns {
				if descMap[id] != nil {
					pbColumn.Description = descMap[id][pbColumn.Name]
				}
			}
		}

		boColumns = append(boColumns, &bo.Columns{TableId: id, Columns: pbColumns})
	}

	return boColumns, nil
}

func (t *tableRepo) GetAppColumnIdsByType(ctx context.Context, appId int64, columnType string) ([]string, error) {
	columns, err := t.GetAppTableColumns(ctx, appId, nil)

	if err != nil {
		return nil, err
	}

	columnIds := make([]string, 0, len(columns))
	for _, column := range columns {
		for _, p := range column.Columns {
			if p.Field != nil && p.Field.Type.String() == columnType {
				columnIds = append(columnIds, p.Name)
			}
		}
	}

	return columnIds, nil
}

func (t *tableRepo) getColumnsMapByTables(ctx context.Context, tableIds []int64, columnIds []string) (map[int64][]*po.TableColumn, error) {
	columnsMap, notInIds, err := t.tableCache.getColumnsByTableIds(ctx, tableIds, columnIds)
	if err != nil {
		return nil, err
	}

	if len(notInIds) > 0 {
		columns := make([]*po.TableColumn, 0, 15)
		err = t.data.mysqlLcGo.Select("id,table_id,column_id,column_type,`schema`").Where("table_id in(?)", tableIds).
			Where("del_flag = ?", 2).
			Find(&columns).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
		needSetCacheIdsMap := make(map[int64]int64, 2)
		for _, column := range columns {
			columnsMap[column.TableId] = append(columnsMap[column.TableId], column)
			needSetCacheIdsMap[column.TableId] = column.TableId
		}

		updateCacheMap := make(map[int64][]*po.TableColumn)
		for _, id := range needSetCacheIdsMap {
			updateCacheMap[id] = columnsMap[id]
		}
		err = t.tableCache.setColumnsBatch(ctx, updateCacheMap)

		for _, id := range needSetCacheIdsMap {
			// 将不需要的字段去除
			columnsMap[id] = t.getLimitColumns(columnsMap[id], columnIds)
		}
	}

	for k, columns := range columnsMap {
		columnsMap[k] = t.sortColumns(columns)
	}

	return columnsMap, nil
}

func (t *tableRepo) sortColumns(columns []*po.TableColumn) []*po.TableColumn {
	sort.Slice(columns, func(i, j int) bool {
		return columns[i].ID < columns[j].ID
	})

	columnsMap := make(map[string]*po.TableColumn, len(columns))
	sortColumns := make([]*po.TableColumn, 0, len(columns))
	for _, column := range columns {
		columnsMap[column.ColumnId] = column
	}
	for _, s := range consts.SummaryColumnIdsSort {
		if column, ok := columnsMap[s]; ok {
			sortColumns = append(sortColumns, column)
		}
	}
	for _, column := range columns {
		if _, ok := consts.SummaryColumnIdsMap[column.ColumnId]; !ok {
			sortColumns = append(sortColumns, column)
		}
	}

	return sortColumns
}

// getColumnsDescription 获取表头描述
func (t *tableRepo) getColumnsDescription(ctx context.Context, tableIds []int64, columnIds []string) (map[int64]map[string]string, error) {
	descMap, notInIds, err := t.tableCache.getColumnsDescriptionByTableIds(ctx, tableIds, columnIds)
	if err != nil {
		return nil, err
	}

	if len(notInIds) > 0 {
		columns := make([]*po.TableColumn, 0, 15)
		err = t.data.mysqlLcGo.Select("id,table_id,column_id,description").Where("table_id in(?)", tableIds).Find(&columns).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
		needSetCacheIdsMap := make(map[int64]int64, 2)
		for _, column := range columns {
			if descMap[column.TableId] == nil {
				descMap[column.TableId] = make(map[string]string)
			}
			descMap[column.TableId][column.ColumnId] = column.Description
			needSetCacheIdsMap[column.TableId] = column.TableId
		}
		updateCacheMap := make(map[int64]map[string]string)
		for _, id := range needSetCacheIdsMap {
			updateCacheMap[id] = descMap[id]
		}
		_ = t.tableCache.setColumnsDescriptionBatch(ctx, updateCacheMap)
	}
	return descMap, nil
}

func (t *tableRepo) getLimitColumns(columns []*po.TableColumn, columnIds []string) []*po.TableColumn {
	// 如果只拿一部分，由于还是要缓存所有的列，所以拿出来后再过滤
	if len(columnIds) > 0 {
		newColumns := make([]*po.TableColumn, 0, len(columnIds))
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

// getColumns 如果columnIds乱传的话还是会缓存击穿
func (t *tableRepo) getColumns(ctx context.Context, tableId int64, columnIds []string) ([]*po.TableColumn, error) {
	columns, err := t.tableCache.getColumns(ctx, tableId, columnIds)
	if err != nil && errors.Cause(err) != redis.Nil {
		return nil, err
	}
	if len(columns) > 0 {
		return columns, nil
	}

	columns = make([]*po.TableColumn, 0, 15)
	err = t.data.mysqlLcGo.Where(&po.TableColumn{TableId: tableId, DelFlag: consts.DeleteFlagNotDel}).Find(&columns).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(columns) == 0 {
		return columns, nil
	}

	_ = t.tableCache.setColumns(ctx, tableId, columns)

	// 如果只拿一部分，由于还是要缓存所有的列，所以拿出来后再过滤
	columns = t.getLimitColumns(columns, columnIds)

	return columns, nil
}

// GetAppTableColumns 获取app下所有表头
func (t *tableRepo) GetAppTableColumns(ctx context.Context, appId int64, columnIds []string) ([]*bo.Columns, error) {
	tables, err := t.GetTables(ctx, appId)
	if err != nil {
		return nil, err
	}

	if len(tables) == 0 {
		return nil, errors.Newf("[GetAppFirstTableColumns] no tables appId:%v", appId)
	}

	tableIds := make([]int64, 0, len(tables))
	for _, table := range tables {
		tableIds = append(tableIds, table.ID)
	}

	return t.GetColumnsByTables(ctx, tableIds, columnIds, false)
}

func (t *tableRepo) GetColumn(ctx context.Context, tableId int64, columnId string, isNeedDescription bool) (*bo.Column, error) {
	column, err := t.tableCache.getColumn(ctx, tableId, columnId)
	if err != nil && errors.Cause(err) != redis.Nil {
		return nil, err
	}
	if column == nil {
		column = &po.TableColumn{}
		err = t.data.mysqlLcGo.Where(&po.TableColumn{TableId: tableId, ColumnId: columnId, DelFlag: consts.DeleteFlagNotDel}).Take(column).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	if isNeedDescription {
		descMap, err := t.getColumnsDescription(ctx, []int64{tableId}, []string{columnId})
		if err != nil {
			return nil, err
		}
		if descMap[tableId] != nil {
			column.Description = descMap[tableId][columnId]
		}
	}

	return covert.ColumnCovert.ToColumn(column)
}

func (t *tableRepo) CreateColumn(ctx context.Context, tc *po.TableColumn, tx ...*gorm.DB) error {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	tc.Creator = ch.UserId

	db := t.data.mysqlLcGo
	if len(tx) >= 1 {
		db = tx[0]
	}

	tc.ID = t.data.snowFlake.Generate().Int64()

	// 如果不是自定义的，因为历史原因，有些高级字段存在，但是不展示，这个时候需要更新，冲突的时候
	//if !strings.HasPrefix(tc.ColumnId, consts.UserDefineColumnPrefix) {
	//	db = db.Clauses(clause.OnConflict{
	//		DoUpdates: clause.AssignmentColumns([]string{"column_type", "schema"}),
	//	})
	//}
	err := db.Create(tc).Error
	if err != nil {
		return err
	}

	if len(tc.Description) > 0 {
		err = t.tableCache.setColumnDescription(ctx, tc.TableId, tc.ColumnId, tc.Description)
		if err != nil {
			return err
		}
	}

	return t.tableCache.deleteColumns(ctx, tc.TableId)
}

func (t *tableRepo) UpdateColumn(ctx context.Context, tc *po.TableColumn, tx ...*gorm.DB) error {
	return t.updateColumn(ctx, tc, tc.ColumnId, nil, tx...)
}

func (t *tableRepo) UpdateColumnWithOldColumnId(ctx context.Context, tc *po.TableColumn, oldColumnId string, tx ...*gorm.DB) error {
	return t.updateColumn(ctx, tc, oldColumnId, nil, tx...)
}

// UpdateColumnAndResetOrgColumnId 需要更新空值的时候要指定下列，要不会被忽略
func (t *tableRepo) UpdateColumnAndResetOrgColumnId(ctx context.Context, tc *po.TableColumn, oldColumnId string, tx ...*gorm.DB) error {
	return t.updateColumn(ctx, tc, oldColumnId, []interface{}{"column_id", "column_type", "schema", "source_org_column_id"}, tx...)
}

func (t *tableRepo) updateColumn(ctx context.Context, column *po.TableColumn, oldColumnId string, updateColumns []interface{}, tx ...*gorm.DB) error {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	column.Updater = ch.UserId

	db := t.data.mysqlLcGo
	if len(tx) >= 1 {
		db = tx[0]
	}

	if len(updateColumns) != 0 {
		db = db.Select(updateColumns[0], updateColumns[1:]...)
	}

	db = db.WithContext(ctx).Where(&po.TableColumn{TableId: column.TableId, ColumnId: oldColumnId}).Updates(column)
	if db.Error != nil {
		return errors.WithStack(db.Error)
	}
	if db.RowsAffected == 0 {
		return nil
	}

	return t.tableCache.deleteColumns(ctx, column.TableId)
}

// ChangeColumnId 将columnId改为另一个名字
func (t *tableRepo) ChangeColumnId(ctx context.Context, oldColumnId string, column *po.TableColumn, tx ...*gorm.DB) error {
	db := t.data.mysqlLcGo
	if len(tx) >= 1 {
		db = tx[0]
	}

	// 将老的columnId更新为新的，并更新schema
	err := db.WithContext(ctx).Where(&po.TableColumn{TableId: column.TableId, ColumnId: oldColumnId}).Updates(column).Error
	if err != nil {
		return errors.WithStack(err)
	}

	// 加入新的column
	return t.tableCache.deleteColumns(ctx, column.TableId)
}

func (t *tableRepo) UpdateColumnDescription(ctx context.Context, tableId int64, columnId, description string) error {
	err := t.data.mysqlLcGo.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Select("description").Where(&po.TableColumn{TableId: tableId, ColumnId: columnId}).Updates(&po.TableColumn{Description: description}).Error
		if err != nil {
			return errors.WithStack(err)
		}

		return t.tableCache.setColumnDescription(ctx, tableId, columnId, description)
	})

	return errors.WithStack(err)
}

// DeleteColumn 删除列
func (t *tableRepo) DeleteColumn(ctx context.Context, tableId int64, columnId string, tx *gorm.DB) error {
	tc := &po.TableColumn{TableId: tableId, ColumnId: columnId}
	err := tx.Where(tc).Delete(tc).Error
	if err != nil {
		return errors.WithStack(err)
	}

	err = t.tableCache.deleteTableDescription(ctx, tableId)
	if err != nil {
		return errors.WithStack(err)
	}

	return t.tableCache.deleteColumns(ctx, tableId)
}

// CheckOrgColumnIdHadUseInOrg 一个组织内，查询下一个字段是否在使用，主要是用于检查组织字段是否使用，在使用的不能删除
func (t *tableRepo) CheckOrgColumnIdHadUseInOrg(ctx context.Context, orgId int64, orgColumnId string) (bool, error) {
	return t.checkHadByCondition(ctx, &po.TableColumn{OrgId: orgId, SourceOrgColumnId: orgColumnId, DelFlag: consts.DeleteFlagNotDel})
}

func (t *tableRepo) CheckOrgColumnIdHadUseInTable(ctx context.Context, tableId int64, orgColumnId string) (bool, error) {
	return t.checkHadByCondition(ctx, &po.TableColumn{TableId: tableId, SourceOrgColumnId: orgColumnId, DelFlag: consts.DeleteFlagNotDel})
}

// checkHadByCondition 根据条件判断是否存在数据
func (t *tableRepo) checkHadByCondition(ctx context.Context, condition *po.TableColumn) (bool, error) {
	c := &po.TableColumn{}
	err := t.data.mysqlLcGo.WithContext(ctx).Where(condition).Take(c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, errors.WithStack(err)
	}

	if c.ID == 0 {
		return false, nil
	}

	return true, nil
}

func (t *tableRepo) GetNewColumnId(category int) string {
	columnId := ""
	size := 4
	switch category {
	case consts.ColumnCategoryOrg:
		size = 3
		columnId = consts.ColumnCategoryOrgPrefix
	case consts.ColumnCategorySummery:
		size = 3
		columnId = consts.ColumnCategorySummeryPrefix
	}

	all := len(consts.ColumnIdRandomKey)
	for i := 0; i < size; i++ {
		columnId += fmt.Sprintf("%c", consts.ColumnIdRandomKey[rand.Intn(all)])
	}

	return columnId
}

// GetColumnCollaboratorRoleIds 获取列的协作人角色配置
func (t *tableRepo) GetColumnCollaboratorRoleIds(column *pb.Column) []string {
	var roleIds []string
	if column.Field != nil && column.Field.Props != nil && t.CheckIsCollaboratorColumn(column.Field.Type.String()) {
		roles := column.Field.Props.Fields[consts.PropertyCollaboratorRoles]
		if roles != nil && roles.GetListValue() != nil && len(roles.GetListValue().Values) > 0 {
			for _, v := range roles.GetListValue().Values {
				roleIds = append(roleIds, v.GetStringValue())
			}
		}
	}
	return roleIds
}

// CheckColumnCollaboratorSwitchOn 检查一下，这个列有没有开启协作人模式
func (t *tableRepo) CheckColumnCollaboratorSwitchOn(column *pb.Column) bool {
	if column.Field != nil && column.Field.Props != nil && t.CheckIsCollaboratorColumn(column.Field.Type.String()) {
		roles := column.Field.Props.Fields[consts.PropertyCollaboratorRoles]
		if roles != nil && roles.GetListValue() != nil && len(roles.GetListValue().Values) > 0 {
			return true
		}
	}
	return false
}

// CheckIsCollaboratorColumn 检查一个列是否是协作人列
func (t *tableRepo) CheckIsCollaboratorColumn(columnType string) bool {
	if _, ok := consts.CollaboratorColumnTypeMap[columnType]; ok {
		return true
	}
	return false
}
