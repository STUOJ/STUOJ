package handler

import (
	"STUOJ/internal/dao"
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
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取题目信息
func ProblemInfo(c *gin.Context) {
	role, _ := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	pd, err := problem.SelectById(pid, role >= entity.RoleAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pd))
}

// 获取题目列表
func ProblemList(c *gin.Context) {
	role, userId := utils.GetUserInfo(c)

	condition := parseProblemWhere(c)

	pds, err := problem.Select(condition, userId, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pds))
}

// 解析题目查询条件
func parseProblemWhere(c *gin.Context) dao.ProblemWhere {
	condition := dao.ProblemWhere{}

	if c.Query("title") != "" {
		condition.Title.Set(c.Query("title"))
	}
	if c.Query("difficulty") != "" {
		difficulty, err := strconv.Atoi(c.Query("difficulty"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Difficulty.Set(entity.Difficulty(difficulty))
		}
	}
	if c.Query("tag") != "" {
		tagsQuery := c.Query("tag")           // 获取URL参数 "ids"
		tags := strings.Split(tagsQuery, ",") // 将字符串分割成字符串切片

		// 假设我们需要将字符串切片转换为int切片
		var tagsInt []uint64
		for _, tagStr := range tags {
			id, err := strconv.Atoi(tagStr)
			if err != nil {
				continue
			}
			tagsInt = append(tagsInt, uint64(id))
		}
		condition.Tag.Set(tagsInt)
	}
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Status.Set(entity.ProblemStatus(status))
		}
	}
	if c.Query("user") != "" {
		user, err := strconv.Atoi(c.Query("user"))
		if err != nil {
			log.Println(err)
		} else {
			condition.UserId.Set(uint64(user))
		}
	}
	if c.Query("page") != "" {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Page.Set(uint64(page))
		}
	}
	if c.Query("size") != "" {
		size, err := strconv.Atoi(c.Query("size"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Size.Set(uint64(size))
		}
	}
	return condition
}

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

// 修改题目
func ProblemModify(c *gin.Context) {
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
func ProblemRemove(c *gin.Context) {
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

// 从文件解析题目
func ProblemParseFromFps(c *gin.Context) {
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
func HistoryListOfProblem(c *gin.Context) {
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
