package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MType int

const (
	TextMType MType = iota
)

type ChatType int

const (
	GroupChatType ChatType = iota + 1
	SingleChatType
)

type ChatLog struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	ConversationId string   `bson:"conversationId"`
	SendId         string   `bson:"sendId"`
	RecvId         string   `bson:"recvId"`
	MsgFrom        int      `bson:"msgFrom"`
	ChatType       ChatType `bson:"chatType"`
	MsgType        MType    `bson:"msgType"`
	MsgContent     string   `bson:"msgContent"`
	SendTime       int64    `bson:"sendTime"`
	Status         int      `bson:"status"`

	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
