package handler

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/service/language"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取语言列表
func LanguageList(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	params := request.QueryLanguageParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	languages, err := language.Select(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, model.RespOk("OK", languages))
}

func UpdateLanguage(c *gin.Context) {
	reqUser := model.NewReqUser(c)

	req := request.UpdateLanguageReq{}
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	if err := language.Update(req, *reqUser); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, model.RespOk("OK", nil))
}

func LanguageStatistics(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	params := request.LanguageStatisticsParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	res, err := language.Statistics(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, model.RespOk("OK", res))
}
