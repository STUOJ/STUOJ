package response

import "STUOJ/internal/domain/user"

type UserData struct {
	// 用户头像
	Avatar string `json:"avatar"`
	// 注册日期
	CreateTime *string `json:"create_time,omitempty"`
	// 用户id，Id 编号
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	// 角色
	Role      uint8   `json:"role"`
	Signature *string `json:"signature,omitempty"`
	// 更新日期
	UpdateTime *string `json:"update_time,omitempty"`
}

type UserStatistics struct {
	ACCount     int64  `json:"ac_count"`
	BlogCount   string `json:"blog_count"`
	SubmitCount string `json:"submit_count"`
}

type UserQueryData struct {
	UserData
	UserStatistics
}

type UserSimpleData struct {
	// 用户头像
	Avatar string `json:"avatar"`
	// 用户id，Id 编号
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	// 角色
	Role uint8 `json:"role"`
}

func Domain2UserSimpleData(u user.User) UserSimpleData {
	return UserSimpleData{
		Avatar: u.Avatar.String(),
		Id:     u.Id,
		Name:   u.Username.String(),
		Role:   uint8(u.Role),
	}
}
