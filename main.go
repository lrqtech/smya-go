package main

import (
	"fmt"
	"os"
	"smya/core"
)

func main() {
	fmt.Println("Starting ......")
	server, subscribe := core.Login()
	client := core.StartMqtt(server, subscribe)
	if subscribe == "" {
		fmt.Println("Subscribe Error")
		os.Exit(0)
	}
	for {
		core.Subscribe(client, subscribe)
	}
}
