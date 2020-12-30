package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"smya/exec"
	"smya/request"
	"smya/resource"
)

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	str := resource.ResDecrypt(msg.Payload())
	command, commandType, commandName := request.GetDetail(str)
	exec.Shell(commandType, command, commandName)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connect success !")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v \n", err)
	fmt.Println("Try to restart program")
}
