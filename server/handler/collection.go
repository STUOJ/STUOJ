package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/collection"
	"STUOJ/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取题单数据
func CollectionInfo(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	cid := uint64(id)
	coll, err := collection.SelectById(cid, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", coll))
}

// 获取题单列表
func CollectionList(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)

	condition := model.CollectionWhere{}
	condition.Parse(c)

	users, err := collection.Select(condition, uid, role)
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
	Status      entity.CollectionStatus `json:"status" binding:"required,statusRange"`
}

func CollectionAdd(c *gin.Context) {
	var req ReqCollectionAdd

	_, uid := utils.GetUserInfo(c)

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", err.Error()))
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
	Status      entity.CollectionStatus `json:"status" binding:"required,statusRange"`
}

func CollectionModify(c *gin.Context) {
	var req ReqCollectionModify

	role, uid := utils.GetUserInfo(c)

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", err.Error()))
		return
	}

	// 初始化题单
	coll := entity.Collection{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	err = collection.Update(coll, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除题单
func CollectionRemove(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	err = collection.Delete(pid, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

// 添加题目到题单
type ReqCollectionAddProblem struct {
	CollectionId uint64 `json:"collection_id" binding:"required"`
	ProblemId    uint64 `json:"problem_id" binding:"required"`
	Serial       uint16 `json:"serial" binding:"required"`
}

func CollectionAddProblem(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	var req ReqCollectionAddProblem

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 添加题单
	err = collection.InsertProblem(req.CollectionId, req.ProblemId, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功", nil))
}

type ReqCollectionModifyProblem struct {
	CollectionId uint64 `json:"collection_id" binding:"required"`
	ProblemId    uint64 `json:"problem_id" binding:"required"`
	Serial       uint16 `json:"serial" binding:"required"`
}

func CollectionModifyProblem(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	var req ReqCollectionModifyProblem

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	cp := entity.CollectionProblem{CollectionId: req.CollectionId, ProblemId: req.ProblemId, Serial: req.Serial}

	err = collection.UpdateProblem(cp, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}
}

// 删除题单的某个题目
type ReqCollectionRemoveProblem struct {
	CollectionId uint64 `json:"collection_id" binding:"required"`
	ProblemId    uint64 `json:"problem_id" binding:"required"`
}

func CollectionRemoveProblem(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	var req ReqCollectionRemoveProblem

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 删除题目
	err = collection.DeleteProblem(req.CollectionId, req.ProblemId, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
