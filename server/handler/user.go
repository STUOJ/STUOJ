package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/user"
	"STUOJ/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户注册
type ReqUserRegister struct {
	Username   string       `json:"username" binding:"required"`
	Password   string       `json:"password" binding:"required"`
	Email      entity.Email `json:"email" binding:"required"`
	VerifyCode string       `json:"verify_code" binding:"required"`
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
	Password string       `json:"password" binding:"required"`
	Email    entity.Email `json:"email" binding:"required"`
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
	Username  string       `json:"username" binding:"required"`
	Email     entity.Email `json:"email" binding:"required"`
	Signature string       `json:"signature" binding:"required"`
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
	Password   string `json:"password" binding:"required"`
	VerifyCode string `json:"verify_code" binding:"required"`
}

func UserChangePassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	uid := uint64(id)
	role, id_ := utils.GetUserInfo(c)
	var req ReqUserChangePassword

	// 参数绑定
	err = c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	userInfo, err := user.SelectById(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	if err := utils.VerifyVerificationCode(userInfo.Email.String(), req.VerifyCode); !err {
		c.JSON(http.StatusUnauthorized, model.RespError("验证码验证失败", nil))
		return
	}

	if id_ != uid && role <= entity.RoleUser {
		c.JSON(http.StatusUnauthorized, model.RespError("权限不足", nil))
		return
	}

	// 修改密码
	err = user.UpdatePasswordById(uid, req.Password)
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

	// 保存文件
	dst := fmt.Sprintf("tmp/%s", utils.GetRandKey())
	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("文件上传失败", nil))
		return
	}

	if id_ != uid && role <= entity.RoleUser {
		c.JSON(http.StatusUnauthorized, model.RespError("权限不足", nil))
		return
	}

	// 更新头像
	err = user.UpdateAvatarById(uid, dst)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("更新成功", nil))
}
