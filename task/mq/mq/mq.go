package mq

import "go-zero-IM/im/dao/models"

type MsgChatTransfer struct {
	ConversationId  string `json:"conversationId"`
	models.ChatType `json:"chatType"`
	SendId          string `json:"sendId"`
	RecvId          string `json:"recvId"`
	SendTime        int64  `json:"sendTime"`

	models.MType `json:"mType"`
	Content      string `json:"content"`
}
