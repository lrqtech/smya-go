package resource

import (
	"encoding/hex"
	"smya/config"
	"smya/util"
	"strconv"
	"time"
)

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

// 资源解密
func ResDecrypt(s []byte) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	CurrentHour := strconv.Itoa(time.Now().In(loc).Hour())
	if len(CurrentHour) == 1 {
		CurrentHour += "0"
	}
	iv := config.Key + CurrentHour
	ds, _ := hex.DecodeString(string(s))
	result, _ := util.DesDecryptCBC(ds, []byte(iv), []byte(iv), util.PKCS5PADDING)
	return string(result)
}
