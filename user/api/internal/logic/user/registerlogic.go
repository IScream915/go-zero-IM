package user

import (
	"context"
	"go-zero-IM/user/rpc/user"

	"go-zero-IM/user/api/internal/svc"
	"go-zero-IM/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 将入参传递到下游函数
	registerResp, err := l.svcCtx.User.Register(l.ctx, &user.RegisterReq{
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Password: req.Password,
		Sex:      int32(req.Sex),
	})
	if err != nil {
		return nil, err
	}

	resp = &types.RegisterResp{
		Token:  registerResp.Token,
		Expire: registerResp.Expire,
	}
	return
}
