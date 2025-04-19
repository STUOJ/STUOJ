package model

import (
	"STUOJ/internal/db/entity"

	"github.com/gin-gonic/gin"
)

type ReqUser struct {
	Id   int64
	Role entity.Role
}

func (r *ReqUser) Parse(c *gin.Context) {
	role, exist := c.Get("req_user_id")
	if !exist {
		role = entity.RoleVisitor
	}
	id, exist := c.Get("req_user_role")
	if !exist {
		id = 0
	}
	r.Id = id.(int64)
	r.Role = role.(entity.Role)
}

func NewReqUser(c *gin.Context) *ReqUser {
	r := &ReqUser{}
	r.Parse(c)
	return r
}
