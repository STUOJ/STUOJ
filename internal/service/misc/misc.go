package misc

import (
	"STUOJ/external/neko"
	"errors"
	"log"
)

// 生成笑话
func TellJoke() (string, error) {
	j, err := neko.TellJoke()
	if err != nil {
		log.Println(err)
		return "", errors.New("生成笑话失败！")
	}

	return j, nil
}
