package consts

import "fmt"

const (
	DataField      = "data"
	groupSelectExt = "Type"

	UserPrefix = "U_"
	DeptPrefix = "D_"

	DbTypePrefix = "postgres_"

	FilterColumns = "filterColumns"

	MaxFloat64ToInt64 = float64(92233720368547758)
)

const (
	Mark         = '"'
	Null         = "null"
	FalseStr     = "false"
	TrueStr      = "true"
	ArrCharLeft  = '['
	ArrCharRight = ']'
	ObjCharLeft  = '{'
	ObjCharRight = '}'
	SpiltChar    = ','
	EqualChar    = ':'
)

const (
	ColumnId                    = "id"
	ColumnDataId                = "dataId"
	ColumnIdOrgId               = "orgId"
	ColumnIdAppId               = "appId"
	ColumnIdProjectId           = "projectId"
	ColumnIdTableId             = "tableId"
	ColumnIdRecycleFlag         = "recycleFlag"
	ColumnIdDelFlag             = "delFlag"
	ColumnIdIssueId             = "issueId"
	ColumnIdCollaborators       = "collaborators"
	ColumnIdWorkHour            = "workHour" // 工时字段
	ColumnIdAuditorIds          = "auditorIds"
	ColumnIdCode                = "code"
	ColumnIdPath                = "path"
	ColumnIdFollowerIds         = "followerIds"
	ColumnIdIssueStatus         = "issueStatus"
	ColumnIdIterationId         = "iterationId"
	ColumnIdOwnerId             = "ownerId"
	ColumnIdParentId            = "parentId"
	ColumnIdPlanEndTime         = "planEndTime"
	ColumnIdPlanStartTime       = "planStartTime"
	ColumnIdCreateTime          = "createTime"
	ColumnIdUpdateTime          = "updateTime"
	ColumnIdProjectObjectTypeId = "projectObjectTypeId"
	ColumnIdRemark              = "remark"
	ColumnIdTitle               = "title"
	ColumnIdOrder               = "order"
	ColumnIdCreator             = "creator"
	ColumnIdUpdator             = "updator"
	ColumnIdData                = "data"
	ColumnIdRelating            = "relating"
	ColumnIdBaRelating          = "baRelating"
)

var NotJsonColumnMap = map[string]struct{}{
	ColumnId:              {},
	ColumnIdOrgId:         {},
	ColumnIdRecycleFlag:   {},
	ColumnIdCreator:       {},
	ColumnIdUpdator:       {},
	ColumnIdCreateTime:    {},
	ColumnIdUpdateTime:    {},
	ColumnIdAppId:         {},
	ColumnIdProjectId:     {},
	ColumnIdTableId:       {},
	ColumnIdPath:          {},
	ColumnIdParentId:      {},
	ColumnIdIssueId:       {},
	ColumnIdCollaborators: {},
	ColumnIdOrder:         {},
	ColumnIdCode:          {},
	ColumnIdData:          {},
}

// SummaryColumnIdsMap 汇总表字段名
var SummaryColumnIdsMap = map[string]struct{}{
	ColumnIdTitle:               {},
	ColumnIdCode:                {},
	ColumnIdOwnerId:             {},
	ColumnIdIssueStatus:         {},
	ColumnIdPlanStartTime:       {},
	ColumnIdPlanEndTime:         {},
	ColumnIdRemark:              {},
	ColumnIdFollowerIds:         {},
	ColumnIdAuditorIds:          {},
	ColumnIdProjectObjectTypeId: {},
	ColumnIdIterationId:         {},
}

var SummaryColumnIdsSort = []string{
	ColumnIdTitle,
	ColumnIdCode,
	ColumnIdOwnerId,
	ColumnIdIssueStatus,
	ColumnIdPlanStartTime,
	ColumnIdPlanEndTime,
	ColumnIdRemark,
	ColumnIdFollowerIds,
	ColumnIdAuditorIds,
	ColumnIdProjectObjectTypeId,
	ColumnIdIterationId,
}

// GetGroupSelectType groupSelect类型的category路径，用于更新
func GetGroupSelectType(columnId string) string {
	return fmt.Sprintf("%s.%s%s", DataField, columnId, groupSelectExt)
}

// GetRecycleFlag 删除标识
func GetRecycleFlag() string {
	return fmt.Sprintf("%s.%s", DataField, ColumnIdRecycleFlag)
}

func GetDelFlag() string {
	return fmt.Sprintf("%s.%s", DataField, ColumnIdDelFlag)
}

func GetCollaborators(columnId string) string {
	return fmt.Sprintf("%s.%s.%s", DataField, ColumnIdCollaborators, columnId)
}
