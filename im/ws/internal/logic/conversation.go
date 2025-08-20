package logic

import (
	"context"
	"go-zero-IM/im/dao/models"
	"go-zero-IM/im/ws/internal/svc"
	"go-zero-IM/im/ws/websocket"
	"go-zero-IM/im/ws/ws"
	"go-zero-IM/pkg/wuid"
	"time"
)

type Conversation struct {
	ctx context.Context
	srv *websocket.Server
	svc *svc.ServiceContext
}

func NewConversation(ctx context.Context, srv *websocket.Server, svc *svc.ServiceContext) *Conversation {
	return &Conversation{
		ctx: ctx,
		srv: srv,
		svc: svc,
	}
}

func (l *Conversation) SingleChat(data *ws.Chat, userId string) error {
	if data.ConversationId == "" {
		data.ConversationId = wuid.CombineId(userId, data.RecvId)
	}

	//time.Sleep(time.Minute)
	// 记录消息
	chatLog := models.ChatLog{
		ConversationId: data.ConversationId,
		SendId:         userId,
		RecvId:         data.RecvId,
		ChatType:       data.ChatType,
		MsgFrom:        0,
		MsgType:        data.MType,
		MsgContent:     data.Content,
		SendTime:       time.Now().UnixNano(),
	}

	// 插入聊天记录到MongoDB
	collection := l.svc.MongoDB.Collection("chat_logs")
	_, err := collection.InsertOne(l.ctx, chatLog)
	if err != nil {
		return err
	}

	return nil
}
