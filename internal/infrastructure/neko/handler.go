package neko

import (
	"STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strings"
)

func ForwardHandler(c *gin.Context) {
	var err error
	url := preUrl + strings.Replace(c.Request.URL.Path, "/ai", "", 1)

	log.Println("NekoACM 请求转发到: " + url)
	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("请求失败！", nil))
		return
	}

	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("请求失败！", nil))
		return
	}
	defer res.Body.Close()

	_, err = io.Copy(c.Writer, res.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("请求失败！", nil))
		return
	}
}
