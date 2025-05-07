package response

import "STUOJ/internal/domain/user"

type UserData struct {
	Id         int64  `json:"id"`
	Username   string `json:"name"`
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
	return UserStatistics{
		ACCount:     m["ac_count"].(int64),
		BlogCount:   m["blog_count"].(int64),
		SubmitCount: m["submit_count"].(int64),
	}
}

type UserQueryData struct {
	UserData
	UserStatistics
}

type UserSimpleData struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
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
	return UserSimpleData{
		Avatar: u["avatar"].(string),
		Id:     u["id"].(int64),
		Name:   u["name"].(string),
		Role:   uint8(u["role"].(int64)),
	}
}
