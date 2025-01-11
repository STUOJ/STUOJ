package handler_admin

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/comment"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 修改评论
type ReqCommentModify struct {
	Id      uint64               `json:"id,omitempty" binding:"required"`
	UserId  uint64               `json:"user_id,omitempty" binding:"required"`
	BlogId  uint64               `json:"blog_id,omitempty" binding:"required"`
	Content string               `json:"content,omitempty" binding:"required"`
	Status  entity.CommentStatus `json:"status,omitempty"`
}

func AdminCommentModify(c *gin.Context) {
	var req ReqCommentModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
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
	err = comment.UpdateById(cmt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}
