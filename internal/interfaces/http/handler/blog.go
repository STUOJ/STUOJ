package handler

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/service/blog"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BlogInfo(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	b, err := blog.SelectById(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", b))
}

func BlogList(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	params := request.QueryBlogParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	blogs, err := blog.Select(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", blogs))
}

// 保存博客
func BlogUpload(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.CreateBlogReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 插入博客
	id, err := blog.Insert(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("发布成功，返回博客ID", id))
}

// 编辑博客
func BlogEdit(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UpdateBlogReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 修改博客
	err = blog.Update(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除博客
func BlogRemove(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 删除博客
	err = blog.DeleteLogic(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

func BlogStatistics(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	params := request.BlogStatisticsParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	res, err := blog.Statistics(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, model.RespOk("OK", res))
}
