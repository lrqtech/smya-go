package core

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"smya/config"
	"smya/model"
	"time"
)

// 生成发布地址
func GenPubAddr() string {
	sum256 := sha256.Sum256([]byte(config.DeviceId + config.Key))
	return base64.RawStdEncoding.EncodeToString(sum256[:])
}

// 格式化Msg
func MsgFmtJson(data string) string {
	j := model.Msg{
		Date: time.Now(),
		Ct:   data,
	}
	d, _ := json.Marshal(j)
	return string(d)
}
