package handler

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/service/testcase"
	"STUOJ/internal/interfaces/http/vo"
	"STUOJ/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取评测点数据
func TestcaseInfo(c *gin.Context) {
	reqUser := request.NewReqUser(c)
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
	c.JSON(http.StatusOK, vo.RespOk("OK", t))
}

func TestcaseList(c *gin.Context) {
	reqUser := request.NewReqUser(c)
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
	c.JSON(http.StatusOK, vo.RespOk("OK", ts))
}

func TestcaseAdd(c *gin.Context) {
	reqUser := request.NewReqUser(c)
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
	c.JSON(http.StatusOK, vo.RespOk("添加成功，返回评测点ID", id))
}

func TestcaseModify(c *gin.Context) {
	reqUser := request.NewReqUser(c)
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
	c.JSON(http.StatusOK, vo.RespOk("修改成功", nil))
}

// 删除评测点数据
func TestcaseRemove(c *gin.Context) {
	reqUser := request.NewReqUser(c)
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
	c.JSON(http.StatusOK, vo.RespOk("删除成功", nil))
}

// 生成测试用例数据
func TestcaseDataMake(c *gin.Context) {
	var t vo.CommonTestcaseInput
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
	c.JSON(http.StatusOK, vo.RespOk("OK", tc.String()))
}
