package main

import (
	"fmt"
	"os"
	"smya/config"
	"smya/core"
)

func main() {
	fmt.Println("Starting ......")
	config.GetMod()
	server, subscribe := core.Login()
	client := core.StartMqtt(server, subscribe)
	for {
		if subscribe == "" {
			fmt.Println("Subscribe Error")
			os.Exit(0)
		}
		core.Subscribe(client, subscribe)
	}
}
