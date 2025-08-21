package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf
	ListenOn string

	MsgChatTransfer kq.KqConf

	Rds RedisConfig

	Mongo MongoConfig

	Ws struct {
		Host string
	}
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type MongoConfig struct {
	Url string
	Db  string
}
