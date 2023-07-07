package po

import "time"

const TableNameLcAppFormBase = "lc_app_form_base"

// LcAppFormBase mapped from table <lc_app_form_base>
type LcAppFormBase struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`                                           // 主键
	OrgID      int64     `gorm:"column:org_id;not null;default:0" json:"org_id"`                           // 组织id
	Config     string    `gorm:"column:config;not null" json:"config"`                                     // 字段配置
	Status     int32     `gorm:"column:status;not null;default:1" json:"status"`                           // 1:可用，2：不可用
	Creator    int64     `gorm:"column:creator;not null;default:0" json:"creator"`                         // 创建人
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	Updator    int64     `gorm:"column:updator;not null;default:0" json:"updator"`                         // 更新人
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	Version    int32     `gorm:"column:version;not null;default:1" json:"version"`                         // 乐观锁
	DelFlag    int32     `gorm:"column:del_flag;not null;default:2" json:"del_flag"`                       // 是否删除,1是,2否
}

// TableName LcAppFormBase's table name
func (*LcAppFormBase) TableName() string {
	return TableNameLcAppFormBase
}
