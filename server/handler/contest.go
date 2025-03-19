package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/contest"
	"STUOJ/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取比赛数据
func ContestInfo(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	cid := uint64(id)
	ct, err := contest.SelectById(cid, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", ct))
}

// 获取比赛列表
func ContestList(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)

	condition := model.ContestWhere{}
	condition.Parse(c)

	users, err := contest.Select(condition, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", users))
}

// 添加比赛
type ReqContestAdd struct {
	CollectionId uint64               `json:"collection_id" binding:"required"`
	Format       entity.ContestFormat `json:"format" binding:"required"`
	TeamSize     uint8                `json:"team_size" binding:"required"`
	StartTime    string               `json:"start_time" binding:"required"`
	EndTime      string               `json:"end_time" binding:"required"`
}

func ContestAdd(c *gin.Context) {
	var req ReqContestAdd

	_, uid := utils.GetUserInfo(c)

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 时间格式转换
	startTime, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("时间格式错误", nil))
		return
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("时间格式错误", nil))
		return
	}

	// 初始化比赛
	ct := entity.Contest{
		UserId:       uid,
		CollectionId: req.CollectionId,
		Format:       req.Format,
		TeamSize:     req.TeamSize,
		StartTime:    startTime,
		EndTime:      endTime,
	}

	// 插入题单
	ct.Id, err = contest.Insert(ct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回比赛ID", ct.Id))
}

// 修改比赛数据
type ReqContestModify struct {
	Id           uint64               `json:"id" binding:"required"`
	CollectionId uint64               `json:"collection_id" binding:"required"`
	Status       entity.ContestStatus `json:"status" binding:"required"`
	Format       entity.ContestFormat `json:"format" binding:"required"`
	TeamSize     uint8                `json:"team_size" binding:"required"`
	StartTime    string               `json:"start_time" binding:"required"`
	EndTime      string               `json:"end_time" binding:"required"`
}

func ContestModify(c *gin.Context) {
	var req ReqContestModify

	role, uid := utils.GetUserInfo(c)

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 时间格式转换
	startTime, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("时间格式错误", nil))
		return
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("时间格式错误", nil))
		return
	}

	// 初始化题单
	ct := entity.Contest{
		Id:           req.Id,
		UserId:       uid,
		CollectionId: req.CollectionId,
		Status:       req.Status,
		Format:       req.Format,
		TeamSize:     req.TeamSize,
		StartTime:    startTime,
		EndTime:      endTime,
	}

	err = contest.Update(ct, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除比赛
func ContestRemove(c *gin.Context) {
	role, uid := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	err = contest.Delete(pid, uid, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
