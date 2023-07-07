package biz

import (
	"context"
	"fmt"
	"strings"

	"github.com/star-table/go-table/internal/data/facade/vo/appvo"

	"github.com/star-table/go-table/internal/data/facade/vo"

	"github.com/star-table/go-common/utils/unsafe"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"github.com/star-table/go-common/pkg/encoding"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-common/pkg/middleware/meta"
	"github.com/star-table/go-table/internal/biz/bo"
	"github.com/star-table/go-table/internal/data/consts"
	"github.com/star-table/go-table/internal/data/facade/vo/datacentervo"
	"github.com/star-table/go-table/internal/data/facade/vo/form"
	"github.com/star-table/go-table/internal/data/po"
	pb "github.com/star-table/interface/golang/table/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

type RowRepo interface {
	List(ctx context.Context, req *pb.ListRequest, query *form.QuerySqlReq, memberColumns []*pb.Column, relateColumnIds []string) (*po.Row, error)
	ListRaw(ctx context.Context, req *pb.ListRawRequest, memberColumns []*pb.Column) (*po.Row, error)
	Delete(ctx context.Context, condition *pb.Condition) (int64, error)

	CheckIsAppCollaborator(ctx context.Context, orgId, appId, userId int64) (bool, error)
	GetUserAppCollaboratorColumns(ctx context.Context, orgId, appId, userId int64) ([]*bo.CollaboratorColumn, error)
	GetAppCollaboratorColumns(ctx context.Context, orgId, appId int64) ([]*bo.CollaboratorColumn, error)
	GetDataCollaborators(ctx context.Context, orgId int64, dataIds []int64) ([]*pb.DataCollaborators, error)
	SwitchColumnCollaboratorOn(ctx context.Context, orgId, appId, tableId int64, columnId string) error
	SwitchColumnCollaboratorOff(ctx context.Context, orgId, appId, tableId int64, columnId string) error
	CopyColumnCollaborator(ctx context.Context, orgId, appId, tableId int64, fromColumnId, toColumnId string) error
}

type RowUseCase struct {
	tableRepo      TableRepo
	rowRepo        RowRepo
	lockRepo       LockRepo
	datacenterRepo DatacenterRepo
	appRepo        AppRepo
	log            *log.Helper
}

func NewRowUseCase(tableRepo TableRepo, rowRepo RowRepo, lockRepo LockRepo, datacenterRepo DatacenterRepo, appRepo AppRepo, logger log.Logger) *RowUseCase {
	return &RowUseCase{tableRepo: tableRepo, rowRepo: rowRepo, lockRepo: lockRepo, datacenterRepo: datacenterRepo, appRepo: appRepo, log: log.NewHelper(logger)}
}

func (r *RowUseCase) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	summeryTableId, err := r.tableRepo.GetSummeryTableId(ctx, ch.OrgId)
	if err != nil {
		return nil, err
	}

	if req.TableId == 0 {
		req.TableId = summeryTableId
	}
	columnsMap, err := r.tableRepo.GetColumnsMap(ctx, req.TableId)
	if err != nil {
		return nil, err
	}
	//pbColumns := make([]*pb.Column, 0, len(columnsMap))
	//for _, column := range columnsMap {
	//	pbColumns = append(pbColumns, column.Schema)
	//}

	err = r.exchangeSummaryCondition(ctx, req)
	if err != nil {
		return nil, err
	}

	memberColumns, referenceMemberColumns, err := r.getMemberAndDeptColumns(ctx, columnsMap)
	if err != nil {
		return nil, err
	}

	referenceColumnInfos := r.getReferenceColumnInfo(columnsMap)
	relateColumnIds := make([]string, 0, len(referenceColumnInfos))
	for _, info := range referenceColumnInfos {
		relateColumnIds = append(relateColumnIds, info.RelateColumnId)
	}
	rows, err := r.rowRepo.List(ctx, req, &form.QuerySqlReq{
		OrgId:          ch.OrgId,
		UserId:         ch.UserId,
		Query:          req.Query,
		SummaryTableId: summeryTableId,
		TableId:        req.TableId,
	}, memberColumns, relateColumnIds)
	if err != nil {
		return nil, err
	}

	// 获取关联数据
	var relateBts []byte
	if len(rows.RelateIssueIds) > 0 {
		var relateRows *po.Row
		relateRows, relateBts, err = r.listRelateData(ctx, referenceColumnInfos, rows.RelateIssueIds, rows.OriginRelateIssueIds, referenceMemberColumns)
		if err != nil {
			return nil, err
		}
		for _, id := range relateRows.UserIds {
			if _, ok := rows.UserIdsMap[id]; !ok {
				rows.UserIds = append(rows.UserIds, id)
			}
		}
		for _, id := range relateRows.DeptIds {
			if _, ok := rows.DeptIdsMap[id]; !ok {
				rows.DeptIds = append(rows.DeptIds, id)
			}
		}
	}

	reply := &pb.ListReply{
		UserIds:        rows.UserIds,
		DeptIds:        rows.DeptIds,
		Data:           rows.Buf.Bytes(),
		RelateData:     relateBts,
		LastUpdateTime: rows.MaxUpdateTime.Format(consts.DateFormat),
		Count:          int32(rows.RowCount),
	}

	return reply, nil
}

func (r *RowUseCase) ExchangeSummaryCondition(ctx context.Context, req *pb.ExchangeSummaryConditionRequest) (*pb.ExchangeSummaryConditionReply, error) {
	condition := &vo.LessCondsData{}
	err := encoding.GetJsonCodec().Unmarshal(unsafe.StringBytes(req.Condition), condition)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = r.changeQueryCondition(ctx, req.TableId, condition)
	if err != nil {
		return nil, err
	}

	bts, err := encoding.GetJsonCodec().Marshal(condition)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Condition = unsafe.BytesString(bts)

	return &pb.ExchangeSummaryConditionReply{Condition: req.Condition}, nil
}

// 根据表类型转换到对应的条件
func (r *RowUseCase) exchangeSummaryCondition(ctx context.Context, req *pb.ListRequest) error {
	query := &form.LessIssueListReq{}
	err := encoding.GetJsonCodec().Unmarshal(unsafe.StringBytes(req.Query), query)
	if err != nil {
		return errors.WithStack(err)
	}

	err = r.changeQueryCondition(ctx, req.TableId, query.Condition)
	if err != nil {
		return err
	}

	bts, err := encoding.GetJsonCodec().Marshal(query)
	if err != nil {
		return errors.WithStack(err)
	}
	req.Query = unsafe.BytesString(bts)

	return nil
}

// 如果是项目汇总表，要将条件的tableId改为这个appId，汇总所有数据
func (r *RowUseCase) changeQueryCondition(ctx context.Context, tableId int64, condition *vo.LessCondsData) error {
	tableInfo, err := r.tableRepo.GetTable(ctx, tableId)
	if err != nil {
		return err
	}

	switch tableInfo.SummeryFlag {
	case consts.SummaryFlagApp:
		// 项目汇总表
		r.changeSummaryTableIdToAppId(condition, tableInfo.ID, []int64{tableInfo.AppId})
	case consts.SummaryFlagFolder:
		// 文件夹汇总表
		ch := meta.GetCommonHeaderFromCtx(ctx)
		list, err := r.appRepo.GetAppList(ctx, &appvo.AppListReqVo{
			OrgId:    ch.OrgId,
			Type:     consts.LcAppTypeForPolaris,
			ParentId: tableInfo.AppId,
		})
		if err != nil {
			return err
		}
		if len(list.Data) > 0 {
			appIds := make([]int64, 0, len(list.Data))
			for _, datum := range list.Data {
				appIds = append(appIds, datum.Id)
			}
			r.changeSummaryTableIdToAppId(condition, tableInfo.ID, appIds)
		}
	case consts.SummaryFlagAll:
		// 如果是全部任务汇总表直接替换成orgId的条件，相当于不需要tableId条件了
		ch := meta.GetCommonHeaderFromCtx(ctx)
		replaceCondition := &vo.LessCondsData{Type: pb.ConditionType_equal.String(), Column: consts.ColumnIdOrgId, Value: ch.OrgId}
		r.changeSummaryTableIdCondition(condition, replaceCondition, tableId)
	}

	return nil
}

func (r *RowUseCase) changeSummaryTableIdToAppId(condition *vo.LessCondsData, tableId int64, appIds []int64) {
	replaceCondition := &vo.LessCondsData{Type: pb.ConditionType_in.String(), Column: consts.ColumnIdAppId, Values: appIds}
	r.changeSummaryTableIdCondition(condition, replaceCondition, tableId)
}

func (r *RowUseCase) changeSummaryTableIdCondition(condition, replaceCondition *vo.LessCondsData, tableId int64) {
	if condition == nil {
		return
	}

	if condition.Type == pb.ConditionType_and.String() || condition.Type == pb.ConditionType_or.String() {
		for _, cond := range condition.Conds {
			r.changeSummaryTableIdCondition(cond, replaceCondition, tableId)
		}
	} else {
		if condition.Column == consts.ColumnIdTableId && cast.ToInt64(condition.Value) == tableId {
			condition.Type = replaceCondition.Type
			condition.Column = replaceCondition.Column
			condition.Value = replaceCondition.Value
			condition.Values = replaceCondition.Values
		}
	}
}

func (r *RowUseCase) listRelateData(ctx context.Context, referenceColumnInfos []*bo.ReferenceColumnInfo, relateDataIds map[int64]struct{},
	idToRelateIdsMap map[string]map[string][]int64, memberColumns []*pb.Column) (*po.Row, []byte, error) {

	if len(referenceColumnInfos) > 0 && len(relateDataIds) > 0 {
		ch := meta.GetCommonHeaderFromCtx(ctx)
		issueIds := make([]string, 0, len(relateDataIds))
		for id := range relateDataIds {
			idStr := cast.ToString(id)
			if len(idStr) < 15 {
				issueIds = append(issueIds, idStr)
			}
		}

		filterColumnIds := make([]string, 0, len(referenceColumnInfos)+3)
		filterColumnIds = append(filterColumnIds, consts.ColumnId, datacentervo.WrapperJsonColumn(consts.ColumnIdIssueId), datacentervo.WrapperJsonColumnAlias(consts.ColumnIdTitle, consts.ColumnIdTitle))
		for _, info := range referenceColumnInfos {
			// 排除下关联列，只有引用列需要数据，关联列默认拿的是title数据
			if !info.IsRelate {
				filterColumnIds = append(filterColumnIds, datacentervo.WrapperJsonColumnAlias(info.ReferenceColumnId, info.OriginColumnId))
			}
		}

		rows, err := r.rowRepo.ListRaw(ctx, &pb.ListRawRequest{
			FilterColumns: filterColumnIds,
			Condition: &pb.Condition{Type: pb.ConditionType_and, Conditions: []*pb.Condition{
				{Type: pb.ConditionType_equal, Column: datacentervo.WrapperJsonColumn(consts.ColumnIdOrgId), Value: fmt.Sprintf("[%d]", ch.OrgId)},
				{Type: pb.ConditionType_in, Column: datacentervo.WrapperJsonColumn(consts.ColumnIdIssueId), Value: fmt.Sprintf("[%s]", strings.Join(issueIds, ","))},
			}},
		}, memberColumns)

		if err != nil {
			return nil, nil, err
		}

		list := make([]map[string]interface{}, 0, len(relateDataIds))
		err = encoding.GetJsonCodec().Unmarshal(rows.Buf.Bytes(), &list)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}

		result := r.aggCalculate(list, referenceColumnInfos, idToRelateIdsMap)
		bts, err := encoding.GetJsonCodec().Marshal(result)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}

		return rows, bts, nil
	}

	return &po.Row{}, nil, nil
}

func (r *RowUseCase) ListRaw(ctx context.Context, req *pb.ListRawRequest) (*pb.ListRawReply, error) {
	rows, err := r.rowRepo.ListRaw(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	reply := &pb.ListRawReply{Data: rows.Buf.Bytes()}

	return reply, nil
}

func (r *RowUseCase) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	count, err := r.rowRepo.Delete(ctx, req.Condition)
	return &pb.DeleteReply{Count: count}, err
}

func (r *RowUseCase) CheckIsAppCollaborator(ctx context.Context, req *pb.CheckIsAppCollaboratorRequest) (*pb.CheckIsAppCollaboratorReply, error) {
	reply := &pb.CheckIsAppCollaboratorReply{}
	ch := meta.GetCommonHeaderFromCtx(ctx)
	result, err := r.rowRepo.CheckIsAppCollaborator(ctx, ch.OrgId, req.AppId, req.UserId)
	if err != nil {
		return nil, err
	}
	reply.Result = result
	return reply, nil
}

func (r *RowUseCase) GetUserAppCollaboratorRoles(ctx context.Context, req *pb.GetUserAppCollaboratorRolesRequest) (*pb.GetUserAppCollaboratorRolesReply, error) {
	reply := &pb.GetUserAppCollaboratorRolesReply{}
	ch := meta.GetCommonHeaderFromCtx(ctx)
	collaboratorColumnIds, err := r.rowRepo.GetUserAppCollaboratorColumns(ctx, ch.OrgId, req.AppId, req.UserId)
	if err != nil {
		return nil, err
	}

	if len(collaboratorColumnIds) == 0 {
		return reply, nil
	}

	columns, err := r.tableRepo.GetAppTableColumns(ctx, req.AppId, nil)
	if err != nil {
		return nil, err
	}
	allColumnMap := make(map[int64]map[string]*pb.Column) // tableId -> columnId -> column
	for _, cs := range columns {
		for _, c := range cs.Columns {
			if m, ok := allColumnMap[cs.TableId]; ok {
				m[c.Name] = c
			} else {
				m = make(map[string]*pb.Column)
				m[c.Name] = c
				allColumnMap[cs.TableId] = m
			}
		}
	}

	columnMap := make(map[string]*pb.Column)
	for _, cc := range collaboratorColumnIds {
		if m, ok := allColumnMap[cc.TableId]; ok {
			if c, ok1 := m[cc.ColumnId]; ok1 {
				key := fmt.Sprintf("%d_%s", cc.TableId, cc.ColumnId)
				columnMap[key] = c // 我是协作人的涉及到的列（注意不同表的列名可能会重复）
			}
		}
	}

	roleIdMap := make(map[string]struct{})
	for _, c := range columnMap {
		roleIds := r.tableRepo.GetColumnCollaboratorRoleIds(c)
		for _, rId := range roleIds {
			roleIdMap[rId] = struct{}{}
		}
	}
	for rId, _ := range roleIdMap {
		reply.RoleIds = append(reply.RoleIds, rId)
	}
	return reply, nil
}

func (r *RowUseCase) GetAppCollaboratorRoles(ctx context.Context, req *pb.GetAppCollaboratorRolesRequest) (*pb.GetAppCollaboratorRolesReply, error) {
	reply := &pb.GetAppCollaboratorRolesReply{}
	ch := meta.GetCommonHeaderFromCtx(ctx)
	collaboratorColumnIds, err := r.rowRepo.GetAppCollaboratorColumns(ctx, ch.OrgId, req.AppId)
	if err != nil {
		return nil, err
	}

	if len(collaboratorColumnIds) == 0 {
		return reply, nil
	}

	columns, err := r.tableRepo.GetAppTableColumns(ctx, req.AppId, nil)
	if err != nil {
		return nil, err
	}
	allColumnMap := make(map[int64]map[string]*bo.ColumnWithCollaboratorRoles) // tableId -> columnId -> column
	for _, cs := range columns {
		for _, c := range cs.Columns {
			if m, ok := allColumnMap[cs.TableId]; ok {
				m[c.Name] = &bo.ColumnWithCollaboratorRoles{Column: c}
			} else {
				m = make(map[string]*bo.ColumnWithCollaboratorRoles)
				m[c.Name] = &bo.ColumnWithCollaboratorRoles{Column: c}
				allColumnMap[cs.TableId] = m
			}
		}
	}

	columnMap := make(map[string]map[string]*bo.ColumnWithCollaboratorRoles) // uid -> key(tableId+columnId) -> column
	for _, cc := range collaboratorColumnIds {
		if m, ok := allColumnMap[cc.TableId]; ok {
			if c, ok1 := m[cc.ColumnId]; ok1 {
				key := fmt.Sprintf("%d_%s", cc.TableId, cc.ColumnId)
				if mm, ok2 := columnMap[cc.Id]; ok2 {
					mm[key] = c
				} else {
					mm = make(map[string]*bo.ColumnWithCollaboratorRoles)
					mm[key] = c
					columnMap[cc.Id] = mm
				}
			}
		}
	}

	for id, cm := range columnMap {
		roleIdMap := make(map[string]struct{})
		for _, c := range cm {
			if !c.GotRoles {
				c.Roles = r.tableRepo.GetColumnCollaboratorRoleIds(c.Column)
				c.GotRoles = true
			}
			for _, rId := range c.Roles {
				roleIdMap[rId] = struct{}{}
			}
		}
		if len(roleIdMap) > 0 {
			roles := &pb.CollaboratorRole{
				Id: id,
			}
			for rId, _ := range roleIdMap {
				roles.RoleIds = append(roles.RoleIds, rId)
			}
			reply.CollaboratorRoles = append(reply.CollaboratorRoles, roles)
		}
	}
	return reply, nil
}

func (r *RowUseCase) GetDataCollaborators(ctx context.Context, req *pb.GetDataCollaboratorsRequest) (*pb.GetDataCollaboratorsReply, error) {
	reply := &pb.GetDataCollaboratorsReply{}
	ch := meta.GetCommonHeaderFromCtx(ctx)
	collaborators, err := r.rowRepo.GetDataCollaborators(ctx, ch.OrgId, req.DataIds)
	if err != nil {
		return nil, err
	}
	reply.Collaborators = collaborators
	return reply, nil
}

func (r *RowUseCase) SwitchColumnCollaboratorOn(ctx context.Context, orgId, appId, tableId int64, columnId string) error {
	return r.rowRepo.SwitchColumnCollaboratorOn(ctx, orgId, appId, tableId, columnId)
}

func (r *RowUseCase) SwitchColumnCollaboratorOff(ctx context.Context, orgId, appId, tableId int64, columnId string) error {
	return r.rowRepo.SwitchColumnCollaboratorOff(ctx, orgId, appId, tableId, columnId)
}

func (r *RowUseCase) CopyColumnCollaborator(ctx context.Context, orgId, appId, tableId int64, fromColumnId, toColumnId string) error {
	return r.rowRepo.CopyColumnCollaborator(ctx, orgId, appId, tableId, fromColumnId, toColumnId)
}

//func (r *RowUseCase) getMemberAndDeptIds(ctx context.Context, list []map[string]interface{}, columns []*bo.Column) ([]int64, []int64, error) {
//	memberColumns, err := r.getMemberAndDeptColumns(ctx, columns)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	if len(memberColumns) > 0 {
//		userIdsMap := make(map[int64]struct{}, 100)
//		deptIdsMap := make(map[int64]struct{}, 30)
//		for _, m := range list {
//			for _, column := range memberColumns {
//				if s, ok := m[column.Name]; ok {
//					if slice, ok := s.([]interface{}); ok {
//						for _, s2 := range slice {
//							s3 := strings.Replace(cast.ToString(s2), consts.UserPrefix, "", 1)
//							s3 = strings.Replace(s3, consts.DeptPrefix, "", 1)
//							r.addToIdMap(userIdsMap, deptIdsMap, cast.ToInt64(s3), column)
//						}
//					} else {
//						r.addToIdMap(userIdsMap, deptIdsMap, cast.ToInt64(s), column)
//					}
//				}
//			}
//		}
//
//		userIds := make([]int64, 0, len(userIdsMap))
//		deptIds := make([]int64, 0, len(deptIdsMap))
//		for id := range userIdsMap {
//			userIds = append(userIds, id)
//		}
//		for id := range deptIdsMap {
//			deptIds = append(deptIds, id)
//		}
//
//		return userIds, deptIds, nil
//	}
//
//	return nil, nil, nil
//}

func (r *RowUseCase) addToIdMap(userIdsMap, deptIdsMap map[int64]struct{}, id int64, column *pb.Column) {
	if id != 0 {
		if column.Field.Type == pb.ColumnType_member {
			userIdsMap[id] = struct{}{}
		} else {
			deptIdsMap[id] = struct{}{}
		}
	}
}

func (r *RowUseCase) CreateRows(ctx context.Context, req *pb.CreateRowsRequest) (*pb.CreateRowsReply, error) {
	if len(req.GetRows()) == 0 {
		return &pb.CreateRowsReply{}, nil
	}

	// 拿表头
	columns, err := r.tableRepo.GetColumnsMap(ctx, req.TableId)
	if err != nil {
		return nil, err
	}

	// 数据校验
	err = r.validateRows(ctx, req.Rows, columns, req.IsImport)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (r *RowUseCase) MoveRow(ctx context.Context, req *pb.MoveRowRequest) (*pb.MoveRowReply, error) {

	return nil, nil
}
func (r *RowUseCase) CopyRow(ctx context.Context, req *pb.CopyRowRequest) (*pb.CopyRowReply, error) {
	return nil, nil
}
func (r *RowUseCase) DeleteRow(ctx context.Context, req *pb.DeleteRowRequest) (*pb.DeleteRowReply, error) {
	return nil, nil
}

func (r *RowUseCase) validateRows(ctx context.Context, rows []*structpb.Struct, columns map[string]*bo.Column, isImport bool) error {
	var ok bool
	var column *bo.Column
	for _, row := range rows {
		// TODO: 校验前回调
		for k, v := range row.GetFields() {
			// 跳过一些字段类型的校验(含公共字段)
			if _, ok = consts.ExcludedValidateColumns[k]; ok {
				continue
			}

			// 获取表头，没有表头则跳过
			if column, ok = columns[k]; !ok {
				continue
			}

			// 校验
			err := r._validateRow(ctx, k, v, column)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *RowUseCase) _validateRow(ctx context.Context, key string, value *structpb.Value, column *bo.Column) error {
	return nil
}

func (r *RowUseCase) _validatePreHook() {
}

func (r *RowUseCase) _validateFailHook() {
}

func (r *RowUseCase) _validateSuccHook() {
}

func (r *RowUseCase) _validateAfterHook() {
}

// updateRowsGroupSelectType 修改需要转换type的row
func (r *RowUseCase) updateRowsGroupSelectType(ctx context.Context, tableId int64, columnId string, changeType map[int64][]interface{}) error {
	if len(changeType) > 0 {
		ch := meta.GetCommonHeaderFromCtx(ctx)
		summeryTableId, err := r.tableRepo.GetSummeryTableId(ctx, ch.OrgId)
		if err != nil {
			return err
		}
		tableName := r.datacenterRepo.WrapperTableName(ch.OrgId, summeryTableId)
		executors := make([]*datacentervo.Executor, 0, len(changeType))
		for t, ids := range changeType {
			condition := r.getAndConditions(r.getTableCondition(tableId),
				r.getConditionWithValues(datacentervo.ConditionIn, columnId, ids))
			executor := r.getSingleColumnUpdateExecutor(tableName, condition, consts.GetGroupSelectType(columnId), t)
			executors = append(executors, executor)
		}
		_, err = r.datacenterRepo.DatacenterExecutors(ctx, consts.DataSource, consts.DataBase, executors)

		return err
	}

	return nil
}

// deleteRowsByTableId 根据tableId删除数据或者标记数据
//func (r *RowUseCase) deleteRowsByTableId(ctx context.Context, tableId int64) error {
//	return r.executeUpdateSingleColumnValue(ctx, r.getTableCondition(tableId), consts.GetDelFlag(), consts.DeleteFlagDel)
//}

// deleteColumnData 删除某一个列的数据
func (r *RowUseCase) deleteColumnData(ctx context.Context, orgId, tableId int64, columnId string) error {
	var dataValue string
	if _, ok := consts.DeleteColumnNotDeleteData[columnId]; !ok {
		dataValue = fmt.Sprintf(`data #- '{%s}'`, columnId)
	}
	if dataValue == "" {
		return nil
	}

	cond := r.getAndConditions(r.getCondition(datacentervo.ConditionEqual, consts.ColumnIdOrgId, orgId))
	if tableId != 0 {
		cond.Conds = append(cond.Conds, r.getTableCondition(tableId))
	}

	return r.executeUpdateSingleColumnValue(ctx, cond, consts.DataField, dataValue, true)
}

// 改列名后把数据拷贝
func (r *RowUseCase) copyColumnData(ctx context.Context, orgId, tableId int64, fromColumnId, toColumnId string) error {
	updateColumnId := fmt.Sprintf("data.%s", toColumnId)
	updateValue := fmt.Sprintf("COALESCE(" + datacentervo.WrapperJsonColumn(fromColumnId) + ", 'null')")

	cond := r.getAndConditions(r.getCondition(datacentervo.ConditionEqual, consts.ColumnIdOrgId, orgId))
	if tableId != 0 {
		cond.Conds = append(cond.Conds, r.getTableCondition(tableId))
	}

	return r.executeUpdateSingleColumnValue(ctx, cond, updateColumnId, updateValue, true)
}

// 移动数据，将老的列数据删掉
func (r *RowUseCase) moveColumnData(ctx context.Context, tableId int64, fromColumnId, toColumnId string, isCollaboratorType bool) error {
	updateValue := fmt.Sprintf("data #- '{%s}'", fromColumnId)
	if isCollaboratorType {
		updateValue += fmt.Sprintf(" #- '{collaborators,%s}'", fromColumnId)
	}
	updateValue += fmt.Sprintf(" || jsonb_build_object('%s', data -> '%s'", toColumnId, fromColumnId)
	if isCollaboratorType {
		updateValue += fmt.Sprintf(", '%s', data->'%s'", toColumnId, fromColumnId)
	}
	updateValue += ")"

	return r.executeUpdateSingleColumnValue(ctx, r.getTableCondition(tableId), consts.DataField, updateValue, true)
}

func (r *RowUseCase) RecycleAttachment(ctx context.Context, req *pb.RecycleAttachmentRequest) (*pb.RecycleAttachmentReply, error) {
	err := r.setAttachmentRecycleFlag(ctx, req.AppId, req.IssueIds, req.ResourceIds, consts.RecycleFlagYes)
	return &pb.RecycleAttachmentReply{}, err
}

func (r *RowUseCase) RecoverAttachment(ctx context.Context, req *pb.RecoverAttachmentRequest) (*pb.RecoverAttachmentReply, error) {
	err := r.setAttachmentRecycleFlag(ctx, req.AppId, req.IssueIds, req.ResourceIds, consts.RecycleFlagNo)
	return &pb.RecoverAttachmentReply{}, err
}

// RecycleAttachment 将附件放入回收站或者从回收站取出来
func (r *RowUseCase) setAttachmentRecycleFlag(ctx context.Context, appId int64, issueIds []int64, resourceIds []int64, flag int) error {
	columnIds, err := r.tableRepo.GetAppColumnIdsByType(ctx, appId, pb.ColumnType_document.String())
	if err != nil {
		return err
	}
	if len(columnIds) == 0 {
		return nil
	}

	dataValue := r.getAttachmentRecycleJson(resourceIds, columnIds, flag)
	condition := r.getCondition(datacentervo.ConditionIn, consts.ColumnIdIssueId, issueIds)
	return r.executeUpdateSingleColumnValue(ctx, condition, consts.DataField, dataValue, true)
}

// getAttachmentRecycleJson 设置附件回收站flag，俄罗斯套娃，一个套一个
func (r *RowUseCase) getAttachmentRecycleJson(resourceIds []int64, columnIds []string, recycleFlag int) string {
	updateJsons := make([]string, 0, len(columnIds))
	defaultValue := `jsonb_set(%s, '{%s,%d,recycleFlag}', '%d', false)`
	for _, resourceId := range resourceIds {
		dataValue := defaultValue
		if len(columnIds) == 1 {
			dataValue = fmt.Sprintf(defaultValue, consts.DataField, columnIds[0], resourceId, recycleFlag)
		} else {
			for i, columnId := range columnIds {
				if i == len(columnIds)-1 {
					dataValue = fmt.Sprintf(dataValue, fmt.Sprintf(defaultValue, consts.DataField, columnId, resourceId, recycleFlag))
				} else if i == 0 {
					dataValue = fmt.Sprintf(defaultValue, "%s", columnId, resourceId, recycleFlag)
				} else {
					dataValue = fmt.Sprintf(defaultValue, dataValue, columnId, resourceId, recycleFlag)
				}
			}
		}
		updateJsons = append(updateJsons, dataValue)
	}

	updateJson := updateJsons[0]
	for i := 1; i < len(updateJsons); i++ {
		updateJson = strings.Replace(updateJson, consts.DataField, updateJsons[i], 1)
	}

	return updateJson
}

// getTableCondition 获取表id的查询条件
func (r *RowUseCase) getTableCondition(tableId int64) *datacentervo.LessCondsData {
	return r.getCondition(datacentervo.ConditionEqual, consts.ColumnIdTableId, cast.ToString(tableId))
}

// getCondition 获取单值查询的条件
func (r *RowUseCase) getCondition(condType, column string, value interface{}) *datacentervo.LessCondsData {
	return &datacentervo.LessCondsData{
		Type:   condType,
		Column: datacentervo.WrapperJsonColumn(column),
		Value:  value,
	}
}

// getConditionWithValues 获取多值查询的条件
func (r *RowUseCase) getConditionWithValues(condType, column string, values interface{}) *datacentervo.LessCondsData {
	return &datacentervo.LessCondsData{
		Type:   condType,
		Column: datacentervo.WrapperJsonColumn(column),
		Values: values,
	}
}

// getAndConditions 获取and的条件
func (r *RowUseCase) getAndConditions(conds ...*datacentervo.LessCondsData) *datacentervo.LessCondsData {
	return &datacentervo.LessCondsData{
		Type:  datacentervo.ConditionAnd,
		Conds: conds,
	}
}

// getSummeryTableName 获取汇总表的名称
func (r *RowUseCase) getSummeryTableName(ctx context.Context) (string, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	summeryTableId, err := r.tableRepo.GetSummeryTableId(ctx, ch.OrgId)
	if err != nil {
		return "", err
	}

	return r.datacenterRepo.WrapperTableName(ch.OrgId, summeryTableId), nil
}

// executeUpdateSingleColumnValue 执行某一列的更新操作
func (r *RowUseCase) executeUpdateSingleColumnValue(ctx context.Context, condition *datacentervo.LessCondsData,
	columnName string, columnValue interface{}, withoutPretreat ...bool) error {

	tableName, err := r.getSummeryTableName(ctx)
	if err != nil {
		return err
	}

	executor := r.getSingleColumnUpdateExecutor(tableName, condition, columnName, columnValue, withoutPretreat...)
	_, err = r.datacenterRepo.DatacenterExecutor(ctx, consts.DataSource, consts.DataBase, executor)
	return errors.WithStack(err)
}

// getSingleColumnUpdateExecutor 获取一个更新列的执行model
func (r *RowUseCase) getSingleColumnUpdateExecutor(tableName string, condition *datacentervo.LessCondsData,
	columnName string, columnValue interface{}, withoutPretreat ...bool) *datacentervo.Executor {
	pretreat := false
	if len(withoutPretreat) > 0 {
		pretreat = true
	}

	return &datacentervo.Executor{
		Type:      datacentervo.ExecuteTypeUpdate,
		Table:     datacentervo.NewTable(tableName),
		Condition: condition,
		Sets: []datacentervo.Set{
			{
				Column:          columnName,
				Type:            datacentervo.SetTypeJson,
				Value:           columnValue,
				WithoutPretreat: pretreat,
			},
		},
	}
}

func (r *RowUseCase) getMemberAndDeptColumns(ctx context.Context, columns map[string]*bo.Column) ([]*pb.Column, []*pb.Column, error) {
	normalList := make([]*pb.Column, 0, 5)
	referenceList := make([]*pb.Column, 0, 5)
	refColumns := make([]*pb.Column, 0, 1)
	for _, column := range columns {
		if column.ColumnType == pb.ColumnType_dept.String() || column.ColumnType == pb.ColumnType_member.String() {
			normalList = append(normalList, column.Schema)
		} else if column.ColumnType == pb.ColumnType_conditionRef.String() || column.ColumnType == pb.ColumnType_reference.String() {
			refColumns = append(refColumns, column.Schema)
		}
	}

	if len(refColumns) > 0 {
		boColumnsMap, err := r.tableRepo.GetRefColumns(ctx, refColumns)
		if err != nil {
			return nil, nil, err
		}
		// 如果引用的列是部门或者成员
		for columnId, refColumn := range boColumnsMap {
			if refColumn.Schema.Field.Type == pb.ColumnType_dept || refColumn.Schema.Field.Type == pb.ColumnType_member {
				if columns[columnId].Schema.Field.Type.String() == pb.ColumnType_reference.String() {
					referenceList = append(referenceList, columns[columnId].Schema)
				} else {
					normalList = append(normalList, columns[columnId].Schema)
				}
			}
		}
	}

	return normalList, referenceList, nil
}

// getReferenceColumnInfo 收集下需要获取的引用相关列信息和关联列信息
func (r *RowUseCase) getReferenceColumnInfo(columns map[string]*bo.Column) []*bo.ReferenceColumnInfo {
	columnInfos := make([]*bo.ReferenceColumnInfo, 0, 1)
	for _, column := range columns {
		if column.ColumnType == pb.ColumnType_reference.String() {
			_, referenceColumnId := r.tableRepo.GetColumnRefTableInfo(column.Schema, nil)
			relateColumnId := r.tableRepo.GetColumnPropsStringValue(column.Schema, consts.RelateColumnId)
			aggFunc := r.tableRepo.GetColumnPropsStringValue(column.Schema, consts.AggFunc)
			if referenceColumnId != "" && relateColumnId != "" {
				columnInfos = append(columnInfos, &bo.ReferenceColumnInfo{
					RelateColumnId:    relateColumnId,
					ReferenceColumnId: referenceColumnId,
					OriginColumnId:    column.Schema.Name,
					AggFunc:           aggFunc,
				})
			}
		} else if column.ColumnType == pb.ColumnType_relating.String() || column.ColumnType == pb.ColumnType_singleRelating.String() {
			columnInfos = append(columnInfos, &bo.ReferenceColumnInfo{
				RelateColumnId: column.Schema.Name,
				OriginColumnId: column.Schema.Name,
				IsRelate:       true,
			})
		}
	}

	return columnInfos
}
