package model

import (
	"STUOJ/internal/infrastructure/repository/entity"

	"github.com/gin-gonic/gin"
)

type ReqUser struct {
	Id   int64
	Role entity.Role
}

func (r *ReqUser) Parse(c *gin.Context) {
	id, exist := c.Get("req_user_id")
	if !exist {
		id = 0
	}
	role, exist := c.Get("req_user_role")
	if !exist {
		role = entity.RoleVisitor
	}
	r.Id = id.(int64)
	r.Role = role.(entity.Role)
}

func NewReqUser(c *gin.Context) *ReqUser {
	r := &ReqUser{}
	r.Parse(c)
	return r
}
