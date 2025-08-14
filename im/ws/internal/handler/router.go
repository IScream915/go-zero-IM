package handler

import (
	"go-zero-IM/im/ws/internal/handler/user"
	"go-zero-IM/im/ws/internal/svc"
	"go-zero-IM/im/ws/websocket"
)

func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	routes := []websocket.Route{
		{
			Method:  "user.online",
			Handler: user.OnLine(svc),
		},
	}
	srv.AddRoutes(routes...)
}
