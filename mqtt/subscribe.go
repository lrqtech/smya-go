package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// 订阅
func Subscribe(client mqtt.Client, subscribe string) {
	topic := fmt.Sprintf("client/%s", subscribe)
	token := client.Subscribe(topic, 1, messageHandler)
	token.Wait()
}
