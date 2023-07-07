package po

var TableNameOrgColumn = "lc_app_org_column"

// OrgColumn 组织(团队)字段
type OrgColumn struct {
	Base
	OrgId      int64  `gorm:"uniqueIndex:idx_org_id_column_id;column:org_id;not null;default:0" json:"org_id"`              // 表id
	ColumnId   string `gorm:"uniqueIndex:idx_org_id_column_id;column:column_id;type:varchar(32);not null" json:"column_id"` // 列id
	ColumnType string `gorm:"column:column_type;type:varchar(32);not null" json:"column_type"`
	Schema     string `gorm:"column:schema;type:json;not null;comment:列配置" json:"schema"` // 列配置
	Creator    int64  `gorm:"column:creator;not null;default:0" json:"creator,omitempty"`
	Updater    int64  `gorm:"column:updater;not null;default:0" json:"updater,omitempty"`
}

func (*OrgColumn) TableName() string {
	return TableNameOrgColumn
}
