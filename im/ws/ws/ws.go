package ws

import "go-zero-IM/im/dao/models"

type (
	Msg struct {
		models.MType `mapstructure:"mType"`
		Content      string `mapstructure:"content"`
	}

	Chat struct {
		ConversationId  string `mapstructure:"conversationId"`
		models.ChatType `mapstructure:"chatType"`
		SendId          string `mapstructure:"sendId"`
		RecvId          string `mapstructure:"recvId"`
		SendTime        int64  `mapstructure:"sendTime"`
		Msg             `mapstructure:"msg"`
	}
)
