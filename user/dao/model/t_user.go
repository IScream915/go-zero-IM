package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey; comment:主键ID" json:"id"`
	Username  string         `gorm:"column:username; type:string; size:30; uniqueIndex:uni_idx_username; not null; comment:用户名" json:"username"`
	Phone     string         `gorm:"column:phone; type:string; size:30; uniqueIndex:uni_idx_phone; not null; comment:手机号" json:"phone"`
	Password  string         `gorm:"column:password; type:string; size:30; not null; comment:密码" json:"password"`
	Status    int            `gorm:"column:status; type:tinyint; size:1; index:idx_status; default:0; not null; comment:状态 0-inactive,1-active" json:"status"`
	CreatedAt time.Time      `gorm:"column:created_at; comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at; comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at; comment:删除时间" json:"deleted_at"`
}

func (User) TableName() string {
	return "t_user"
}
