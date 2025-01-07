package handler

import (
	"STUOJ/internal/dao"
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

func CommentList(c *gin.Context) {
	role, userId := utils.GetUserInfo(c)
	condition := parseCommentWhere(c)
	commonts, err := comment.Select(condition, userId, role >= entity.RoleAdmin)
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
	err = comment.DeleteByIdCheckUserId(cid, uid, role >= entity.RoleAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

func parseCommentWhere(c *gin.Context) dao.CommentWhere {
	condition := dao.CommentWhere{}

	if c.Query("user") != "" {
		user, err := strconv.Atoi(c.Query("user"))
		if err != nil {
			log.Println(err)
		} else {
			condition.UserId.Set(uint64(user))
		}
	}
	if c.Query("blog") != "" {
		blog, err := strconv.Atoi(c.Query("blog"))
		if err != nil {
			log.Println(err)
		} else {
			condition.BlogId.Set(uint64(blog))
		}
	}
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Status.Set(entity.CommentStatus(status))
		}
	}
	timePreiod, err := utils.GetPeriod(c)
	if err != nil {
		log.Println(err)
	} else {
		condition.StartTime.Set(timePreiod.StartTime)
		condition.EndTime.Set(timePreiod.EndTime)
	}
	if c.Query("page") != "" {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Page.Set(uint64(page))
		}
	}
	if c.Query("size") != "" {
		size, err := strconv.Atoi(c.Query("size"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Size.Set(uint64(size))
		}
	}
	return condition
}
