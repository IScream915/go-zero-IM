package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"primarykey; comment:主键ID" json:"id"`
	Nickname  string         `gorm:"column:nickname; type:string; size:30; uniqueIndex:uni_idx_nickname; not null; comment:用户名" json:"nickname"`
	Phone     string         `gorm:"column:phone; type:string; size:30; uniqueIndex:uni_idx_phone; not null; comment:手机号" json:"phone"`
	Password  string         `gorm:"column:password; type:string; size:30; not null; comment:密码" json:"password"`
	Status    int            `gorm:"column:status; type:tinyint; size:1; index:idx_status; default:0; not null; comment:状态 0-inactive,1-active" json:"status"`
	Sex       int            `gorm:"column:sex; type:tinyint; size:1; index:idx_sex; default:0; not null; comment:状态 0-女性,1-男性" json:"sex"`
	CreatedAt time.Time      `gorm:"column:created_at; type:datetime(0); comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at; type:datetime(0); comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at; type:datetime(0); comment:删除时间" json:"deleted_at"`
}

func (User) TableName() string {
	return "t_user"
}
