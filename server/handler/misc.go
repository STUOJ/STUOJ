package handler

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/service/image"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UploadImageReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	req.File, err = c.FormFile("file")
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	url, err := image.Insert(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, model.RespOk("上传成功", url))
}

type ReqEmail struct {
	Email string `json:"email" binding:"required"`
}

func SendVerificationCode(c *gin.Context) {
	var req ReqEmail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	if err := utils.SendVerificationCode(req.Email); err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, model.RespOk("OK", nil))
}

type ReqVerify struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
