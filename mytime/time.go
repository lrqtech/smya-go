package mytime

import (
	"encoding/json"
	"smya/model"
	"smya/util"
)

func GetTime() string {
	url := "https://time.hl98.cn"
	req := util.Request("GET", url, []byte(""), 512)
	var r model.GetTime
	_ = json.Unmarshal(req, &r)
	return r.CurrentHour
}
