package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/history"
	"STUOJ/internal/service/problem"
	"STUOJ/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取题目信息
func ProblemInfo(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	where := model.ProblemWhere{}
	where.Parse(c)
	p, err := problem.SelectById(pid, uid, role, where)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}

// 获取题目列表
func ProblemList(c *gin.Context) {
	role, userId := utils.GetUserInfo(c)

	condition := model.ProblemWhere{}
	condition.Parse(c)

	ps, err := problem.Select(condition, userId, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", ps))
}

// 解析题目查询条件

// 添加题目
type ReqProblemAdd struct {
	Title        string               `json:"title" binding:"required"`
	Source       string               `json:"source"`
	Difficulty   entity.Difficulty    `json:"difficulty"`
	TimeLimit    float64              `json:"time_limit" binding:"required"`
	MemoryLimit  uint64               `json:"memory_limit" binding:"required"`
	Description  string               `json:"description" binding:"required"`
	Input        string               `json:"input" binding:"required"`
	Output       string               `json:"output" binding:"required"`
	SampleInput  string               `json:"sample_input" binding:"required"`
	SampleOutput string               `json:"sample_output" binding:"required"`
	Hint         string               `json:"hint"`
	Status       entity.ProblemStatus `json:"status"`
}

func ProblemAdd(c *gin.Context) {
	_, uid := utils.GetUserInfo(c)
	var req ReqProblemAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", err.Error()))
		return
	}

	// 初始化题目
	p := entity.Problem{
		Title:        req.Title,
		Source:       req.Source,
		Difficulty:   req.Difficulty,
		TimeLimit:    req.TimeLimit,
		MemoryLimit:  req.MemoryLimit,
		Description:  req.Description,
		Input:        req.Input,
		Output:       req.Output,
		SampleInput:  req.SampleInput,
		SampleOutput: req.SampleOutput,
		Hint:         req.Hint,
		Status:       req.Status,
	}
	p.Id, err = problem.Insert(p, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回题目ID", p.Id))
}

// 修改题目
type ReqProblemModify struct {
	Id           uint64               `json:"id" binding:"required"`
	Title        string               `json:"title" binding:"required"`
	Source       string               `json:"source"`
	Difficulty   entity.Difficulty    `json:"difficulty" binding:"required"`
	TimeLimit    float64              `json:"time_limit" binding:"required"`
	MemoryLimit  uint64               `json:"memory_limit" binding:"required"`
	Description  string               `json:"description" binding:"required"`
	Input        string               `json:"input" binding:"required"`
	Output       string               `json:"output" binding:"required"`
	SampleInput  string               `json:"sample_input" binding:"required"`
	SampleOutput string               `json:"sample_output" binding:"required"`
	Hint         string               `json:"hint"`
	Status       entity.ProblemStatus `json:"status" binding:"required"`
}

// 修改题目
func ProblemModify(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	var req ReqProblemModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", err.Error()))
		return
	}

	// 初始化题目对象
	p := entity.Problem{
		Id:           req.Id,
		Title:        req.Title,
		Source:       req.Source,
		Difficulty:   req.Difficulty,
		TimeLimit:    req.TimeLimit,
		MemoryLimit:  req.MemoryLimit,
		Description:  req.Description,
		Input:        req.Input,
		Output:       req.Output,
		SampleInput:  req.SampleInput,
		SampleOutput: req.SampleOutput,
		Hint:         req.Hint,
		Status:       req.Status,
	}

	err = problem.Update(p, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除题目
func ProblemRemove(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	err = problem.Delete(pid, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

// 获取题目历史记录
func HistoryListOfProblem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	histories, err := history.SelectByProblemId(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", histories))
}
