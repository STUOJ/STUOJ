package judge0

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
)

func InitJudge(host, port, token string) error {
	config = JudgeConf{
		Host:  host,
		Port:  port,
		Token: token,
	}
	preUrl = config.Host + ":" + config.Port
	log.Println("Connecting to judge server: " + preUrl)
	response, err := About()
	if err != nil || response.StatusCode != http.StatusOK {
		log.Println("Judge server is not available!")
		return err
	}

	log.Println("Judge server is available.")
	return nil
}

func httpInteraction(route string, httpMethod string, reader *bytes.Reader) (string, error) {
	url := preUrl + route
	var req *http.Request
	var err error
	if route == "/submissions" && httpMethod == "POST" {
		log.Println("Wait for judge0 server to finish checking...")
		url = url + "?wait=true"
	}
	if reader == nil {
		req, err = http.NewRequest(httpMethod, url, nil)
	} else {
		req, err = http.NewRequest(httpMethod, url, reader)
	}
	if err != nil {
		return "", err
	}

	req.Header.Set("X-Auth-Token", config.Token)
	req.Header.Set("X-Auth-User", config.Token)
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
	bodystr := string(body)
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		return "", errors.New(bodystr)
	}
	return bodystr, nil
}
