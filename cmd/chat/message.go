package main

const (
	MESSAGE_TYPE_TEXT  = 1
	MESSAGE_TYPE_IMAGE = 2
	MESSAGE_TYPE_FILE  = 3
	MESSAGE_TYPE_AUDIO = 4
	MESSAGE_TYPE_VIDEO = 5
)

type Message struct {
	id          int
	messageType int
	senderId    int
	receiverId  int
	data        interface{}
}

var messageId int = 0

func NewMessage(messageType int, senderId int, receiverId int, data interface{}) *Message {
	messageId++
	return &Message{
		id:          messageId,
		messageType: messageType,
		senderId:    senderId,
		receiverId:  receiverId,
		data:        data,
	}
}
