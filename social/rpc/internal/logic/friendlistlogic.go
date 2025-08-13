package logic

import (
	"context"
	"errors"
	"fmt"
	"go-zero-IM/social/dao/models"
	"go-zero-IM/social/rpc/internal/svc"
	"go-zero-IM/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *social.FriendListReq) (*social.FriendListResp, error) {
	friendTable := models.Friends{}.TableName()

	friends := make([]*models.Friends, 0)
	// 获取好友列表
	if err := l.svcCtx.DB.
		WithContext(l.ctx).
		Table(fmt.Sprintf("%s AS f", friendTable)).
		Where(fmt.Sprintf("f.`user_id` = '%s'", in.UserId)).
		Find(&friends).Error; err != nil {
		return nil, errors.New("查询" + friendTable + "时出现问题")
	}

	// 将好友信息进行处理
	result := make([]*social.Friends, 0)
	for _, friend := range friends {
		result = append(result, &social.Friends{
			Id:        friend.FriendUID,
			Remark:    friend.Remark,
			AddSource: int32(friend.AddSource),
		})
	}

	return &social.FriendListResp{
		List: result,
	}, nil
}
