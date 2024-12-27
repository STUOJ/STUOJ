package neko

import (
	"STUOJ/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
)

// 生成题目
func GenerateProblem(pi model.NekoProblemInstruction) (model.NekoProblem, error) {
	// 转换为json
	data, err := json.Marshal(pi)
	if err != nil {
		return model.NekoProblem{}, err
	}

	// 发送请求
	bodyStr, err := httpInteraction("/problem", "POST", bytes.NewReader(data))
	if err != nil {
		return model.NekoProblem{}, err
	}

	// 解析返回值
	var resp model.NekoResponse
	err = json.Unmarshal([]byte(bodyStr), &resp)
	if err != nil {
		return model.NekoProblem{}, err
	}
	if resp.Code == 0 {
		return model.NekoProblem{}, errors.New(resp.Msg)
	}

	// 解析题目
	var p model.NekoProblem
	err = mapstructure.Decode(resp.Data, &p)
	if err != nil {
		return model.NekoProblem{}, err
	}

	return p, nil
}

// 翻译题目
func TranslateProblem(pi model.NekoTranslateInstruction) (model.NekoProblem, error) {
	// 转换为json
	data, err := json.Marshal(pi)
	if err != nil {
		return model.NekoProblem{}, err
	}

	// 发送请求
	bodyStr, err := httpInteraction("/translate", "POST", bytes.NewReader(data))
	if err != nil {
		return model.NekoProblem{}, err
	}

	// 解析返回值
	var resp model.NekoResponse
	err = json.Unmarshal([]byte(bodyStr), &resp)
	if err != nil {
		return model.NekoProblem{}, err
	}
	if resp.Code == 0 {
		return model.NekoProblem{}, errors.New(resp.Msg)
	}

	// 解析题目
	var p model.NekoProblem
	err = mapstructure.Decode(resp.Data, &p)
	if err != nil {
		return model.NekoProblem{}, err
	}

	return p, nil
}
