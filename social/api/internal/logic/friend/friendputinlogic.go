package friend

import (
	"context"
	"go-zero-IM/pkg/ctxData"
	"go-zero-IM/social/rpc/social"

	"go-zero-IM/social/api/internal/svc"
	"go-zero-IM/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友申请
func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInLogic) FriendPutIn(req *types.FriendPutInReq) (resp *types.FriendPutInResp, err error) {
	// 从ctx中获取id
	id, err := ctxData.GetUid(l.ctx)
	if err != nil {
		return nil, err
	}

	if _, err = l.svcCtx.Social.FriendPutIn(l.ctx, &social.FriendPutInReq{
		UserId: id,
		ReqUid: req.ReqUid,
		ReqMsg: req.ReqMsg,
	}); err != nil {
		return nil, err
	}

	return
}
