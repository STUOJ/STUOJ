package response

import (
	"STUOJ/internal/domain/user"
	"STUOJ/pkg/utils"
)

type UserData struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Role       uint8  `json:"role"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Signature  string `json:"signature"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type UserStatistics struct {
	ACCount     int64 `json:"ac_count"`
	BlogCount   int64 `json:"blog_count"`
	SubmitCount int64 `json:"submit_count"`
}

func Map2UserStatistics(m map[string]any) UserStatistics {
	var res UserStatistics
	utils.SafeTypeAssert(m["ac_count"], &res.ACCount)
	utils.SafeTypeAssert(m["blog_count"], &res.BlogCount)
	utils.SafeTypeAssert(m["submit_count"], &res.SubmitCount)
	return res
}

type UserQueryData struct {
	UserData
	UserStatistics
}

type UserSimpleData struct {
	Id     int64  `json:"id"`
	Name   string `json:"username"`
	Role   uint8  `json:"role"`
	Avatar string `json:"avatar"`
}

func Domain2UserSimpleData(u user.User) UserSimpleData {
	return UserSimpleData{
		Avatar: u.Avatar.String(),
		Id:     u.Id.Value(),
		Name:   u.Username.String(),
		Role:   uint8(u.Role.Value()),
	}
}

func Map2UserSimpleData(u map[string]any) UserSimpleData {
	var res UserSimpleData
	utils.SafeTypeAssert(u["avatar"], &res.Avatar)
	utils.SafeTypeAssert(u["id"], &res.Id)
	utils.SafeTypeAssert(u["username"], &res.Name)
	utils.SafeTypeAssert(u["role"], &res.Role)
	return res
}
