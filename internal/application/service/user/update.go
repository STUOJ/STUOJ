package user

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/image"
	imgval "STUOJ/internal/domain/image/valueobject"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

// Update 根据Id更新用户基本信息
func Update(req request.UserUpdateReq, reqUser model.ReqUser) error {
	// 检查权限
	err := isPermission(req.Id, reqUser)
	if err != nil {
		return err
	}

	// 更新字段
	u1 := user.NewUser(
		user.WithId(req.Id),
		user.WithUsername(req.Username),
		user.WithSignature(req.Signature),
	)

	return u1.Update()
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
	err = isPermission(u0.Id.Value(), reqUser)
	if err != nil {
		return err
	}

	u1 := user.NewUser(
		user.WithId(u0.Id.Value()),
		user.WithPassword(req.Password),
	)

	return u1.Update()
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
	err = isPermission(u0.Id.Value(), reqUser)
	if err != nil {
		return err
	}

	u1 := user.NewUser(
		user.WithId(u0.Id.Value()),
		user.WithRole(newRole),
	)

	return u1.Update()
}

// UpdateAvatar 更新用户头像
func UpdateAvatar(req request.UserChangeAvatarReq, reqUser model.ReqUser) (string, error) {
	// 检查权限
	err := isPermission(req.Id, reqUser)
	if err != nil {
		return "", err
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
	// 更新头像
	u1 := user.NewUser(
		user.WithId(u0.Id.Value()),
		user.WithAvatar(url),
	)
	err = u1.Update()

	if err != nil {
		return "", err
	}

	if u0.Avatar.String() == "https://avatars.githubusercontent.com/u/188169408?s=200&v=4" {
		return url, nil
	}

	// 删除旧头像
	oldImage := image.NewImage(
		image.WithUrl(u0.Avatar.String()),
	)
	err = oldImage.Delete()
	if err != nil {
		return "", errors.ErrInternalServer.WithMessage("旧头像删除失败")
	}

	return url, nil
}
