package model

import (
	"STUOJ/internal/infrastructure/repository/entity"

	"github.com/gin-gonic/gin"
)

type ReqUser struct {
	Id   int64
	Role entity.Role
}

// Parse 从gin.Context中解析用户信息
// 如果字段不存在或类型不匹配，会设置默认值
func (r *ReqUser) Parse(c *gin.Context) {
	id, exist := c.Get("req_user_id")
	if !exist {
		id = 0
	}
	if v, ok := id.(int64); ok {
		r.Id = v
	} else {
		r.Id = 0
	}

	role, exist := c.Get("req_user_role")
	if !exist {
		role = entity.RoleVisitor
	}
	if v, ok := role.(entity.Role); ok {
		r.Role = v
	} else {
		r.Role = entity.RoleVisitor
	}
}

func NewReqUser(c *gin.Context) *ReqUser {
	r := &ReqUser{}
	r.Parse(c)
	return r
}
