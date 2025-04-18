package request

type QueryUserParams struct {
	EndTime   *string `form:"end-time,omitempty"`
	Id        *string `form:"id,omitempty"`
	Order     *string `form:"order,omitempty"`
	OrderBy   *string `form:"order_by,omitempty"`
	Page      *int64  `form:"page,omitempty"`
	Role      *string `form:"role,omitempty"`
	Size      *int64  `form:"size,omitempty"`
	StartTime *string `form:"start-time,omitempty"`
}

type UserRegisterReq struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Username   string `json:"username"`
	VerifyCode string `json:"verify_code"`
}

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserUpdateReq struct {
	Username  string `json:"username"`
	Signature string `json:"signature"`
}

type UserChangeEmailReq struct {
	Email      string `json:"email"`
	VerifyCode string `json:"verify_code"`
}

type UserForgetPasswordReq struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	VerifyCode string `json:"verify_code"`
}

type UserUpdateRoleReq struct {
	Id   int64 `json:"id"`
	Role int8  `json:"role"`
}
