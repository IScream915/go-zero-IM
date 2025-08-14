package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type Server struct {
	addr     string
	upgrader *websocket.Upgrader
	logx.Logger

	// 路由以及对应的处理方法 map[路由]路由的处理函数
	routes map[string]HandlerFunc

	// 并发安全的读写锁
	sync.RWMutex

	// 连接存储
	connToUser map[*websocket.Conn]string
	userToConn map[string]*websocket.Conn

	// 鉴权
	authentication Authentication
}

func NewServer(addr string) *Server {
	return &Server{
		addr:     addr,
		upgrader: &websocket.Upgrader{},
		Logger:   logx.WithContext(context.Background()),

		routes: make(map[string]HandlerFunc),

		connToUser: make(map[*websocket.Conn]string),
		userToConn: make(map[string]*websocket.Conn),

		authentication: new(authentication),
	}
}

// ServerWs WebSocket主服务
func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			s.Errorf("server handler ws recover err %v", r)
		}
	}()

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Errorf("upgrader ws err %v", err)
		return
	}

	// 根据连接对象执行任务处理
	go s.handelConn(conn)
}

// AddRoutes 添加路由
func (s *Server) AddRoutes(rs ...Route) {
	for _, route := range rs {
		s.routes[route.Method] = route.Handler
	}
}

// handelConn 根据连接对象执行任务处理
func (s *Server) handelConn(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.Errorf("websocket conn message read err %v", err)
			// todo: 关闭连接
			return
		}

		var message Message
		if err = json.Unmarshal(msg, &message); err != nil {
			s.Errorf("json message unmarshal err %v, msg %v", err, string(msg))
			// todo: 关闭连接
			return
		}

		// 根据请求的method分发路由并执行
		if handler, ok := s.routes[message.Method]; ok {
			handler(s, conn, &message)
		} else {
			conn.WriteMessage(websocket.TextMessage,
				[]byte(fmt.Sprintf("路由:% 没有可以执行的方法, 请检查路由以及对应方法", message.Method)))
		}
	}
}

func (s *Server) addConn(conn *websocket.Conn, req *http.Request) {
	uid := s.authentication.UserId(req)

	// 上锁保证对map访问的协程安全
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	// 验证用户是否之前登入过
	if c := s.userToConn[uid]; c != nil {
		// 关闭之前的连接
		c.Close()
	}

	s.connToUser[conn] = uid
	s.userToConn[uid] = conn
}

func (s *Server) Close(conn *websocket.Conn) {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	uid := s.connToUser[conn]
	if uid == "" {
		// 已经被关闭
		return
	}

	delete(s.connToUser, conn)
	delete(s.userToConn, uid)

	conn.Close()
}

func (s *Server) GetConn(uid string) *websocket.Conn {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	if conn, ok := s.userToConn[uid]; ok {
		return conn
	}

	return nil
}

func (s *Server) GetConns(uids ...string) []*websocket.Conn {
	if len(uids) == 0 {
		return nil
	}

	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	res := make([]*websocket.Conn, 0, len(uids))
	for _, uid := range uids {
		res = append(res, s.userToConn[uid])
	}
	return res
}

func (s *Server) GetUsers(conns ...*websocket.Conn) []string {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	var res []string
	if len(conns) == 0 {
		// 获取全部
		res = make([]string, 0, len(s.connToUser))
		for _, uid := range s.connToUser {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, s.connToUser[conn])
		}
	}

	return res
}

func (s *Server) Start() {
	http.HandleFunc("/ws", s.ServerWs)
	fmt.Println("websocket已启动, 正在监听...")
	s.Info(http.ListenAndServe(s.addr, nil))
}

func (s *Server) Stop() {
	fmt.Println("websocket服务已停止")
}
