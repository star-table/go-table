package po

var TableNameTableColumn = "lc_app_table_column"

type TableColumn struct {
	Base
	OrgId             int64  `gorm:"index:idx_org_id_del_flag_column_id;column:org_id;not null;default:0" json:"org_id,omitempty"`
	TableId           int64  `gorm:"uniqueIndex:idx_table_id_column_id;column:table_id;not null;default:0" json:"table_id"`                                              // 表id
	ColumnId          string `gorm:"uniqueIndex:idx_table_id_column_id;index:idx_org_id_del_flag_column_id;column:column_id;type:varchar(64);not null" json:"column_id"` // 列id
	ColumnType        string `gorm:"column:column_type;type:varchar(32);not null" json:"column_type"`
	Schema            string `gorm:"column:schema;type:json;not null;comment:列配置" json:"schema"`         // 列配置
	Description       string `gorm:"column:description;type:varchar(8000);not null;default:''" json:"-"` // 列描述
	SourceOrgColumnId string `gorm:"column:source_org_column_id;type:varchar(32);not null;default:'';comment:来源的组织字段名，用于关联组织字段" json:"source_org_column_id,omitempty"`
	DelFlag           int32  `gorm:"index:idx_org_id_del_flag_column_id;column:del_flag;not null;default:2;comment:是否删除,1是,2否" json:"del_flag,omitempty"`
	Creator           int64  `gorm:"column:creator;not null;default:0" json:"creator,omitempty"`
	Updater           int64  `gorm:"column:updater;not null;default:0" json:"updater,omitempty"`
}

func (*TableColumn) TableName() string {
	return TableNameTableColumn
}
