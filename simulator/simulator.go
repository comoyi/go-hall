package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go simulate(i)
	}
	select {}
}

func simulate(id int) {
	var dial *websocket.Dialer
	conn, response, err := dial.Dial("ws://127.0.0.1:8080/ws", nil)
	if err != nil {
		fmt.Println(conn, response, err)
		return
	}
	go receive(conn, id)
	go send(conn, id)
}

func send(conn *websocket.Conn, id int) {
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("client["+strconv.Itoa(id)+"]"+time.Now().Format("2006-01-02 15:04:05")))
		time.Sleep(1000 * time.Millisecond)
	}
}

func receive(conn *websocket.Conn, id int) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error: ", err)
			return
		}

		fmt.Printf("client[%d]received: %s\n", id, message)
	}
}
