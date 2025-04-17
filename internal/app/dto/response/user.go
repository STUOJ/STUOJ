package response

import "STUOJ/internal/domain/user"

type UserData struct {
	// 用户头像
	Avatar string `json:"avatar"`
	// 注册日期
	CreateTime *string `json:"create_time,omitempty"`
	// 用户id，ID 编号
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// 角色
	Role      int64   `json:"role"`
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
	// 用户id，ID 编号
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// 角色
	Role int64 `json:"role"`
}

func Domain2UserSimpleData(u user.User) UserSimpleData {
	return UserSimpleData{
		Avatar: u.Avatar.String(),
		ID:     int64(u.Id),
		Name:   u.Username.String(),
		Role:   int64(u.Role),
	}
}
