package svc

import (
	"user/api/internal/config"
	"user/api/internal/middleware"
	"user/rpc/userclient"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	User   userclient.User

	// 中间件
	LoginVerification rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		User:              userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle,
	}
}
