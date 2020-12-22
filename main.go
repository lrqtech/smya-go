package main

import (
	"fmt"
	"smya/config"
	"smya/core"
)

func main() {
	fmt.Println("Starting ......")
	config.GetMod()
	server, subscribe := core.Login()
	client := core.StartMqtt(server, subscribe)
	for {
		core.Subscribe(client, subscribe)
	}
}
