package user

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"STUOJ/utils"
)

// Register 用户注册
func Register(req request.UserRegisterReq, reqUser model.ReqUser) (uint64, error) {
	u := user.NewUser(
		user.WithUsername(req.Username),
		user.WithPassword(req.Password),
		user.WithEmail(req.Email),
	)

	// 验证码校验
	if err := utils.VerifyVerificationCode(req.Email, req.VerifyCode); err != nil {
		return 0, errors.ErrUnauthorized.WithMessage(err.Error())
	}

	return u.Create()
}
