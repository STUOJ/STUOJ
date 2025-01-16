package handler

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/user"
	"STUOJ/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户注册
type ReqUserRegister struct {
	Username   string       `json:"username" binding:"required,min=3,max=16"`
	Password   string       `json:"password" binding:"required,min=6,max=16"`
	Email      entity.Email `json:"email" binding:"required,max=50"`
	VerifyCode string       `json:"verify_code" binding:"required,max=10"`
}

func UserRegister(c *gin.Context) {
	var req ReqUserRegister

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	if err := utils.VerifyVerificationCode(req.Email.String(), req.VerifyCode); !err {
		c.JSON(http.StatusUnauthorized, model.RespError("验证码验证失败", nil))
		return
	}

	// 初始化用户
	u := entity.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	if err := u.Email.Verify(); err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError(err.Error(), nil))
		return
	}

	u.Id, err = user.Register(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("注册成功，返回用户ID", u.Id))
}

// 用户登录
type ReqUserLogin struct {
	Password string       `json:"password" binding:"required,min=6,max=16"`
	Email    entity.Email `json:"email" binding:"required,max=50"`
}

func UserLogin(c *gin.Context) {
	var req ReqUserLogin

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化用户
	u := entity.User{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := user.VerifyByEmail(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 登录成功，返回token
	c.JSON(http.StatusOK, model.RespOk("登录成功，返回token", token))
}

// 获取用户信息
func UserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	uid := uint64(id)
	u, err := user.SelectById(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", u))
}

// 获取当前用户id
func UserCurrentId(c *gin.Context) {
	_, id := utils.GetUserInfo(c)
	if id == 0 {
		c.JSON(http.StatusUnauthorized, model.RespError("未登录", nil))
		return
	}
	c.JSON(http.StatusOK, model.RespOk("OK", id))
}

// 修改用户信息
type ReqUserModify struct {
	Username  string       `json:"username" binding:"required,min=3,max=16"`
	Email     entity.Email `json:"email" binding:"required,max=50"`
	Signature string       `json:"signature" binding:"required,max=50"`
}

func UserModify(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	uid := uint64(id)

	role, id_ := utils.GetUserInfo(c)
	var req ReqUserModify

	// 参数绑定
	err = c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	if id_ != uid && role <= entity.RoleUser {
		c.JSON(http.StatusUnauthorized, model.RespError("权限不足", nil))
		return
	}

	// 修改用户
	u := entity.User{
		Id:        uid,
		Username:  req.Username,
		Email:     req.Email,
		Signature: req.Signature,
	}
	err = user.UpdateByIdExceptPassword(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 修改用户密码
type ReqUserChangePassword struct {
	Email      entity.Email `json:"email" binding:"required,min=3,max=16"`
	Password   string       `json:"password" binding:"required,min=6,max=16"`
	VerifyCode string       `json:"verify_code" binding:"required,max=10"`
}

func UserChangePassword(c *gin.Context) {
	role, _ := utils.GetUserInfo(c)
	var req ReqUserChangePassword

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	if role < entity.RoleAdmin {
		if err := utils.VerifyVerificationCode(req.Email.String(), req.VerifyCode); !err {
			c.JSON(http.StatusUnauthorized, model.RespError("验证码验证失败", nil))
			return
		}
	}
	u, err := user.SelectByEmail(req.Email.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 修改密码
	err = user.UpdatePasswordById(u.Id, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 修改用户头像
func ModifyUserAvatar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	uid := uint64(id)
	role, id_ := utils.GetUserInfo(c)
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("文件上传失败", nil))
		return
	}

	ext := filepath.Ext(file.Filename)

	// 保存文件
	dst := fmt.Sprintf("tmp/%s%s", utils.GetRandKey(), ext)
	log.Println(dst)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("文件解析失败", nil))
		return
	}
	defer os.Remove(dst)

	// 更新头像
	err = user.UpdateAvatarById(uid, dst, id_, role)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("更新成功", nil))
}

// 获取用户列表
func UserList(c *gin.Context) {
	condition := parseUserWhere(c)
	users, err := user.Select(condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", users))
}

// 添加普通用户
type ReqUserAdd struct {
	Username  string       `json:"username" binding:"required,min=3,max=16"`
	Password  string       `json:"password" binding:"required,min=6,max=20"`
	Email     entity.Email `json:"email" binding:"required,max=50"`
	Avatar    string       `json:"avatar" binding:"required"`
	Signature string       `json:"signature" binding:"required,max=50"`
}

func UserAdd(c *gin.Context) {
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
func UserRemove(c *gin.Context) {
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

func UserModifyRole(c *gin.Context) {
	role, _ := utils.GetUserInfo(c)
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
	err = user.UpdateRoleById(u, role)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 条件查询用户
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
