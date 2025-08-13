package models

import (
	"time"

	"gorm.io/gorm"
)

type GroupMembers struct {
	ID          uint           `gorm:"column:id; type:int(11) unsigned; primarykey; autoIncrement; comment:主键ID" json:"id"`
	GroupID     string         `gorm:"column:group_id; type:varchar(64); not null; comment:群组ID" json:"group_id"`
	UserID      string         `gorm:"column:user_id; type:varchar(64); not null; comment:用户ID" json:"user_id"`
	RoleLevel   int            `gorm:"column:role_level; type:tinyint; not null; default:0; comment:角色等级 0-普通成员, 1-管理员, 2-群主" json:"role_level"`
	JoinTime    time.Time      `gorm:"column:join_time; type:timestamp; comment:加入时间" json:"join_time"`
	JoinSource  int            `gorm:"column:join_source; type:tinyint; default:0; comment:加入来源 0-搜索群聊, 1-推荐进群, 2-创建群聊进群" json:"join_source"`
	InviterUID  string         `gorm:"column:inviter_uid; type:varchar(64); comment:邀请者ID" json:"inviter_uid"`
	OperatorUID string         `gorm:"column:operator_uid; type:varchar(64); comment:操作者ID" json:"operator_uid"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at; type:timestamp; comment:删除时间" json:"deleted_at"`
}

func (GroupMembers) TableName() string {
	return "group_members"
}
