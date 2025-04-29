package handler

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/service/solution"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取题解数据
func SolutionInfo(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	s, err := solution.SelectById(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", s))
}

func SolutionList(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	params := request.QuerySolutionParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	solutions, err := solution.Select(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, model.RespOk("OK", solutions))
}

func SolutionAdd(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.CreateSolutionReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 插入题解
	id, err := solution.Insert(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回题解ID", id))
}

// 修改题解
func SolutionModify(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UpdateSolutionReq
	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	err = solution.Update(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除题解
func SolutionRemove(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 删除题解
	err = solution.Delete(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
