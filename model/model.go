package model

type Info struct {
	DeviceId string `json:"device_id"`
	Str      string `json:"str"`
}

type Response struct {
	Code int
	Msg  string
	Data map[string]string
}

type ResultJson struct {
	Command     string `json:"command"`
	CommandType int    `json:"type"`
	CommandName string `json:"command_name"`
}

type GetTime struct {
	CurrentHour string `json:"h"`
}
