package core

import "encoding/json"

type ResultJson struct {
	Command     string `json:"command"`
	CommandType int    `json:"type"`
	CommandName string `json:"command_name"`
}

func GetDetail(j string) (string, int, string) {
	var r ResultJson
	_ = json.Unmarshal([]byte(j), &r)
	return r.Command, r.CommandType, r.CommandName
}
