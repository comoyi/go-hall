package main

type Room struct {
	id    int
	users []User
}

var roomId int = 0

func NewRoom() *Room {
	roomId++
	return &Room{
		id: roomId,
	}
}
