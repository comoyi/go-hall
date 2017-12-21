package main

var users = make(map[*User]bool)

func AddUser(u *User) {
	users[u] = true
}

func RemoveUser(u *User) {
	delete(users, u)
}

func Users() map[*User]bool {
	return users
}
