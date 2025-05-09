package handler

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/interfaces/http/vo"
	"STUOJ/pkg/errors"

	"STUOJ/internal/application/service/history"
	"STUOJ/internal/application/service/problem"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取题目信息
func ProblemInfo(c *gin.Context) {
	reqUser := request.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	p, err := problem.SelectById(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", p))
}

// 获取题目列表
func ProblemList(c *gin.Context) {
	reqUser := request.NewReqUser(c)
	params := request.QueryProblemParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	ps, err := problem.Select(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", ps))
}

// 添加题目

func ProblemAdd(c *gin.Context) {
	reqUser := request.NewReqUser(c)
	var req request.CreateProblemReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 插入题目
	id, err := problem.Insert(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, vo.RespOk("添加成功，返回题目ID", id))
}

// 修改题目

// 修改题目
func ProblemModify(c *gin.Context) {
	reqUser := request.NewReqUser(c)
	var req request.UpdateProblemReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 修改题目
	err = problem.Update(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, vo.RespOk("修改成功", nil))
}

func HistoryInfo(c *gin.Context) {
	reqUser := request.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	h, err := history.SelectById(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", h))
}

// 获取题目历史记录
func HistoryListOfProblem(c *gin.Context) {
	reqUser := request.NewReqUser(c)
	params := request.QueryHistoryParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	hs, err := history.Select(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, vo.RespOk("OK", hs))
}

func ProblemStatistics(c *gin.Context) {
	reqUser := request.NewReqUser(c)
	params := request.ProblemStatisticsParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	res, err := problem.Statistics(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, vo.RespOk("OK", res))
}
