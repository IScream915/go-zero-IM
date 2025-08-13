package logic

import (
	"context"
	"errors"
	"go-zero-IM/pkg/ctxData"
	"go-zero-IM/pkg/encrypt"
	"go-zero-IM/pkg/wuid"
	"go-zero-IM/user/dao/models"
	"go-zero-IM/user/dao/query/implement"
	"time"

	"go-zero-IM/user/rpc/internal/svc"
	"go-zero-IM/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneIsRegistered = errors.New("该手机号已经注册过")
	ErrEncrypt           = errors.New("密码加密错误")
	ErrInsertUser        = errors.New("创建新用户失败")
	ErrGenToken          = errors.New("生成token失败")
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// 查询手机号是否已经注册
	dbTool := implement.NewDbToolHelper[models.User](l.svcCtx.DB)
	record, err := dbTool.SearchSingleByField(l.ctx, "phone", in.Phone)
	if err != nil {
		return nil, errors.New("查询时出现错误")
	}

	// 如果查询的到手机号则说明已经注册过, 返回报错
	if record != nil {
		return nil, ErrPhoneIsRegistered
	}

	// 注册信息登记
	// 对密码加密
	genPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
	if err != nil {
		return nil, ErrEncrypt
	}

	newUser, err := dbTool.InsertSingleRecordAndReturn(l.ctx, &models.User{
		ID:       wuid.GenUid(svc.DSN),
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Password: string(genPassword),
		Status:   0,
		Sex:      int(in.Sex),
	})
	if err != nil {
		return nil, ErrInsertUser
	}

	// jwt-token生成
	nowTime := time.Now().Unix()
	token, err := ctxData.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, nowTime, l.svcCtx.Config.Jwt.AccessExpire, newUser.ID)
	if err != nil {
		return nil, ErrGenToken
	}

	return &user.RegisterResp{
		Token:  token,
		Expire: nowTime + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
