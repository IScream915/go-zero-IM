package user

import (
	"context"
	"errors"
	"go-zero-IM/pkg/ctxData"
	"go-zero-IM/user/rpc/user"

	"go-zero-IM/user/api/internal/svc"
	"go-zero-IM/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	uid, err := ctxData.GetUid(l.ctx)
	if err != nil {
		return nil, err
	}

	infoResp, err := l.svcCtx.User.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: uid,
	})
	if err != nil {
		return nil, errors.New("获取用户详情失败" + err.Error())
	}

	resp = &types.UserInfoResp{
		Info: types.User{
			Id:       infoResp.User.Id,
			Phone:    infoResp.User.Phone,
			Nickname: infoResp.User.Nickname,
			Status:   infoResp.User.Status,
			Sex:      infoResp.User.Sex,
		},
	}
	return
}
