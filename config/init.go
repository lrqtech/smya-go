package config

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	Cfg      *ini.File
	DeviceId string
	Key      string
)

func init() {
	var err error
	Cfg, err = ini.Load("./conf.ini")
	if err != nil {
		fmt.Printf("Fail to parse 'conf.ini': %v \n", err)
		return
	}
	LoadCfg()
}

func LoadCfg() {
	sec, err := Cfg.GetSection("account")
	if err != nil {
		fmt.Printf("Fail to get section 'account': %v \n", err)
		return
	}
	DeviceId = sec.Key("DeviceId").MustString("")
	Key = sec.Key("Key").MustString("")
}
