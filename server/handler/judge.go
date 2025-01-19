package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/judge"
	"STUOJ/internal/service/language"
	"STUOJ/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取语言列表
func JudgeLanguageList(c *gin.Context) {
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

// 提交评测
type ReqJudgeSubmit struct {
	LanguageId uint64 `json:"language_id" binding:"required"`
	ProblemId  uint64 `json:"problem_id" binding:"required"`
	SourceCode string `json:"source_code" binding:"required"`
}

func JudgeSubmit(c *gin.Context) {
	_, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	var req ReqJudgeSubmit

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化提交对象
	s := entity.Submission{
		UserId:     uid,
		ProblemId:  req.ProblemId,
		LanguageId: req.LanguageId,
		SourceCode: req.SourceCode,
	}

	// 提交代码
	//s.Id, err = judge.AsyncSubmit(s)
	s.Id, err = judge.AsyncSubmit(s)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回提交ID
	c.JSON(http.StatusOK, model.RespOk("提交成功，返回记录提交ID", s.Id))

}

type ReqJudgeTestRun struct {
	LanguageId uint64 `json:"language_id" binding:"required"`
	SourceCode string `json:"source_code" binding:"required"`
	Stdin      string `json:"stdin"`
}

func JudgeTestRun(c *gin.Context) {
	var req ReqJudgeTestRun
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化提交对象
	s := entity.Submission{
		LanguageId: req.LanguageId,
		SourceCode: req.SourceCode,
	}

	// 测试运行
	j, err := judge.TestRun(s, req.Stdin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", j))
}
