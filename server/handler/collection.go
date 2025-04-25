package handler

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/service/collection"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取题单数据
func CollectionInfo(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	coll, err := collection.SelectById(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", coll))
}

// 获取题单列表
func CollectionList(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	params := request.QueryCollectionParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	collections, err := collection.Select(params, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", collections))
}

// 添加题单
func CollectionAdd(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.CreateCollectionReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 插入题单
	id, err := collection.Insert(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回题单ID", id))
}

// 修改题单数据
func CollectionModify(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UpdateCollectionReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	err = collection.Update(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除题单
func CollectionRemove(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	err = collection.DeleteLogic(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

func CollectionModifyProblem(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UpdateCollectionProblemReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	err = collection.UpdateProblem(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

func CollectionModifyUser(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UpdateCollectionUserReq

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	err = collection.UpdateUser(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}
