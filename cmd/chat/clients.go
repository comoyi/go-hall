package main

var clients = make(map[*Client]bool)

func AddClient(c *Client) {
	clients[c] = true
}

func RemoveClient(c *Client) {
	delete(clients, c)
}

func Clients() map[*Client]bool {
	return clients
}
