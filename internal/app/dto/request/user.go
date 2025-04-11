package request

type QueryUserParams struct {
	EndTime   *string `form:"end-time,omitempty"`
	ID        *string `form:"id,omitempty"`
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
