package handler_admin

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/conf"
	"STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 获取设置列表
func AdminConfigList(c *gin.Context) {
	var err error
	configuration := model.Configuration{}

	configuration.System = *conf.Conf
	configuration.Judge, err = judge0.GetConfigInfo()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("获取配置信息失败", nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", configuration))
}
