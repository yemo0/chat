package handlers

import (
	"fmt"
	"log"
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

type client struct {
	Id   string
	Conn *websocket.Conn
}

// 在线列表
var connList []*client

func NewSocketClient(w http.ResponseWriter, r *http.Request, id string) (cl *client, err error) {
	ws, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade Error: %v", err)
		return
	}

	// defer ws.Close()

	cl = &client{
		Id:   id,
		Conn: ws,
	}

	connList = append(connList, cl)

	return cl, nil
	// wslist = append(wslist, ws)

	// for {
	// 	mt, message, err := ws.ReadMessage()
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Printf("massageType: %v "+string(message)+"\n", mt)

	// 	err = ws.WriteMessage(1, message)
	// 	if err != nil {
	// 		fmt.Printf("WriteMessage Error: %v", err)
	// 		break
	// 	}

	// 	fmt.Println("100")
	// 	err = cl.Conn.WriteMessage(1, []byte("早上好"))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	time.Sleep(time.Second)
	// }
	// return cl, nil

}

func (cl *client) Run() {
	for {
		_, message, err := cl.Conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("massageType: %v "+string(message)+"\n", cl.Id)

		for _, v := range connList {
			fmt.Println(cl.Id + "              " + v.Id)
			if v.Id != cl.Id {
				err = v.Conn.WriteMessage(1, message)
				if err != nil {
					fmt.Printf("WriteMessage Error: %v", err)
					break
				}
			}
		}

		// cl.Conn.WriteMessage(1, []byte("ok"))
		// time.Sleep(time.Second * 1)
	}
}
