package user

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/image"
	imgval "STUOJ/internal/domain/image/valueobject"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/domain/user/valueobject"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"io"
)

// Update 根据Id更新用户
func Update(uid uint64, req request.UserUpdateReq, reqUser model.ReqUser) error {
	// 读取用户
	qt := querycontext.UserQueryContext{}
	qt.Id.Add(uid)
	qt.Field.SelectId().SelectUsername().SelectSignature()
	u0, _, err := user.Query.SelectOne(qt)
	if err != nil {
		return err
	}

	// 更新字段
	u0.Username = valueobject.Username(req.Username)
	u0.Signature = valueobject.Signature(req.Signature)

	return u0.Update(false)
}

// UpdatePassword 根据Id更新用户密码
func UpdatePassword(req request.UserForgetPasswordReq, reqUser model.ReqUser) error {
	// 读取用户
	qt := querycontext.UserQueryContext{}
	qt.Id.Add(reqUser.Id)
	qt.Field.SelectId().SelectPassword()
	u0, _, err := user.Query.SelectOne(qt)
	if err != nil {
		return err
	}

	u0.Password = valueobject.Password(req.Password)

	return u0.Update(true)
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
		return errors.ErrUnauthorized.WithMessage("没有权限修改用户权限组")
	}

	u0.Role = newRole

	return u0.Update(false)
}

// UpdateAvatar 更新用户头像
func UpdateAvatar(uid uint64, reader io.Reader, filename string, reqUser model.ReqUser) (string, error) {
	// 读取用户
	qt := querycontext.UserQueryContext{}
	qt.Id.Add(uid)
	qt.Field.SelectId().SelectAvatar()
	u0, _, err := user.Query.SelectOne(qt)
	if err != nil {
		return "", err
	}

	if u0.Id != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		return "", errors.ErrUnauthorized.WithMessage("没有权限修改其他用户的头像")
	}

	// 上传头像
	newImage := image.NewImage(
		image.WithReader(reader),
		image.WithKey(filename),
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
	u0.Avatar = valueobject.Avatar(url)
	err = u0.Update(false)
	if err != nil {
		return "", err
	}

	return url, nil
}
