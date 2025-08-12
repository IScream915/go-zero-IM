package svc

import (
	"fmt"
	"log"
	"user/rpc/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DSN 数据库连接地址
var DSN string

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initDatabase(c.Database)

	return &ServiceContext{
		Config: c,
		DB:     db,
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
