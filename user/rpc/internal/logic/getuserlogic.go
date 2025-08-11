package logic

import (
	"context"
	"errors"
	"user/dao/model"
	"user/dao/query/implement"
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
	dbTool := implement.NewDbToolHelper[model.User](l.svcCtx.DB)
	record, err := dbTool.SearchSingleByField(l.ctx, "id", in.Id)
	if err != nil {
		return nil, err
	}

	if record == nil {
		return nil, errors.New("查询用户不存在")
	} else {
		return &user.GetUserResp{
			Id:    record.ID,
			Name:  record.Name,
			Phone: record.Phone,
		}, nil
	}
}
