package models

import (
	"time"

	"gorm.io/gorm"
)

type GroupRequests struct {
	ID            uint           `gorm:"column:id; type:int(11) unsigned; primarykey; autoIncrement; comment:主键ID" json:"id"`
	ReqID         string         `gorm:"column:req_id; type:varchar(64); not null; comment:请求者ID" json:"req_id"`
	GroupID       string         `gorm:"column:group_id; type:varchar(64); not null; comment:群组ID" json:"group_id"`
	ReqMsg        string         `gorm:"column:req_msg; type:varchar(255); default:''; comment:请求消息" json:"req_msg"`
	ReqTime       time.Time      `gorm:"column:req_time; type:timestamp; comment:请求时间" json:"req_time"`
	JoinSource    int            `gorm:"column:join_source; type:tinyint; default:0; comment:加入来源 0-搜索群聊, 1-推荐进群, 2-创建群聊进群" json:"join_source"`
	InviterUserID string         `gorm:"column:inviter_user_id; type:varchar(64); comment:邀请者用户ID" json:"inviter_user_id"`
	HandleUserID  string         `gorm:"column:handle_user_id; type:varchar(64); comment:处理者用户ID" json:"handle_user_id"`
	HandleTime    *time.Time     `gorm:"column:handle_time; type:timestamp; comment:处理时间" json:"handle_time"`
	HandleResult  int            `gorm:"column:handle_result; type:tinyint; default:0; comment:处理结果 0-未处理, 1-拒绝, 2-成功" json:"handle_result"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at; type:timestamp; comment:删除时间" json:"deleted_at"`
}

func (GroupRequests) TableName() string {
	return "t_group_requests"
}
