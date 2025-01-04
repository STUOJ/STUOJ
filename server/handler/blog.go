package handler

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/blog"
	"STUOJ/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BlogInfo(c *gin.Context) {
	role, userId := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	bid := uint64(id)
	b, err := blog.SelectById(bid, userId, role >= entity.RoleAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", b))
}

func BlogList(c *gin.Context) {
	role, userId := utils.GetUserInfo(c)
	condition := parseBlogWhere(c)
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}
	blogs, err := blog.Select(condition, userId, uint64(page), uint64(size), role >= entity.RoleAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", blogs))
}

// 保存博客
type ReqBlogSave struct {
	ProblemId uint64            `json:"problem_id,omitempty"`
	Title     string            `json:"title" binding:"required"`
	Content   string            `json:"content" binding:"required"`
	Status    entity.BlogStatus `json:"status,omitempty"`
}

func BlogUpload(c *gin.Context) {
	role, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	var req ReqBlogSave

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	b := entity.Blog{
		UserId:    uid,
		ProblemId: req.ProblemId,
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
	}

	// 插入博客
	b.Id, err = blog.BlogUpload(b, role >= entity.RoleAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("保存成功，需要提交审核，返回博客ID", b.Id))
}

// 编辑博客
type ReqBlogEdit struct {
	Id        uint64            `json:"id,omitempty" binding:"required"`
	ProblemId uint64            `json:"problem_id,omitempty" binding:"required"`
	Title     string            `json:"title,omitempty" binding:"required"`
	Content   string            `json:"content,omitempty" binding:"required"`
	Status    entity.BlogStatus `json:"status,omitempty" binding:"required"`
}

func BlogEdit(c *gin.Context) {
	role, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	var req ReqBlogEdit

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	b := entity.Blog{
		Id:        req.Id,
		UserId:    uid,
		ProblemId: req.ProblemId,
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
	}

	// 修改博客
	err = blog.EditByIdCheckUserId(b, role >= entity.RoleAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功，需要提交审核", nil))
}

// 提交博客
func BlogSubmit(c *gin.Context) {
	role, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	bid := uint64(id)
	err = blog.SubmitByIdCheckUserId(bid, uid, role >= entity.RoleAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("发布成功，等待管理员审核", nil))
}

// 删除博客
func BlogRemove(c *gin.Context) {
	role, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 删除博客
	bid := uint64(id)
	err = blog.DeleteByIdCheckUserId(bid, uid, role >= entity.RoleAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

func parseBlogWhere(c *gin.Context) dao.BlogWhere {
	condition := dao.BlogWhere{}
	if c.Query("title") != "" {
		condition.Title.Set(c.Query("title"))
	}
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Status.Set(entity.BlogStatus(status))
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
	if c.Query("problem") != "" {
		problem, err := strconv.Atoi(c.Query("problem"))
		if err != nil {
			log.Println(err)
		} else {
			condition.ProblemId.Set(uint64(problem))
		}
	}
	timePreiod, err := utils.GetPeriod(c)
	if err != nil {
		log.Println(err)
	} else {
		condition.StartTime.Set(timePreiod.StartTime)
		condition.EndTime.Set(timePreiod.EndTime)
	}

	return condition
}
