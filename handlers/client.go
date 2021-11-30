package handlers

import (
	"fmt"
	"log"
	"net/http"
)

var aliveList *AliveList

func init() {
	aliveList = NewAliveList()
	go aliveList.run()
}

func NewSocketClient(w http.ResponseWriter, r *http.Request, id string) (cl *Client, err error) {
	ws, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade Error: %v", err)
		return nil, err
	}

	// defer ws.Close()

	cl = &Client{
		Id:     id,
		Conn:   ws,
		cancel: make(chan int, 1),
	}

	aliveList.ConnList[id] = ws
	aliveList.Register(cl)

	return cl, nil

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

func (cl *Client) Broadcast(message string) {
	aliveList.broadcast <- message
}

//
func (al *AliveList) run() {
	for {
		select {
		case client := <-al.register:
			fmt.Println(client.Id + "成功注册了!")
		case client := <-al.broadcast:
			fmt.Println(client)

		}

	}
}

func NewAliveList() *AliveList {
	return &AliveList{
		ConnList:  make(map[string]*Client, 100),
		register:  make(chan *Client, 100),
		destroy:   make(chan *Client, 100),
		broadcast: make(chan string, 100),
		cancel:    make(chan int),
		Len:       0,
	}
}

func (al *AliveList) Register(cl *Client) error {
	aliveList.register <- cl
	return nil
}

func (al *AliveList) Destroy(cl *Client) error {
	aliveList.destroy <- cl
	return nil
}

func (al *AliveList) Broadcast(message string) error {
	al.broadcast <- message
	return nil
}
