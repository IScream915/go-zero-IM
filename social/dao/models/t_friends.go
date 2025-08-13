package models

import (
	"time"

	"gorm.io/gorm"
)

type Friends struct {
	ID        uint           `gorm:"column:id; type:int(11) unsigned; primarykey; autoIncrement; comment:主键ID" json:"id"`
	UserID    string         `gorm:"column:user_id; type:varchar(64); not null; comment:用户ID" json:"user_id"`
	FriendUID string         `gorm:"column:friend_uid; type:varchar(64); not null; comment:好友ID" json:"friend_uid"`
	Remark    string         `gorm:"column:remark; type:varchar(255); default:''; comment:备注" json:"remark"`
	AddSource int            `gorm:"column:add_source; type:tinyint; default:0; not null; comment:添加来源 0-搜索添加, 1-群聊添加" json:"add_source"`
	CreatedAt time.Time      `gorm:"column:created_at; type:timestamp; comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at; type:timestamp; comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at; type:timestamp; comment:删除时间" json:"deleted_at"`
}

func (Friends) TableName() string {
	return "friends"
}
