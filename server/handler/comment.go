package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/comment"
	"STUOJ/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 发表评论
type ReqCommentAdd struct {
	BlogId  uint64 `json:"blog_id,omitempty" binding:"required"`
	Content string `json:"content,omitempty" binding:"required"`
}

func CommentAdd(c *gin.Context) {
	_, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	var req ReqCommentAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	cmt := entity.Comment{
		UserId:  uid,
		BlogId:  req.BlogId,
		Content: req.Content,
	}

	// 插入评论
	cmt.Id, err = comment.Insert(cmt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("发布成功，返回评论ID", cmt.Id))
}

// 获取评论列表
func CommentList(c *gin.Context) {
	role, userId := utils.GetUserInfo(c)
	condition := model.CommentWhere{}
	condition.Parse(c)
	commonts, err := comment.Select(condition, userId, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, model.RespOk("查询成功", commonts))
}

// 删除评论
func CommentRemove(c *gin.Context) {
	role, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 删除评论
	cid := uint64(id)
	err = comment.Delete(cid, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

// 修改评论
type ReqCommentModify struct {
	Id      uint64               `json:"id,omitempty" binding:"required"`
	UserId  uint64               `json:"user_id,omitempty" binding:"required"`
	BlogId  uint64               `json:"blog_id,omitempty" binding:"required"`
	Content string               `json:"content,omitempty" binding:"required"`
	Status  entity.CommentStatus `json:"status,omitempty" binding:"required"`
}

func CommentModify(c *gin.Context) {
	var req ReqCommentModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", err.Error()))
		return
	}

	cmt := entity.Comment{
		Id:      req.Id,
		UserId:  req.UserId,
		BlogId:  req.BlogId,
		Content: req.Content,
		Status:  req.Status,
	}

	// 修改评论
	err = comment.Update(cmt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}
