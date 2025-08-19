package main

import (
	"flag"
	"fmt"
	"go-zero-IM/im/ws/internal/config"
	"go-zero-IM/im/ws/internal/handler"
	"go-zero-IM/im/ws/internal/svc"
	"go-zero-IM/im/ws/websocket"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "im/ws/etc/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c)
	srv := websocket.NewServer(c.ListenOn,
		websocket.WithServerAuthentication(handler.NewJwt(ctx)),
		websocket.WithServerMaxConnectionIdle(10*time.Second),
	)
	//websocket.WithServerAck(websocket.RigorAck),

	defer srv.Stop()

	handler.RegisterHandlers(srv, ctx)

	fmt.Println("start websocket server at ", c.ListenOn, " ..... ")
	srv.Start()

}
