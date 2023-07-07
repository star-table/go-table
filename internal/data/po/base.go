package po

import "time"

type Base struct {
	ID       int64     `gorm:"primarykey;autoIncrement:false" json:"id,string"`
	CreateAt time.Time `gorm:"type:datetime;column:create_at;not null;default:CURRENT_TIMESTAMP" json:"createAt,omitempty"`                             // 创建日期
	UpdateAt time.Time `gorm:"type:datetime;column:update_at;not null;default:CURRENT_TIMESTAMP on update current_timestamp" json:"updateAt,omitempty"` // 更新日期
}
