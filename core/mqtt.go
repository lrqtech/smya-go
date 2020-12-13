package core

import (
	"encoding/hex"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"smya/config"
	"smya/util"
	"strconv"
	"time"
)

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	str := ResDecrypt(msg.Payload())
	command, commandType, commandName := GetDetail(str)
	fmt.Printf("Run: %s \n", commandName)
	ExecShell(commandType, command)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connect success !")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v \n", err)
	fmt.Println("Try to restart program")
}

// 连接并启动mqtt
func StartMqtt(server string, subscribe string) mqtt.Client {
	var broker = server
	var port = 5201
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetClientID(config.DeviceId)
	opts.SetUsername(config.DeviceId)
	opts.SetPassword(subscribe)
	opts.SetKeepAlive(60 * time.Second)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Error() != nil {
		fmt.Println(token.Error())
	}
	return client
}

// 订阅
func Subscribe(client mqtt.Client, subscribe string) {
	topic := fmt.Sprintf("client/%s", subscribe)
	token := client.Subscribe(topic, 1, messageHandler)
	token.Wait()
}

// 资源解密
func ResDecrypt(s []byte) string {
	CurrentHour := strconv.Itoa(time.Now().Hour())
	if len(CurrentHour) == 1 {
		CurrentHour += "0"
	}
	iv := config.Key + CurrentHour
	ds, _ := hex.DecodeString(string(s))
	result, _ := util.DesDecryptCBC(ds, []byte(iv), []byte(iv), util.PKCS5PADDING)
	return string(result)
}
