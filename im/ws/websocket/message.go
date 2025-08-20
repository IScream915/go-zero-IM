package websocket

type FrameType uint8

const (
	FrameData  FrameType = 0x0
	FramePing  FrameType = 0x1
	FrameAck   FrameType = 0x2
	FrameNoAck FrameType = 0x3
	FrameErr   FrameType = 0x9
)

type Message struct {
	FrameType `json:"frameType"`
	Method    string      `json:"method"`
	FromId    string      `json:"fromId"` // 消息发送方的Id
	Data      interface{} `json:"data"`
}

func NewMessage(fromId string, data interface{}) *Message {
	return &Message{
		FrameType: FrameData,
		FromId:    fromId,
		Data:      data,
	}
}

func NewErrMessage(err error) *Message {
	return &Message{
		FrameType: FrameErr,
		Data:      err.Error(),
	}
}
