package main

import (
	"fmt"
	"os"
	"smya/cache"
	"smya/corn"
	"smya/mqtt"
	"smya/mytime"
	"smya/request"
)

func main() {
	fmt.Println("Starting ......")
	cache.InitCache()
	go corn.TimeTask()
	mytime.SetHour()
	server, subscribe := request.Login()
	client := mqtt.Client(server, subscribe)
	if subscribe == "" {
		fmt.Println("Subscribe address error")
		fmt.Println("Please restart program again")
		os.Exit(0)
	}
	for {
		mqtt.Subscribe(client, subscribe)
	}
}
