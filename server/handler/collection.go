package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/collection"
	"STUOJ/internal/service/problem"
	"STUOJ/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取题单列表
func CollectionList(c *gin.Context) {

	condition := model.CollectionWhere{}
	condition.Parse(c)
	users, err := collection.Select(condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", users))
}

// 添加题单
type ReqCollectionAdd struct {
	Title       string                  `json:"title" binding:"required"`
	Description string                  `json:"description" binding:"required"`
	Status      entity.CollectionStatus `json:"status" binding:"required"`
}

func CollectionAdd(c *gin.Context) {
	var req ReqCollectionAdd

	_, uid := utils.GetUserInfo(c)

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化题单
	coll := entity.Collection{
		UserId:      uid,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	// 插入题单
	coll.Id, err = collection.Insert(coll)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回题单ID", coll.Id))
}

// 修改题单数据
type ReqCollectionModify struct {
	Id          uint64                  `json:"id" binding:"required"`
	Title       string                  `json:"title" binding:"required"`
	Description string                  `json:"description" binding:"required"`
	Status      entity.CollectionStatus `json:"status" binding:"required"`
}

func CollectionModify(c *gin.Context) {
	var req ReqCollectionModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化题单
	coll := entity.Collection{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	err = collection.UpdateById(coll)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除题单
func CollectionRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 删除题单
	tid := uint64(id)
	err = collection.DeleteById(tid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

// 添加题单到题目
type ReqProblemAddCollection struct {
	ProblemId    uint64 `json:"problem_id,omitempty" binding:"required"`
	CollectionId uint64 `json:"collection_id,omitempty" binding:"required"`
}

func ProblemAddCollection(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	var req ReqProblemAddCollection

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 添加题单
	err = problem.InsertCollection(req.ProblemId, req.CollectionId, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功", nil))
}

// 删除题目的某个题单
type ReqProblemRemoveCollection struct {
	ProblemId    uint64 `json:"problem_id,omitempty" binding:"required"`
	CollectionId uint64 `json:"collection_id,omitempty" binding:"required"`
}

// 删除题目的某个题单
func ProblemRemoveCollection(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	var req ReqProblemRemoveCollection

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 删除题单
	err = problem.DeleteCollection(req.ProblemId, req.CollectionId, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
