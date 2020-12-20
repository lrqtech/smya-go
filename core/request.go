package core

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"smya/config"
	"smya/util"
	"strconv"
	"time"
)

type Info struct {
	DeviceId string `json:"device_id"`
	Str      string `json:"str"`
}

type Response struct {
	Code int
	Msg  string
	Data map[string]string
}

// 登入
func Login() (string, string) {
	str := ResEncrypt()
	url := "https://smya.cn/client/login"
	data := Info{
		DeviceId: config.DeviceId,
		Str:      str,
	}
	body, err := json.Marshal(data)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	req := util.Request("POST", url, body, 1024)
	var r Response
	_ = json.Unmarshal(req, &r)
	return r.Data["server"], r.Data["subscribe"]
}

// 数据加密
func ResEncrypt() string {
	CurrentHour := strconv.Itoa(time.Now().Hour())
	if len(CurrentHour) == 1 {
		CurrentHour += "0"
	}
	k := config.Key + CurrentHour
	s := config.Key + config.DeviceId
	result, _ := util.DesEncryptCBC([]byte(s), []byte(k), []byte(k), util.PKCS5PADDING)
	return hex.EncodeToString(result)
}
