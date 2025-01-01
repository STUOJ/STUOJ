package handler_admin

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/user"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
func UserList(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}
	condition := parseUserWhere(c)
	users, err := user.Select(condition, uint64(page), uint64(size))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", users))
}

// 添加普通用户
type ReqUserAdd struct {
	Username  string       `json:"username" binding:"required"`
	Password  string       `json:"password" binding:"required"`
	Email     entity.Email `json:"email" binding:"required"`
	Avatar    string       `json:"avatar" binding:"required"`
	Signature string       `json:"signature" binding:"required"`
}

func AdminUserAdd(c *gin.Context) {
	var req ReqUserAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化用户
	u := entity.User{
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		Avatar:    req.Avatar,
		Signature: req.Signature,
	}
	u.Id, err = user.InsertUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回用户ID", u.Id))
}

// 删除用户
func AdminUserRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	uid := uint64(id)
	err = user.DeleteById(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

// 设置用户角色
type ReqUserModifyRole struct {
	Id   uint64      `json:"id" binding:"required"`
	Role entity.Role `json:"role" binding:"required"`
}

func AdminUserModifyRole(c *gin.Context) {
	var req ReqUserModifyRole

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化用户
	u := entity.User{
		Id:   req.Id,
		Role: req.Role,
	}

	// 修改用户
	err = user.UpdateRoleById(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

func parseUserWhere(c *gin.Context) dao.UserWhere {
	condition := dao.UserWhere{}
	if c.Query("role") != "" {
		role, err := strconv.Atoi(c.Query("role"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Role.Set(entity.Role(role))
		}
	}

	return condition
}
