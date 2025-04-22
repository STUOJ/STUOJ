package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/language"
	"STUOJ/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取语言列表
func LanguageList(c *gin.Context) {
	role, _ := utils.GetUserInfo(c)
	con := model.LanguageWhere{}
	con.Parse(c)
	languages, err := language.Select(con, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", languages))
}

type ReqLanguageUpdate struct {
	Id     uint64                `json:"id" binding:"required"`
	Name   string                `json:"name"`
	Serial uint16                `json:"serial" binding:"required"`
	Status entity.LanguageStatus `json:"status" binding:"required"`
	MapId  uint32                `json:"map_id"`
}

func UpdateLanguage(c *gin.Context) {
	role, _ := utils.GetUserInfo(c)
	var req ReqLanguageUpdate
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", err.Error()))
		return
	}
	lang := entity.Language{
		Id:     req.Id,
		Name:   req.Name,
		Serial: req.Serial,
		Status: req.Status,
		MapId:  req.MapId,
	}
	if err := language.Update(lang, role); err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, model.RespOk("OK", nil))
}
