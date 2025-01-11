package handler_admin

import (
	"STUOJ/internal/model"
	"STUOJ/internal/service/record"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 删除提交记录（提交信息+评测结果）
func AdminRecordRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	sid := uint64(id)
	err = record.DeleteBySubmissionId(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
