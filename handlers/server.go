package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

func WS(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !websocket.IsWebSocketUpgrade(r) {
		log.Printf("this is not a WebSocket")
		w.Write([]byte("暂无法连接..."))
		return
	}

	log.Printf("WebSocket connected")
	id := r.Header.Get("Sec-WebSocket-Key")
	cl, err := NewSocketClient(w, r, id)
	if err != nil {
		fmt.Println("创建新连接失败")
		return
	}

	for {
		_, message, err := cl.Conn.ReadMessage()
		if err != nil {
			log.Printf("ReadMessag Error: %v", err)
			break
		}
		cl.Broadcast(string(message))
	}
}
