package handler

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/service/comment"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommentAdd(c *gin.Context) {
	reqUser := model.NewReqUser(c)

	var req request.CreateCommentReq
	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 插入评论
	id, err := comment.Insert(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("发布成功，返回评论ID", id))
}

// 获取评论列表
func CommentList(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	params := request.QueryCommentParams{}
	// 参数绑定
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 获取评论列表
	commonts, err := comment.Select(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("查询成功", commonts))
}

func CommentModify(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UpdateCommentReq
	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 修改评论
	err = comment.Update(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

func CommentStatistics(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	params := request.CommentStatisticsParams{}
	// 参数绑定
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	res, err := comment.Statistics(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("查询成功", res))
}
