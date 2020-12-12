package main

import (
	"smya/core"
)

func main() {
	server, subscribe := core.Login()
	client := core.StartMqtt(server, subscribe)
	for {
		core.Subscribe(client, subscribe)
	}
}
