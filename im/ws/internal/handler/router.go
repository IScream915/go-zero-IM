package handler

import (
	"go-zero-IM/im/ws/internal/handler/conversation"
	"go-zero-IM/im/ws/internal/handler/push"
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
		{
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
		{
			Method:  "push",
			Handler: push.Push(svc),
		},
	}
	srv.AddRoutes(routes...)
}
