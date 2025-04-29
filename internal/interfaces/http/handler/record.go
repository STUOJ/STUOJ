package handler

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/service/record"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取提交记录信息（提交信息+评测结果）
func RecordInfo(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	sid := uint64(id)
	r, err := record.SelectById(int64(sid), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", r))
}

// 获取提交记录列表
func RecordList(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	params := request.QuerySubmissionParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	records, err := record.Select(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", records))
}

// 获取通过用户列表
func SelectACUsers(c *gin.Context) {
	pidQuery := c.Query("problem")
	pid, err := strconv.Atoi(pidQuery)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	sizeQuery := c.Query("size")
	size, err := strconv.Atoi(sizeQuery)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	users, err := record.SelectAcUsers(int64(pid), int64(size))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", users))
}
