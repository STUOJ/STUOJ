package middlewares

import (
	"STUOJ/internal/conf"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/user"
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

		c.Set("id", uid)
		c.Set("role", role)
		c.Next()
	}
}

func TokenAuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		switch role {
		case entity.RoleVisitor:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，未登录用户", nil))
			c.Abort()
			return
		case entity.RoleBanned:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，用户已被封禁", nil))
			c.Abort()
			return
		default:
			break
		}

		// 放行
		c.Next()
	}
}

func TokenAuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		//log.Println(role)
		switch role {
		case entity.RoleVisitor:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，未登录用户", nil))
			c.Abort()
			return
		case entity.RoleBanned:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，用户已被封禁", nil))
			c.Abort()
			return
		case entity.RoleAdmin:
			break
		case entity.RoleRoot:
			break
		default:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，用户权限不足", nil))
			c.Abort()
			return
		}

		// 放行
		c.Next()
	}
}

func TokenAuthRoot() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		switch role {
		case entity.RoleVisitor:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，未登录用户", nil))
			c.Abort()
			return
		case entity.RoleBanned:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，用户已被封禁", nil))
			c.Abort()
			return
		case entity.RoleRoot:
			break
		default:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，用户权限不足", nil))
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

	_, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)

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
		return -1, err
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
