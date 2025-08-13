package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	ResultUnhandled = iota
	ResultRefuse
	ResultApprove
)

type FriendRequests struct {
	ID           uint           `gorm:"column:id; type:int(11) unsigned; primarykey; autoIncrement; comment:主键ID" json:"id"`
	UserID       string         `gorm:"column:user_id; type:varchar(64); not null; comment:用户ID" json:"user_id"`
	ReqUID       string         `gorm:"column:req_uid; type:varchar(64); not null; comment:请求用户ID" json:"req_uid"`
	ReqMsg       string         `gorm:"column:req_msg; type:varchar(255); default:'' ;comment:请求消息" json:"req_msg"`
	ReqTime      time.Time      `gorm:"column:req_time; type:timestamp; not null; comment:请求时间" json:"req_time"`
	HandleResult int            `gorm:"column:handle_result; type:tinyint; default:0; comment:处理结果 0-未处理, 1-拒绝, 2-通过" json:"handle_result"`
	HandleMsg    string         `gorm:"column:handle_msg; type:varchar(255); default:''; comment:处理消息" json:"handle_msg"`
	HandledAt    *time.Time     `gorm:"column:handled_at; type:timestamp; comment:处理时间" json:"handled_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at; type:timestamp; comment:删除时间" json:"deleted_at"`
}

func (FriendRequests) TableName() string {
	return "t_friend_requests"
}
