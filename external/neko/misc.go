package neko

import (
	"STUOJ/internal/model"
	"encoding/json"
	"errors"
)

// 翻译题目
func TellJoke() (string, error) {
	// 发送请求
	bodyStr, err := httpInteraction("/joke", "GET", nil)
	if err != nil {
		return "", err
	}

	// 解析返回值
	var resp model.NekoRespStr
	err = json.Unmarshal([]byte(bodyStr), &resp)
	if err != nil {
		return "", err
	}
	if resp.Code == 0 {
		return "", errors.New(resp.Msg)
	}

	return resp.Data, nil
}
