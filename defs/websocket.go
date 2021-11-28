package defs

// type Client struct {
// 	Id   string
// 	Conn *websocket.Conn
// }

type Message struct {
	Id      string
	Content string
	SentAt  int64
	Type    int // 消息类型
}
