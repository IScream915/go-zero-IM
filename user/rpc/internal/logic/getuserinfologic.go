package logic

import (
	"context"
	"errors"
	"user/dao/models"
	"user/dao/query/implement"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrUserNotFound = errors.New("无法查询到该用户")
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// 根据id进行查找
	dbTool := implement.NewDbToolHelper[models.User](l.svcCtx.DB)
	record, err := dbTool.SearchSingleByField(l.ctx, "id", in.Id)
	if err != nil {
		return nil, errors.New("查询时出现错误")
	}

	if record == nil {
		// 查找不到该用户
		return nil, ErrUserNotFound
	}

	return &user.GetUserInfoResp{
		User: &user.UserEntity{
			Id:       record.ID,
			Nickname: record.Nickname,
			Phone:    record.Phone,
			Status:   int32(record.Status),
			Sex:      int32(record.Sex),
		},
	}, nil
}
