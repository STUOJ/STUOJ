package user

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/pkg/errors"
	"STUOJ/pkg/utils"
)

// Register 用户注册
func Register(req request.UserRegisterReq, reqUser request.ReqUser) (int64, error) {
	u := user.NewUser(
		user.WithUsername(req.Username),
		user.WithPassword(req.Password),
		user.WithEmail(req.Email),
	)

	if reqUser.Role < entity.RoleAdmin {
		// 验证码校验
		if err := utils.VerifyVerificationCode(req.Email, req.VerifyCode); err != nil {
			return 0, errors.ErrUnauthorized.WithMessage(err.Error())
		}
	}

	return u.Create()
}
