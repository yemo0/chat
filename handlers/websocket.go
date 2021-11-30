package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// CheckOrigin防止跨站点请求伪造
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 在线列表
type AliveList struct {
	ConnList  map[string]*Client
	register  chan *Client
	destroy   chan *Client
	broadcast chan string // Message
	cancel    chan int
	Len       int
}

type Client struct {
	Id     string
	Conn   *websocket.Conn
	cancel chan int
}

type Message struct {
	ID         string                 // 发送消息id
	Content    string                 // 消息内容
	SentAt     int64                  `json:"sent_at"` // 发送时间
	Type       int                    // 消息类型, 如 BroadcastMessage
	From       string                 // 发送人client id
	To         []string               // 接收人client id, 根据消息类型来说, 单发, 群发, 广播什么的, 具体处理在Event中处理
	FromUserID string                 `json:"from_user_id"` // 发送者用户业务id
	ToUserID   string                 `json:"to_user_id"`   // 接受者用户业务id
	Ext        map[string]interface{} `json:"ext"`          // 扩展字段, 按需使用
}
