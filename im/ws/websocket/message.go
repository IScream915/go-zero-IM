package websocket

type Message struct {
	Method string      `json:"method"`
	FromId string      `json:"fromId"` // 消息发送方的Id
	Data   interface{} `json:"data"`
}

func NewMessage(fromId string, data interface{}) *Message {
	return &Message{
		FromId: fromId,
		Data:   data,
	}
}
