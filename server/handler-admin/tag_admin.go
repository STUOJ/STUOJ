package handler_admin

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/tag"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取标签列表
func AdminTagList(c *gin.Context) {
	tags, err := tag.SelectAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.RespOk("OK", tags))
}

// 添加标签
type ReqTagAdd struct {
	Name string `json:"name,omitempty" binding:"required"`
}

func AdminTagAdd(c *gin.Context) {
	var req ReqTagAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化标签
	t := entity.Tag{
		Name: req.Name,
	}

	// 插入标签
	t.Id, err = tag.Insert(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回标签ID", t.Id))
}

// 修改标签数据
type ReqTagModify struct {
	Id   uint64 `json:"id,omitempty" binding:"required"`
	Name string `json:"name,omitempty" binding:"required"`
}

func AdminTagModify(c *gin.Context) {
	var req ReqTagModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	err = tag.UpdateById(req.Id, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除标签
func AdminTagRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 删除标签
	tid := uint64(id)
	err = tag.DeleteById(tid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
