package logic

import (
	"context"
	"errors"
	"go-zero-IM/pkg/ctxData"
	"go-zero-IM/pkg/encrypt"
	"go-zero-IM/user/dao/models"
	"go-zero-IM/user/dao/query/implement"
	"time"

	"go-zero-IM/user/rpc/internal/svc"
	"go-zero-IM/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneIsNotRegistered = errors.New("该手机号未进行注册")
	ErrPassword             = errors.New("密码错误")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// 根据手机号进行查询
	dbTool := implement.NewDbToolHelper[models.User](l.svcCtx.DB)
	record, err := dbTool.SearchSingleByField(l.ctx, "phone", in.Phone)
	if err != nil {
		return nil, errors.New("查询时出现错误")
	}

	// 查询不到手机号, 直接返回
	if record == nil {
		return nil, ErrPhoneIsNotRegistered
	}

	// 校验密码
	if !encrypt.ValidatePasswordHash(in.Password, record.Password) {
		// 校验失败, 直接返回
		return nil, ErrPassword
	}

	// jwt-token生成
	nowTime := time.Now().Unix()
	token, err := ctxData.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, nowTime, l.svcCtx.Config.Jwt.AccessExpire, record.ID)
	if err != nil {
		return nil, ErrGenToken
	}

	return &user.LoginResp{
		Token:  token,
		Expire: nowTime + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
