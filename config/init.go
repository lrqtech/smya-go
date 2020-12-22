package config

import (
	"flag"
	"fmt"
	"os"
)

var (
	DeviceId string
	Key      string
	Mode     string
)

var CliDeviceId = flag.String("id", "", "Input Your Device Id")
var CliKey = flag.String("passwd", "", "Input Your Safe Code")
var CliMode = flag.String("mode", "Normal", "Ruing mode selection [ Normal / Debug ]\n Default: Normal")

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
		Mode = *CliMode
	}
}

// 获取运行模式
func GetMod() {
	var pmode string
	switch Mode {
	case "Normal":
		pmode = "Normal"
	case "Debug", "debug", "DEBUG":
		pmode = "Debug"
	}
	fmt.Printf("Ruing mode: %s \n", pmode)
}
