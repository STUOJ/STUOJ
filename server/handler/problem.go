package handler

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/problem"
	"STUOJ/internal/service/tag"
	"STUOJ/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取题目信息
func ProblemInfo(c *gin.Context) {
	role, _ := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	pd, err := problem.SelectById(pid, role >= entity.RoleAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pd))
}

// 获取题目列表
func ProblemList(c *gin.Context) {
	role, _ := utils.GetUserInfo(c)

	condition := parseProblemWhere(c)

	if role < entity.RoleAdmin {
		condition.Status.Set(entity.ProblemStatusPublic)
	}

	pds, err := problem.Select(condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pds))
}

// 获取标签列表
func TagList(c *gin.Context) {
	tags, err := tag.SelectAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", tags))
}

func parseProblemWhere(c *gin.Context) dao.ProblemWhere {
	condition := dao.ProblemWhere{}

	if c.Query("title") != "" {
		condition.Title.Set(c.Query("title"))
	}
	if c.Query("difficulty") != "" {
		difficulty, err := strconv.Atoi(c.Query("difficulty"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Difficulty.Set(entity.Difficulty(difficulty))
		}
	}
	if c.Query("tag") != "" {
		tagsQuery := c.Query("tag")           // 获取URL参数 "ids"
		tags := strings.Split(tagsQuery, ",") // 将字符串分割成字符串切片

		// 假设我们需要将字符串切片转换为int切片
		var tagsInt []uint64
		for _, tagStr := range tags {
			id, err := strconv.Atoi(tagStr)
			if err != nil {
				continue
			}
			tagsInt = append(tagsInt, uint64(id))
		}
		condition.Tag.Set(tagsInt)
	}
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Status.Set(entity.ProblemStatus(status))
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
