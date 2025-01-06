package handler

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
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
	r, err := record.SelectBySubmissionId(id_, sid, role <= entity.RoleUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", r))
}

// 获取提交记录列表
func RecordList(c *gin.Context) {
	role, id_ := utils.GetUserInfo(c)

	condition := parseRecordWhere(c)

	records, err := record.Select(condition, id_, role <= entity.RoleUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", records))
}

func parseRecordWhere(c *gin.Context) dao.SubmissionWhere {
	condition := dao.SubmissionWhere{}
	if c.Query("problem") != "" {
		problem, err := strconv.Atoi(c.Query("problem"))
		if err != nil {
			log.Println(err)
		} else {
			condition.ProblemId.Set(uint64(problem))
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
	if c.Query("language") != "" {
		language, err := strconv.Atoi(c.Query("language"))
		if err != nil {
			log.Println(err)
		} else {
			condition.LanguageId.Set(uint64(language))
		}
	}
	timePreiod, err := utils.GetPeriod(c)
	if err != nil {
		log.Println(err)
	} else {
		condition.StartTime.Set(timePreiod.StartTime)
		condition.EndTime.Set(timePreiod.EndTime)
	}
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Status.Set(uint64(status))
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
