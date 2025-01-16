package user

import (
	"STUOJ/external/yuki"
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"html"
	"log"
	"strings"
	"time"
)

// 根据ID更新用户
func UpdateById(u entity.User) error {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	err := u.HashPassword()
	if err != nil {
		return err
	}

	// 读取用户
	user, err := SelectById(u.Id)
	if err != nil {
		return err
	}

	// 更新用户
	updateTime := time.Now()
	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password
	user.Signature = u.Signature
	user.Role = u.Role
	user.Avatar = u.Avatar
	user.UpdateTime = updateTime

	err = dao.UpdateUserById(user)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID更新用户（除了密码）
func UpdateByIdExceptPassword(u entity.User) error {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	// 读取用户
	user, err := SelectById(u.Id)
	if err != nil {
		log.Println(err)
		return errors.New("用户不存在")
	}

	updateTime := time.Now()
	user.Username = u.Username
	user.Email = u.Email
	user.Signature = u.Signature
	user.Role = u.Role
	user.Avatar = u.Avatar
	user.UpdateTime = updateTime

	// 更新用户
	err = dao.UpdateUserById(user)
	if err != nil {
		log.Println(err)
		return errors.New("更新用户失败，用户名或邮箱已存在")
	}

	return nil
}

// 根据ID更新用户密码
func UpdatePasswordById(uid uint64, pw string) error {
	// 读取用户
	u, err := dao.SelectUserById(uid)
	if err != nil {
		log.Println(err)
		return errors.New("用户不存在")
	}

	updateTime := time.Now()
	u.UpdateTime = updateTime

	// 预处理
	u.Password = pw
	err = u.HashPassword()
	if err != nil {
		log.Println(err)
		return errors.New("密码加密失败")
	}

	// 更新用户
	err = dao.UpdateUserById(u)
	if err != nil {
		log.Println(err)
		return errors.New("更新用户失败")
	}

	return nil
}

// 根据ID更新用户角色
func UpdateRoleById(u entity.User, role entity.Role) error {
	// 读取用户
	u0, err := SelectById(u.Id)
	if err != nil {
		return errors.New("用户不存在")
	}
	if u.Role >= role {
		return errors.New("权限不足")
	}
	if u0.Role >= role {
		return errors.New("权限不足")
	}

	updateTime := time.Now()
	u0.UpdateTime = updateTime
	u0.Role = u.Role

	// 更新用户
	err = dao.UpdateUserById(u0)
	if err != nil {
		return err
	}

	return nil
}

// 更新用户头像
func UpdateAvatarById(uid uint64, dst string, userId uint64, role entity.Role) error {
	// 读取用户
	u, err := SelectById(uid)
	if err != nil {
		log.Println(err)
		return errors.New("用户不存在")
	}

	if u.Id != userId && role < entity.RoleAdmin {
		return errors.New("权限不足")
	}

	// 上传头像
	image, err := yuki.UploadImage(dst, model.YukiAvatarAlbum)
	if err != nil {
		log.Println(err)
		return errors.New("上传失败")
	}

	// 删除旧头像
	err = yuki.DeleteOldAvatar(u.Avatar)
	if err != nil {
		log.Println(err)
	}

	updateTime := time.Now()
	u.Avatar = image.Url
	u.UpdateTime = updateTime

	// 更新用户
	err = dao.UpdateUserById(u)
	if err != nil {
		log.Println(err)
		return errors.New("更新用户头像失败")
	}

	return nil
}
