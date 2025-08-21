package svc

import (
	"context"
	"fmt"
	"go-zero-IM/pkg/ctxData"
	"go-zero-IM/user/rpc/internal/config"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DSN 数据库连接地址
var DSN string

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Rds    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initDatabase(c.Database)
	rds := initRedis(c.Rds)

	return &ServiceContext{
		Config: c,
		DB:     db,
		Rds:    rds,
	}
}

func initDatabase(cfg config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
		cfg.Charset,
	)

	DSN = dsn

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

func initRedis(cfg config.RedisConfig) *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	fmt.Println("Redis", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), "connect success")
	return rds
}

func (svc *ServiceContext) SetRootToken() error {
	// 生成jwt
	systemToken, err := ctxData.GetJwtToken(
		svc.Config.Jwt.AccessSecret,
		time.Now().Unix(),
		999999999,
		ctxData.SYSTEM_ROOT_UID,
	)
	if err != nil {
		return err
	}

	// 写入到redis
	return svc.Rds.Set(
		context.Background(),
		ctxData.REDIS_SYSTEM_ROOT_TOKEN,
		systemToken,
		time.Duration(999999999)*time.Second,
	).Err()
}
