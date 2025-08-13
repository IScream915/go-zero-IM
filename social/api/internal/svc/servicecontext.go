package svc

import (
	"go-zero-IM/social/api/internal/config"
	"go-zero-IM/social/rpc/socialclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Social socialclient.Social
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
	}
}
