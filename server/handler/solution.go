package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/solution"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取题解数据
func SolutionInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 获取评测点数据
	sid := uint64(id)
	s, err := solution.SelectById(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", s))
}

// 添加题解
type ReqSolutionAdd struct {
	LanguageId uint64 `json:"language_id" binding:"required"`
	ProblemId  uint64 `json:"problem_id" binding:"required"`
	SourceCode string `json:"source_code" binding:"required"`
}

func SolutionAdd(c *gin.Context) {
	var req ReqSolutionAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	s := entity.Solution{
		LanguageId: req.LanguageId,
		ProblemId:  req.ProblemId,
		SourceCode: req.SourceCode,
	}

	// 插入题解
	s.Id, err = solution.Insert(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回题解ID", s.Id))
}

// 修改题解
type ReqSolutionModify struct {
	Id         uint64 `json:"id,omitempty" binding:"required"`
	LanguageId uint64 `json:"language_id,omitempty" binding:"required"`
	ProblemId  uint64 `json:"problem_id,omitempty" binding:"required"`
	SourceCode string `json:"source_code,omitempty" binding:"required"`
}

func SolutionModify(c *gin.Context) {
	var req ReqSolutionModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 修改题解
	s := entity.Solution{
		Id:         req.Id,
		LanguageId: req.LanguageId,
		ProblemId:  req.ProblemId,
		SourceCode: req.SourceCode,
	}
	err = solution.UpdateById(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除题解
func SolutionRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 删除题解
	sid := uint64(id)
	err = solution.DeleteById(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
