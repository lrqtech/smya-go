package request

import (
	"encoding/json"
	"smya/config"
	"smya/model"
	"smya/resource"
	"smya/util"
)

// 登入
func Login() (string, string) {
	str := resource.ResEncrypt()
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
