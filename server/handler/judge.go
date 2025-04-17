package handler

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/service/judge"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"

	"net/http"

	"github.com/gin-gonic/gin"
)

// 提交评测
func JudgeSubmit(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.JudgeReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	id, err := judge.Submit(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回提交ID
	c.JSON(http.StatusOK, model.RespOk("提交成功，返回记录提交ID", id))
}

func JudgeTestRun(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.TestRunReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 测试运行
	j, err := judge.TestRun(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", j))
}
