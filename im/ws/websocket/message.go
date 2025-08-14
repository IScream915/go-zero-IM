package websocket

type Message struct {
	Method string      `json:"method,omitempty"`
	FormId string      `json:"formId,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
