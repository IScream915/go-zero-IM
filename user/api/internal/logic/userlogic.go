package logic

import (
	"context"
	"errors"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	getUserResp, err := l.svcCtx.User.GetUser(l.ctx, &user.GetUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}

	resp = &types.UserResp{
		Id:    getUserResp.Id,
		Name:  getUserResp.Name,
		Phone: getUserResp.Phone,
	}

	return resp, nil
}
