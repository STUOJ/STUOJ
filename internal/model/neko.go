package model

type NekoResp struct {
	Code uint8       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
