package friend

import (
	"context"
	"go-zero-IM/pkg/ctxData"
	"go-zero-IM/social/rpc/social"

	"go-zero-IM/social/api/internal/svc"
	"go-zero-IM/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友申请处理
func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInHandleLogic) FriendPutInHandle(req *types.FriendPutInHandleReq) (resp *types.FriendPutInHandleResp, err error) {
	// 从ctx中获取id
	id, err := ctxData.GetUid(l.ctx)
	if err != nil {
		return nil, err
	}

	if _, err = l.svcCtx.Social.FriendPutInHandle(l.ctx, &social.FriendPutInHandleReq{
		FriendReqId:  req.FriendReqId,
		UserId:       id,
		HandleResult: req.HandleResult,
	}); err != nil {
		return nil, err
	}

	return
}
