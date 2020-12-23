package core

import (
	"encoding/hex"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"smya/config"
	"smya/util"
	"strconv"
	"time"
)

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	str := ResDecrypt(msg.Payload())
	command, commandType, commandName := GetDetail(str)
	fmt.Printf("Run: %s \n", commandName)
	ExecShell(commandType, command, commandName)
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
	switch config.Mode {
	case "Debug", "debug", "DEBUG":
		mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
		mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
		fmt.Println("Probe: ")
		fmt.Printf("Get mqtt address: %s \n", server)
	}
	var broker = server
	var port = 5201
	opts := mqtt.NewClientOptions()
	switch config.Mode {
	case "Debug", "debug", "DEBUG":
		fmt.Println("Probe: ")
		fmt.Printf("Mqtt address: %s \n", fmt.Sprintf("tcp://%s:%d", broker, port))
	}
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	switch config.Mode {
	case "Debug", "debug", "DEBUG":
		fmt.Println("Probe: ")
		fmt.Printf("Mqtt config: \n SetClientID: %s \n SetUsername: %s \n SetPassword: %s \n SetKeepAlive: %d \n", config.DeviceId, config.DeviceId, subscribe, 60*time.Second)
	}
	opts.SetClientID(config.DeviceId)
	opts.SetUsername(config.DeviceId)
	opts.SetPassword(subscribe)
	opts.SetKeepAlive(60 * time.Second)
	switch config.Mode {
	case "Debug", "debug", "DEBUG":
		fmt.Println("Probe: ")
		fmt.Println("Starting mqtt connection ...... ")
	}
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Error() != nil {
		fmt.Println(token.Error())
	}
	return client
}

// 订阅
func Subscribe(client mqtt.Client, subscribe string) {
	switch config.Mode {
	case "Debug", "debug", "DEBUG":
		fmt.Println("Probe: ")
		fmt.Printf("Subscribe address: %s \n", fmt.Sprintf("client/%s", subscribe))
		fmt.Println("Qos: 1")
	}
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
