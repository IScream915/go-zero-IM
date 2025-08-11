package logic

import (
	"context"
	"user/dao/model"
	"user/dao/query/implement"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserResp, error) {
	newUser := &model.User{
		Name:     in.Name,
		Phone:    in.Phone,
		Password: in.Password,
	}
	if txErr := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		dbTool := implement.NewDbToolHelper[model.User](tx)
		err := dbTool.InsertSingleRecord(l.ctx, newUser)
		if err != nil {
			return err
		}
		return nil
	}); txErr != nil {
		return nil, txErr
	}

	return &user.CreateUserResp{
		Id: newUser.ID,
	}, nil
}
