package main

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	id    int
	conn  *websocket.Conn
	queue chan []byte
}

var clientId = 0

func NewClient(conn *websocket.Conn) *Client {

	clientId++

	return &Client{
		id:    clientId,
		conn:  conn,
		queue: make(chan []byte, 256),
	}
}

func (c *Client) Read() {
	defer func() {
		c.conn.Close()
	}()
	for {
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Info("read error: ", err, messageType, message)
			delete(clients, c)
			break
		}
		log.Info("[recv data]", "[client id:", c.id, "] ", "[message type:", messageType, "] ", "[message:", string(message), "]")
		for i := range clients {
			i.queue <- message
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.conn.Close()
	}()
	for {
		message, _ := <-c.queue
		log.Info("[send data]", "[client id:", c.id, "] ", string(message))
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Warning(err, c.id)
			break
		}
	}
}
