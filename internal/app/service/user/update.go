package user

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/image"
	imgval "STUOJ/internal/domain/image/valueobject"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

// Update 根据Id更新用户基本信息
func Update(req request.UserUpdateReq, reqUser model.ReqUser) error {
	// 检查权限
	if reqUser.Id != req.Id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}

	// 更新字段
	u1 := user.NewUser(
		user.WithId(req.Id),
		user.WithUsername(req.Username),
		user.WithSignature(req.Signature),
	)

	return u1.Update(false)
}

// UpdatePassword 根据Email更新用户密码
func UpdatePassword(req request.UserForgetPasswordReq, reqUser model.ReqUser) error {
	// 读取用户
	qt := querycontext.UserQueryContext{}
	qt.Email.Set(req.Email)
	qt.Field.SelectId().SelectPassword()
	u0, _, err := user.Query.SelectOne(qt)
	if err != nil {
		return err
	}

	// 检查权限
	if reqUser.Id != u0.Id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}

	u1 := user.NewUser(
		user.WithId(u0.Id),
		user.WithPassword(req.Password),
	)

	return u1.Update(true)
}

// UpdateRole 根据Id更新用户权限组
func UpdateRole(req request.UserUpdateRoleReq, reqUser model.ReqUser) error {
	// 读取用户
	qt := querycontext.UserQueryContext{}
	qt.Id.Add(req.Id)
	qt.Field.SelectId().SelectRole()
	u0, _, err := user.Query.SelectOne(qt)
	if err != nil {
		return err
	}

	newRole := entity.Role(req.Role)

	// 检查权限
	if u0.Role >= reqUser.Role || newRole >= reqUser.Role {
		return &errors.ErrUnauthorized
	}

	u1 := user.NewUser(
		user.WithId(u0.Id),
		user.WithRole(newRole),
	)

	return u1.Update(false)
}

// UpdateAvatar 更新用户头像
func UpdateAvatar(req request.UserChangeAvatarReq, reqUser model.ReqUser) (string, error) {
	if req.Id != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		return "", &errors.ErrUnauthorized
	}
	// 读取用户
	qt := querycontext.UserQueryContext{}
	qt.Id.Add(req.Id)
	qt.Field.SelectId().SelectAvatar()
	u0, _, err := user.Query.SelectOne(qt)
	if err != nil {
		return "", err
	}

	reader, err := req.File.Open()
	if err != nil {
		return "", errors.ErrInternalServer.WithMessage("头像读取失败")
	}
	// 上传头像
	newImage := image.NewImage(
		image.WithReader(reader),
		image.WithKey(req.File.Filename),
		image.WithAlbum(uint8(imgval.Avatar)),
	)
	url, err := newImage.Upload()
	if err != nil {
		return "", errors.ErrInternalServer.WithMessage("头像上传失败")
	}

	// 删除旧头像
	oldImage := image.NewImage(
		image.WithUrl(u0.Avatar.String()),
	)
	err = oldImage.Delete()
	if err != nil {
		return "", errors.ErrInternalServer.WithMessage("旧头像删除失败")
	}

	// 更新头像
	u1 := user.NewUser(
		user.WithId(u0.Id),
		user.WithAvatar(url),
	)
	err = u1.Update(false)
	if err != nil {
		return "", err
	}

	return url, nil
}
