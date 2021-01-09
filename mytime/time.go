package mytime

import (
	"encoding/json"
	"smya/cache"
	"smya/model"
	"smya/util"
)

// 获取时间
func GetTime() string {
	url := "https://time.hl98.cn"
	req := util.Request("GET", url, []byte(""), 512)
	var r model.GetTime
	_ = json.Unmarshal(req, &r)
	return r.CurrentHour
}

// 设置缓存时间
func SetHour() {
	CurrentHour := GetTime()
	cache.SetCache("CurrentHour", CurrentHour)
}

func GetHour() string {
	return cache.GetCache("CurrentHour")
}
