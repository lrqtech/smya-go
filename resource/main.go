package resource

import (
	"encoding/hex"
	"smya/config"
	"smya/mytime"
	"smya/util"
)

// 数据加密
func ResEncrypt() string {
	CurrentHour := mytime.GetHour()
	k := config.Key + CurrentHour
	s := config.Key + config.DeviceId
	result, _ := util.DesEncryptCBC([]byte(s), []byte(k), []byte(k), util.PKCS5PADDING)
	return hex.EncodeToString(result)
}

// 资源解密
func ResDecrypt(s []byte) string {
	CurrentHour := mytime.GetHour()
	iv := config.Key + CurrentHour
	ds, _ := hex.DecodeString(string(s))
	result, _ := util.DesDecryptCBC(ds, []byte(iv), []byte(iv), util.PKCS5PADDING)
	return string(result)
}
