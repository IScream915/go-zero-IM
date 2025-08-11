package logic

import (
	"context"
	"errors"
	"user/rpc/user"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserReq) (resp *types.UserResp, err error) {
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
