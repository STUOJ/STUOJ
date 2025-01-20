package handler

import (
	"STUOJ/internal/model"
	"STUOJ/internal/service/record"
	"STUOJ/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取提交记录信息（提交信息+评测结果）
func RecordInfo(c *gin.Context) {
	role, id_ := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	sid := uint64(id)
	r, err := record.SelectBySubmissionId(id_, sid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", r))
}

// 获取提交记录列表
func RecordList(c *gin.Context) {
	role, id_ := utils.GetUserInfo(c)

	condition := model.SubmissionWhere{}
	condition.Parse(c)

	records, err := record.Select(condition, id_, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", records))
}

// 删除提交记录（提交信息+评测结果）
func RecordRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	sid := uint64(id)
	err = record.DeleteSubmission(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
