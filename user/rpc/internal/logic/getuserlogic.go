package logic

import (
	"context"
	"errors"
	"fmt"
	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	// 添加调试日志
	l.Infof("GetUser called with ID: %d", in.Id)
	l.Infof("Request object: %+v", in)

	if u, ok := users[fmt.Sprintf("%d", in.Id)]; ok {
		l.Infof("Found user: %+v", u)
		return &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}, nil
	}
	l.Infof("User with ID %d not found", in.Id)
	return nil, errors.New("查询用户不存在")
}
