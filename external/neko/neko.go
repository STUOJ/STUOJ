package neko

import (
	"STUOJ/internal/conf"
	"STUOJ/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func InitNekoAcm() error {
	config = conf.Conf.NekoAcm
	preUrl = config.Host + ":" + config.Port + "/api"
	log.Println("Connecting to NekoACM service: " + preUrl)

	// 发送请求
	bodyStr, err := httpInteraction("/", "GET", nil)
	if err != nil {
		return err
	}

	// 解析返回值
	var resp model.NekoResp
	err = json.Unmarshal([]byte(bodyStr), &resp)
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func httpInteraction(route string, httpMethod string, reader *bytes.Reader) (string, error) {
	url := preUrl + route
	var req *http.Request
	var err error
	if reader == nil {
		req, err = http.NewRequest(httpMethod, url, nil)
	} else {
		req, err = http.NewRequest(httpMethod, url, reader)
	}
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	bodyStr := string(body)
	return bodyStr, nil
}
