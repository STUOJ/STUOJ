package handler

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/service/user"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var req request.UserRegisterReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	id, err := user.Register(req, model.ReqUser{})
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("注册成功，返回用户ID", id))
}

func UserLogin(c *gin.Context) {
	var req request.UserLoginReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	token, err := user.LoginByEmail(req, model.ReqUser{})
	if err != nil {
		c.Error(err)
		return
	}

	// 登录成功，返回token
	c.JSON(http.StatusOK, model.RespOk("登录成功，返回token", token))
}

// 获取用户信息
func UserInfo(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	u, err := user.SelectById(int64(id), *reqUser)

	c.JSON(http.StatusOK, model.RespOk("OK", u))
}

// 获取当前用户id
func UserCurrentId(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	if reqUser.Id == 0 {
		c.Error(errors.ErrUnauthorized.WithMessage("未登录"))
		return
	}
	c.JSON(http.StatusOK, model.RespOk("OK", reqUser.Id))
}

func UserModify(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UserUpdateReq
	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	err = user.Update(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

func UserChangePassword(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UserForgetPasswordReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}

	// 修改密码
	err = user.UpdatePassword(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 修改用户头像
func ModifyUserAvatar(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UserChangeAvatarReq

	// 参数绑定
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	req.File, err = c.FormFile("file")
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	url, err := user.UpdateAvatar(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("更新成功", url))
}

// 获取用户列表
func UserList(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.QueryUserParams

	// 参数绑定
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	// 查询用户
	users, err := user.Select(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", users))
}

func UserAdd(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.CreateUserReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	req_ := request.UserRegisterReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	id, err := user.Register(req_, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回用户ID", id))
}

// 删除用户
func UserRemove(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	err = user.Delete(int64(id), *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}

func UserModifyRole(c *gin.Context) {
	reqUser := model.NewReqUser(c)
	var req request.UserUpdateRoleReq

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.Error(&errors.ErrValidation)
		return
	}
	err = user.UpdateRole(req, *reqUser)
	if err != nil {
		c.Error(err)
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}
