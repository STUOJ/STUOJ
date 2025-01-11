package handler_admin

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/testcase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取评测点数据
func AdminTestcaseInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 获取评测点数据
	tid := uint64(id)
	t, err := testcase.SelectById(tid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.RespOk("OK", t))
}

// 添加评测点数据
type ReqTestcaseAdd struct {
	Serial     uint64 `json:"serial,omitempty" binding:"required"`
	ProblemId  uint64 `json:"problem_id,omitempty" binding:"required"`
	TestInput  string `json:"test_input,omitempty" binding:"required"`
	TestOutput string `json:"test_output,omitempty" binding:"required"`
}

func AdminTestcaseAdd(c *gin.Context) {
	var req ReqTestcaseAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化题目
	t := entity.Testcase{
		Serial:     req.Serial,
		ProblemId:  req.ProblemId,
		TestInput:  req.TestInput,
		TestOutput: req.TestOutput,
	}

	// 插入评测点数据
	t.Id, err = testcase.Insert(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回评测点ID", t.Id))
}

// 修改评测点数据
type ReqTestcaseModify struct {
	Id         uint64 `json:"id,omitempty" binding:"required"`
	Serial     uint64 `json:"serial,omitempty" binding:"required"`
	ProblemId  uint64 `json:"problem_id,omitempty" binding:"required"`
	TestInput  string `json:"test_input,omitempty" binding:"required"`
	TestOutput string `json:"test_output,omitempty" binding:"required"`
}

func AdminTestcaseModify(c *gin.Context) {
	var req ReqTestcaseModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 修改评测点数据
	t := entity.Testcase{
		Id:         req.Id,
		Serial:     req.Serial,
		ProblemId:  req.ProblemId,
		TestInput:  req.TestInput,
		TestOutput: req.TestOutput,
	}

	// 更新评测点数据
	err = testcase.UpdateById(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除评测点数据
func AdminTestcaseRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 删除评测点
	tid := uint64(id)
	err = testcase.DeleteById(tid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

func AdminTestcaseDataMake(c *gin.Context) {
	var t model.CommonTestcaseInput
	if err := c.ShouldBindJSON(&t); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	tc, err := t.Unfold()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("OK", tc.String()))
}
