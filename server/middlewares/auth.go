package middlewares

import (
	"STUOJ/internal/app/service/user"
	"STUOJ/internal/conf"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TokenGetInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.GetToken(c) == "" {
			c.Set("id", uint64(0))
			c.Set("role", entity.RoleVisitor)
			c.Next()
			return
		}
		uid, err := utils.GetTokenUid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.RespError("获取用户id失败", nil))
			c.Abort()
			return
		}
		role, err := getUserRole(uid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.RespError("无法查询到用户组", nil))
			c.Abort()
			return
		}

		if role != entity.RoleVisitor {
			err = tokenAutoRefresh(c)
			if err != nil {
				log.Println(err)
			}
		}

		c.Set("req_user_id", uid)
		c.Set("req_user_role", role)
		c.Next()
	}
}

func TokenAuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqUser := model.NewReqUser(c)
		if reqUser.Role < entity.RoleUser {
			c.JSON(http.StatusForbidden, model.RespError("请登录", nil))
			c.Abort()
			return
		}
		// 放行
		c.Next()
	}
}

func TokenAuthEditor() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqUser := model.NewReqUser(c)
		if reqUser.Role < entity.RoleEditor {
			c.JSON(http.StatusForbidden, model.RespError("权限不足", nil))
			c.Abort()
			return
		}
		// 放行
		c.Next()
	}
}

func TokenAuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqUser := model.NewReqUser(c)
		if reqUser.Role < entity.RoleAdmin {
			c.JSON(http.StatusForbidden, model.RespError("权限不足", nil))
			c.Abort()
			return
		}

		// 放行
		c.Next()
	}
}

func TokenAuthRoot() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqUser := model.NewReqUser(c)
		if reqUser.Role < entity.RoleRoot {
			c.JSON(http.StatusForbidden, model.RespError("权限不足", nil))
			c.Abort()
			return
		}

		// 放行
		c.Next()
	}
}

func tokenAutoRefresh(c *gin.Context) error {
	config := conf.Conf.Token
	exp, err := utils.GetTokenExpire(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token无效，获取用户信息失败", nil))
		c.Abort()
		return err
	}

	// 计算token剩余时间
	timeLeft := exp - uint64(time.Now().Unix())
	//log.Println(timeLeft, config.Refresh)
	if timeLeft > config.Refresh {
		return nil
	}

	reqUser := model.NewReqUser(c)
	uid := uint64(reqUser.ID)

	// 生成新token
	token, err := utils.GenerateToken(uid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token刷新失败", nil))
		c.Abort()
		return err
	}

	c.JSON(http.StatusUnauthorized, model.RespRetry("token已刷新，请重新发送请求", token))
	c.Abort()
	return nil
}

func getUserRole(uid uint64) (entity.Role, error) {
	// 获取用户信息
	u, err := user.SelectById(uid)
	if err != nil {
		return 0, err
	}

	return u.Role, nil
}

func tokenVerify(c *gin.Context) error {
	err := utils.VerifyToken(c)
	if err != nil {
		return errors.New("token无效")
	}

	return nil
}
