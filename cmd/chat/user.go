package main

type User struct {
	id     int
	name   string
	avatar string
	gender int
}

func NewUser(id int) *User {
	return &User{
		id: id,
	}
}
