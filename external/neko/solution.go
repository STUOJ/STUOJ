package neko

import (
	"STUOJ/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
)

// 生成测试用例
func GenerateSolution(ti model.NekoSolutionInstruction) (model.NekoSolution, error) {
	// 转换为json
	data, err := json.Marshal(ti)
	if err != nil {
		return model.NekoSolution{}, err
	}

	// 发送请求
	bodyStr, err := httpInteraction("/solution", "POST", bytes.NewReader(data))
	if err != nil {
		return model.NekoSolution{}, err
	}

	// 解析返回值
	var resp model.NekoRespObj
	err = json.Unmarshal([]byte(bodyStr), &resp)
	if err != nil {
		return model.NekoSolution{}, err
	}
	if resp.Code == 0 {
		return model.NekoSolution{}, errors.New(resp.Msg)
	}

	// 解析题解
	var s model.NekoSolution
	err = mapstructure.Decode(resp.Data, &s)
	if err != nil {
		return model.NekoSolution{}, err
	}

	return s, nil
}
