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
	"strings"

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

// 条件查询提交记录
func parseRecordWhere(c *gin.Context) dao.SubmissionWhere {
	condition := dao.SubmissionWhere{}
	if c.Query("problem") != "" {
		problemQuery := c.Query("problem")
		problems := strings.Split(problemQuery, ",")
		var problemsInt []uint64
		for _, problem := range problems {
			problemInt, err := strconv.Atoi(problem)
			if err != nil {
				log.Println(err)
			} else {
				problemsInt = append(problemsInt, uint64(problemInt))
			}
		}
		condition.ProblemId.Set(problemsInt)
	}
	if c.Query("user") != "" {
		userQuery := c.Query("user")
		users := strings.Split(userQuery, ",")
		var usersInt []uint64
		for _, user := range users {
			userInt, err := strconv.Atoi(user)
			if err != nil {
				log.Println(err)
			} else {
				usersInt = append(usersInt, uint64(userInt))
			}
		}
		condition.UserId.Set(usersInt)
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

// 删除提交记录（提交信息+评测结果）
func RecordRemove(c *gin.Context) {
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
