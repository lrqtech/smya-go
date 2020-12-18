package config

import (
	"flag"
	"fmt"
	"os"
)

var (
	DeviceId string
	Key      string
)

var CliDeviceId = flag.String("id", "", "Input Your Device Id")
var CliKey = flag.String("passwd", "", "Input Your Safe Code")

func init() {
	flag.Parse()
	if *CliDeviceId == "" || *CliKey == "" {
		fmt.Println("Param error")
		fmt.Println("Please try again")
		fmt.Println("Use 'smya -h' for help")
		os.Exit(0)
	} else {
		if len(*CliDeviceId) != 32 {
			fmt.Printf("The length of device id %s isn't enough \n ", *CliDeviceId)
			fmt.Println("Please try again")
			os.Exit(0)
		} else {
			DeviceId = *CliDeviceId
		}
		if len(*CliKey) != 6 {
			fmt.Printf("The length of safe code %s isn't enough \n ", *CliKey)
			fmt.Println("Please try again")
			os.Exit(0)
		} else {
			Key = *CliKey
		}
	}
}
