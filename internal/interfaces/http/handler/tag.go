package handler

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/service/tag"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取标签信息
func TagInfo(c *gin.Context) {
	// 获取用户信息
	reqUser := model.NewReqUser(c)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 获取标签数据
	t, err := tag.SelectById(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", t))
}

// 获取标签列表
func TagList(c *gin.Context) {
	// 获取用户信息
	reqUser := model.NewReqUser(c)

	// 解析查询参数
	params := request.QueryTagParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 查询标签列表
	tags, err := tag.Select(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", tags))
}

// 添加标签
func TagAdd(c *gin.Context) {
	// 获取用户信息
	reqUser := model.NewReqUser(c)

	var req request.CreateTagReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 插入标签
	id, err := tag.Insert(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回标签ID", id))
}

// 修改标签数据
func TagModify(c *gin.Context) {
	// 获取用户信息
	reqUser := model.NewReqUser(c)

	var req request.UpdateTagReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	err = tag.Update(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除标签
func TagRemove(c *gin.Context) {
	// 获取用户信息
	reqUser := model.NewReqUser(c)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 删除标签
	err = tag.Delete(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
