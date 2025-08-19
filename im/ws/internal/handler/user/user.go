package user

import (
	"go-zero-IM/im/ws/internal/svc"
	"go-zero-IM/im/ws/websocket"
)

// OnLine 获取所有在线用户
func OnLine(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		uids := srv.GetUsers()
		u := srv.GetUsers(conn)
		err := srv.Send(websocket.NewMessage(u[0], uids), conn)
		srv.Info("err ", err)
	}
}
