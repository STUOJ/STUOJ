package neko

import (
	"STUOJ/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
)

// 生成测试用例
func GenerateTestcase(ti model.NekoTestcaseInstruction) (model.NekoTestcase, error) {
	// 转换为json
	data, err := json.Marshal(ti)
	if err != nil {
		return model.NekoTestcase{}, err
	}

	// 发送请求
	bodyStr, err := httpInteraction("/testcase", "POST", bytes.NewReader(data))
	if err != nil {
		return model.NekoTestcase{}, err
	}

	// 解析返回值
	var resp model.NekoRespObj
	err = json.Unmarshal([]byte(bodyStr), &resp)
	if err != nil {
		return model.NekoTestcase{}, err
	}
	if resp.Code == 0 {
		return model.NekoTestcase{}, errors.New(resp.Msg)
	}

	// 解析测试用例
	var t model.NekoTestcase
	err = mapstructure.Decode(resp.Data, &t)
	if err != nil {
		return model.NekoTestcase{}, err
	}

	return t, nil
}
