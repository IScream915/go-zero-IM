package config

import "github.com/zeromicro/go-zero/core/service"

type Config struct {
	service.ServiceConf

	ListenOn string

	Jwt struct {
		AccessSecret string
	}

	Mongo MongoConfig

	MsgChatTransfer struct {
		Topic string
		Addrs []string
	}
}

type MongoConfig struct {
	Url string
	Db  string
}
