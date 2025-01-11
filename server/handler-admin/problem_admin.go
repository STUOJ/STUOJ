package handler_admin

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/history"
	"STUOJ/internal/service/problem"
	"STUOJ/utils"
	"STUOJ/utils/fps"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func AdminProblemAdd(c *gin.Context) {
	_, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	var req ReqProblemAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
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

func AdminProblemModify(c *gin.Context) {
	_, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	var req ReqProblemModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
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

	err = problem.UpdateById(p, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除题目
func AdminProblemRemove(c *gin.Context) {
	_, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	err = problem.DeleteByProblemId(pid, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

// 添加标签到题目
type ReqProblemAddTag struct {
	ProblemId uint64 `json:"problem_id,omitempty" binding:"required"`
	TagId     uint64 `json:"tag_id,omitempty" binding:"required"`
}

func AdminProblemAddTag(c *gin.Context) {
	var req ReqProblemAddTag

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 添加标签
	err = problem.InsertTag(req.ProblemId, req.TagId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功", nil))
}

// 删除题目的某个标签
type ReqProblemRemoveTag struct {
	ProblemId uint64 `json:"problem_id,omitempty" binding:"required"`
	TagId     uint64 `json:"tag_id,omitempty" binding:"required"`
}

func AdminProblemRemoveTag(c *gin.Context) {
	var req ReqProblemRemoveTag

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 删除标签
	err = problem.DeleteTag(req.ProblemId, req.TagId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

func AdminProblemParseFromFps(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("文件上传失败", nil))
		return
	}

	// 保存文件
	dst := fmt.Sprintf("tmp/%s", utils.GetRandKey())
	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("文件上传失败", nil))
		return
	}
	defer os.Remove(dst)

	// 解析文件
	f, err := fps.Read(dst)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("文件解析失败", nil))
		return
	}
	p := fps.Parse(f)

	c.JSON(http.StatusOK, model.RespOk("文件解析成功", p))
}

// 获取题目历史记录
func AdminHistoryListOfProblem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	histories, err := history.SelectHistoriesByProblemId(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", histories))
}
