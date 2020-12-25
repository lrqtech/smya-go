package core

import (
	"encoding/hex"
	"encoding/json"
	"smya/config"
	"smya/model"
	"smya/util"
	"strconv"
	"time"
)

// 登入
func Login() (string, string) {
	str := ResEncrypt()
	url := "https://smya.cn/client/login"
	data := model.Info{
		DeviceId: config.DeviceId,
		Str:      str,
	}
	body, _ := json.Marshal(data)
	req := util.Request("POST", url, body, 1024)
	var r model.Response
	_ = json.Unmarshal(req, &r)
	return r.Data["server"], r.Data["subscribe"]
}

// 数据加密
func ResEncrypt() string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	CurrentHour := strconv.Itoa(time.Now().In(loc).Hour())
	if len(CurrentHour) == 1 {
		CurrentHour += "0"
	}
	k := config.Key + CurrentHour
	s := config.Key + config.DeviceId
	result, _ := util.DesEncryptCBC([]byte(s), []byte(k), []byte(k), util.PKCS5PADDING)
	return hex.EncodeToString(result)
}
