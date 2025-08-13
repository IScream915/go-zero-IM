package main

import (
	"flag"
	"fmt"
	"go-zero-IM/social/dao/models"
	"log"

	"go-zero-IM/social/rpc/internal/config"
	"go-zero-IM/social/rpc/internal/server"
	"go-zero-IM/social/rpc/internal/svc"
	"go-zero-IM/social/rpc/social"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/social.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// 执行数据库迁移
	if err := ctx.DB.AutoMigrate(
		&models.Friends{},
		&models.FriendRequests{},
		&models.Groups{},
		&models.GroupMembers{},
		&models.GroupRequests{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	fmt.Println("Database migration completed successfully")

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		social.RegisterSocialServer(grpcServer, server.NewSocialServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
