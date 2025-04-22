package handler

import (
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
	b, err := blog.SelectById(bid, userId, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", b))
}

func BlogList(c *gin.Context) {
	role, userId := utils.GetUserInfo(c)
	condition := model.BlogWhere{}
	condition.Parse(c)
	blogs, err := blog.Select(condition, userId, role)
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
	Status    entity.BlogStatus `json:"status,omitempty" binding:"required"`
}

func BlogUpload(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	var req ReqBlogSave

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", err.Error()))
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
	b.Id, err = blog.Upload(b, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("发布成功，返回博客ID", b.Id))
}

// 编辑博客
type ReqBlogEdit struct {
	Id        uint64            `json:"id,omitempty" binding:"required"`
	ProblemId uint64            `json:"problem_id,omitempty"`
	Title     string            `json:"title,omitempty" binding:"required"`
	Content   string            `json:"content,omitempty" binding:"required"`
	Status    entity.BlogStatus `json:"status,omitempty" binding:"required"`
}

func BlogEdit(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	var req ReqBlogEdit

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", err.Error()))
		return
	}

	b := entity.Blog{
		Id:        req.Id,
		ProblemId: req.ProblemId,
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
	}

	// 修改博客
	err = blog.Update(b, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除博客
func BlogRemove(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 删除博客
	bid := uint64(id)
	err = blog.Delete(bid, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
