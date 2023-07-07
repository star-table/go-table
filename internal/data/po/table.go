package po

var (
	TableNameTable = "lc_app_table"

	TableSimpleSelect = "id,app_id,name,column_flag,auto_schedule_flag,summery_flag"
)

type Table struct {
	Base
	Name             string `gorm:"column:name; type:varchar(255);not null" json:"name"`                                                                                 // 表名
	OrgId            int64  `gorm:"index:idx_org_id_summery_flag;column:org_id;not null;default:0" json:"orgId,omitempty"`                                               // 组织id
	AppId            int64  `gorm:"index:idx_app_id;column:app_id;not null;default:0" json:"appId,string"`                                                               // 应用id
	Config           string `gorm:"column:config;type:json;not null;comment:配置" json:"config,omitempty"`                                                                 // 配置
	Creator          int64  `gorm:"column:creator;not null;default:0;comment:创建人" json:"creator,omitempty"`                                                              // 创建人
	Updater          int64  `gorm:"column:updater;not null;default:0;comment:更新人" json:"updater,omitempty"`                                                              // 更新人
	ColumnFlag       int32  `gorm:"column:column_flag;not null;default:1;comment:是否需要列数据,1是,2否" json:"columnFlag,omitempty"`                                             // 如果没有列数据，查询的时候会直接忽略，不浪费查询资源
	DelFlag          int32  `gorm:"column:del_flag;not null;default:2;comment:是否删除,1是,2否" json:"delFlag,omitempty"`                                                      // 是否删除,1是,2否
	SummeryFlag      int32  `gorm:"index:idx_org_id_summery_flag;column:summery_flag;not null;default:2;comment:1:全部任务汇总表,2:普通表, 3: 项目汇总表" json:"summeryFlag,omitempty"` //汇总表标识
	AutoScheduleFlag int32  `gorm:"column:auto_schedule_flag;not null;default:1" json:"autoScheduleFlag,omitempty"`                                                      // 自动排期是否开启,1否,2是
}

func (*Table) TableName() string {
	return TableNameTable
}
