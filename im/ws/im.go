package main

import (
	"flag"
	"fmt"
	"go-zero-IM/im/ws/internal/config"
	"go-zero-IM/im/ws/internal/svc"
	"go-zero-IM/im/ws/websocket"

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

	svc.NewServiceContext(c)
	srv := websocket.NewServer(c.ListenOn) //websocket.WithServerAuthentication(handler.NewJwtAuth(ctx)),
	//websocket.WithServerAck(websocket.RigorAck),
	//websocket.WithServerMaxConnectionIdle(10*time.Second),

	defer srv.Stop()

	//handler.RegisterHandlers(srv, ctx)

	fmt.Println("start websocket server at ", c.ListenOn, " ..... ")
	srv.Start()

}
