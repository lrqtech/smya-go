package boot

import (
	"fmt"
	"os"
	"smya/cache"
	"smya/corn"
	"smya/mqtt"
	"smya/mytime"
	"smya/request"
)

func init() {
	cache.InitCache()
	go corn.TimeTask()
	mytime.SetHour()
}

func Run() {
	fmt.Println("Starting ......")
	mqttRun()
}

func mqttRun() {
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
