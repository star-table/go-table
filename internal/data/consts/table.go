package consts

import (
	"github.com/star-table/go-common/pkg/encoding"
	pb "github.com/star-table/interface/golang/table/v1"
)

const (
	DeleteFlagDel    = 1
	DeleteFlagNotDel = 2

	RecycleFlagYes = 1
	RecycleFlagNo  = 2
)

const (
	DataSource = 2 // pg使用的id
	DataBase   = 1 // pg使用的id
)

const (
	ColumnCategoryNormal = iota
	ColumnCategoryOrg
	ColumnCategorySummery
)

const (
	ColumnCategoryOrgPrefix     = "_"
	ColumnCategorySummeryPrefix = "-"
	ColumnIdRandomKey           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

const (
	SummaryFlagAll    = 1 // 全部任务表
	SummaryFlagNormal = 2 // 普通表
	SummaryFlagApp    = 3 // 项目汇总表
	SummaryFlagFolder = 4 // 文件夹汇总表
)

const (
	UserDefineColumnPrefix   = "_field"
	ColumnRelation           = "_field_relation"
	ColumnBeforeAfter        = "_field_before_after"
	RelationField            = "link_to"
	BeRelatedField           = "link_from"
	BeforeField              = "before_link_to"
	AfterField               = "after_link_from"
	TitleField               = "title"
	IssueStatusField         = "issueStatus"
	ProjectObjectTypeIdField = "projectObjectTypeId"
	RelateTableId            = "relateTableId"
	RelateAppId              = "relateAppId"
	Reference                = "reference"
	ReferenceColumnId        = "referenceColumnId" // 真正引用的列
	RelateColumnId           = "relateColumnId"    // 关联列
	AggFunc                  = "aggFunc"           // 计算函数
)

const (
	PropertyCollaboratorRoles = "collaboratorRoles" // 协作人标识
)

var NoNeedColumnIdsMap = map[string]struct{}{
	"parentId": {},
}

var Document = &pb.Column{}
var DocumentBts = []byte(`{
	"name": "document", 
	"field": {
		"type": "document", 
		"customType":"",
		"dataType": "CUSTOM",
		"props": {
			"checked": true,
			"disabled": false, 
			"hide": false,
			"multiple": false, 
			"required": false
		}
	}, 
	"label": "附件", 
	"editable": true, 
	"writable": true
}`)

var CanNotCopyAndOrgColumns = map[string]struct{}{
	pb.ColumnType_workHour.String():   {},
	pb.ColumnType_relating.String():   {},
	pb.ColumnType_baRelating.String(): {},
}

var CollaboratorColumnTypeMap = map[string]struct{}{
	pb.ColumnType_workHour.String(): {},
	pb.ColumnType_member.String():   {},
	pb.ColumnType_dept.String():     {},
}

var ExcludedValidateColumns = map[string]interface{}{
	"issueStatus":         struct{}{},
	"projectObjectTypeId": struct{}{},
	"iterationId":         struct{}{},
	"creator":             struct{}{},
	"createTime":          struct{}{},
	"updator":             struct{}{},
	"updateTime":          struct{}{},
	"status":              struct{}{},
	"parentId":            struct{}{},
}

// DeleteColumnNotDeleteData 有些列删除后，不能删除数据，因为这些数据是默认存在的，只是表头删减
var DeleteColumnNotDeleteData = map[string]struct{}{
	"creator":    {},
	"createTime": {},
	"updator":    {},
	"updateTime": {},
	"status":     {},
	"parentId":   {},
	"relating":   {},
	"baRelating": {},
	"workHour":   {},
}

type commonColumn struct {
	Label      string
	ColumnType pb.ColumnType
	DataType   pb.StorageColumnType
}

var CommonColumns = map[string]*pb.Column{}

func init() {
	_ = encoding.GetJsonCodec().Unmarshal(DocumentBts, Document)

	defaultCommonColumns := map[string]*commonColumn{
		"creator":    {"创建人", pb.ColumnType_member, pb.StorageColumnType_STRING},
		"createTime": {"创建时间", pb.ColumnType_datepicker, pb.StorageColumnType_DATE},
		"updator":    {"更新人", pb.ColumnType_member, pb.StorageColumnType_STRING},
		"updateTime": {"更新时间", pb.ColumnType_datepicker, pb.StorageColumnType_DATE},
		"status":     {"启用状态", pb.ColumnType_status, pb.StorageColumnType_INT},
		"parentId":   {"数据ID", pb.ColumnType_input, pb.StorageColumnType_STRING},
		"relating":   {"关联", pb.ColumnType_relating, pb.StorageColumnType_CUSTOM},
		"baRelating": {"前后置", pb.ColumnType_baRelating, pb.StorageColumnType_CUSTOM},
	}
	for s, column := range defaultCommonColumns {
		CommonColumns[s] = &pb.Column{
			Name:  s,
			Label: column.Label,
			IsSys: true,
			Field: &pb.ColumnOption{
				Type:     column.ColumnType,
				DataType: column.DataType,
			},
		}
	}
}
