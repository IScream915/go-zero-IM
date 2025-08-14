package user

import (
	"go-zero-IM/im/ws/internal/svc"
	_websocket "go-zero-IM/im/ws/websocket"

	"github.com/gorilla/websocket"
)

// OnLine 获取所有在线用户
func OnLine(svc *svc.ServiceContext) _websocket.HandlerFunc {
	return func(srv *_websocket.Server, conn *websocket.Conn, msg *_websocket.Message) {
		uids := srv.GetUsers()
		u := srv.GetUsers(conn)
		err := srv.Send(_websocket.NewMessage(u[0], uids), conn)
		srv.Info("err ", err)
	}
}
