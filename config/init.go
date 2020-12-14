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
		DeviceId = *CliDeviceId
		Key = *CliKey
		fmt.Println("Set Device Id: ", *CliDeviceId)
		fmt.Println("Set Safe Code: ", *CliKey)
	}
}
