package handler

import (
	"STUOJ/internal/interfaces/http/vo"
	"STUOJ/pkg/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取设置列表
func ConfigList(c *gin.Context) {
	var err error
	configuration := vo.Configuration{}

	configuration.System = *config.Conf
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, vo.RespError("获取配置信息失败", nil))
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", configuration))
}
