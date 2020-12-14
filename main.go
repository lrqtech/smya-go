package main

import (
	"fmt"
	"smya/core"
)

func main() {
	fmt.Println("Starting ......")
	server, subscribe := core.Login()
	client := core.StartMqtt(server, subscribe)
	for {
		core.Subscribe(client, subscribe)
	}
}
