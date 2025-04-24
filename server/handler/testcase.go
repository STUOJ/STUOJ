package handler

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/service/testcase"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取评测点数据
func TestcaseInfo(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	t, err := testcase.SelectById(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.RespOk("OK", t))
}

func TestcaseList(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	params := request.QueryTestcaseParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	ts, err := testcase.Select(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.RespOk("OK", ts))
}

func TestcaseAdd(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.CreateTestcaseReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 插入评测点数据
	id, err := testcase.Insert(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回评测点ID", id))
}

func TestcaseModify(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UpdateTestcaseReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 更新评测点数据
	err = testcase.Update(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除评测点数据
func TestcaseRemove(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	err = testcase.Delete(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

// 生成测试用例数据
func TestcaseDataMake(c *gin.Context) {
	var t model.CommonTestcaseInput
	if err := c.ShouldBindJSON(&t); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	tc, err := t.Unfold()
	if err != nil {
		c.Error(errors.ErrInternalServer.WithMessage(err.Error()))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("OK", tc.String()))
}
