package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Database DatabaseConfig
	Jwt      JwtConfig
}

type DatabaseConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	Charset      string
}

type JwtConfig struct {
	AccessSecret string
	AccessExpire int64
}
