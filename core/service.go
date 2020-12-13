package core

import "fmt"

// 重启连接
func ReStartServer() {
	server, subscribe := Login()
	client := StartMqtt(server, subscribe)
	fmt.Println("Reconnect success !")
	for {
		Subscribe(client, subscribe)
	}
}
