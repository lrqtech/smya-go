package request

import (
	"encoding/json"
	"smya/model"
)

// 解析下发命令json
func GetDetail(j string) (string, int, string) {
	var r model.ResultJson
	_ = json.Unmarshal([]byte(j), &r)
	return r.Command, r.CommandType, r.CommandName
}
