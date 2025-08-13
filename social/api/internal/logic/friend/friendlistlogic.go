package friend

import (
	"context"
	"go-zero-IM/pkg/ctxData"
	"go-zero-IM/social/rpc/social"

	"go-zero-IM/social/api/internal/svc"
	"go-zero-IM/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友列表
func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	// 从ctx中获取id
	id, err := ctxData.GetUid(l.ctx)
	if err != nil {
		return nil, err
	}

	// 从下游函数获取结果
	result, err := l.svcCtx.Social.FriendList(l.ctx, &social.FriendListReq{
		UserId: id,
	})
	if err != nil {
		return nil, err
	}

	// 对结果进行转换
	records := make([]*types.Friends, 0)
	for _, record := range result.List {
		records = append(records, &types.Friends{
			Id:        record.Id,
			Nickname:  record.Nickname,
			Remark:    record.Remark,
			AddSource: record.AddSource,
		})
	}

	resp = &types.FriendListResp{
		List: records,
	}

	return
}
