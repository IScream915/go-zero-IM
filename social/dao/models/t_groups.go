package models

import (
	"time"

	"gorm.io/gorm"
)

type Groups struct {
	ID              string         `gorm:"column:id; type:varchar(24); primarykey; comment:主键ID" json:"id"`
	Name            string         `gorm:"column:name; type:varchar(255); not null; comment:群组名称" json:"name"`
	Icon            string         `gorm:"column:icon; type:varchar(255); not null; comment:群组图标" json:"icon"`
	Status          int            `gorm:"column:status; type:tinyint; default:0; not null; comment:状态 0-inactive,1-active" json:"status"`
	CreatorUID      string         `gorm:"column:creator_uid; type:varchar(64); not null; comment:创建者ID" json:"creator_uid"`
	GroupType       int            `gorm:"column:group_type; type:int(11); not null; comment:群组类型" json:"group_type"`
	IsVerify        bool           `gorm:"column:is_verify; type:boolean; not null; comment:是否需要验证" json:"is_verify"`
	Notification    string         `gorm:"column:notification; type:varchar(255); default:''; comment:群公告" json:"notification"`
	NotificationUID string         `gorm:"column:notification_uid; type:varchar(64); default:''; comment:公告发布者ID" json:"notification_uid"`
	CreatedAt       time.Time      `gorm:"column:created_at; type:timestamp; comment:创建时间" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at; type:timestamp; comment:更新时间" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at; type:timestamp; comment:删除时间" json:"deleted_at"`
}

func (Groups) TableName() string {
	return "groups"
}
