package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"smya/config"
	"time"
)

// 连接mqtt
func Client(server string, subscribe string) mqtt.Client {
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
	fmt.Println("Starting mqtt connection ...... ")
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Error() != nil {
		fmt.Println(token.Error())
	}
	return client
}
