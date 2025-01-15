package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"errors"
	"log"
)

type UserPage struct {
	Users []entity.User `json:"users"`
	model.Page
}

// 根据ID查询用户
func SelectById(id uint64) (entity.User, error) {
	u, err := dao.SelectUserById(id)
	if err != nil {
		log.Println(err)
		return entity.User{}, errors.New("用户不存在")
	}

	// 不返回密码
	u.Password = ""

	return u, nil
}

func SelectByEmail(email string) (entity.User, error) {
	u, err := dao.SelectUserByEmail(email)
	if err != nil {
		log.Println(err)
		return entity.User{}, errors.New("用户不存在")
	}

	// 不返回密码
	u.Password = ""

	return u, nil
}

// 查询所有用户
func Select(condition model.UserWhere) (UserPage, error) {
	if !condition.Page.Exist() {
		condition.Page.Set(1)
	}
	if !condition.Size.Exist() {
		condition.Size.Set(10)
	}
	users, err := dao.SelectUsers(condition)
	if err != nil {
		log.Println(err)
		return UserPage{}, errors.New("查询用户失败")
	}

	hidePassword(users)

	count, err := dao.CountUsers(condition)
	if err != nil {
		log.Println(err)
		return UserPage{}, errors.New("查询统计失败")
	}
	uPage := UserPage{
		Users: users,
		Page: model.Page{
			Total: count,
			Page:  condition.Page.Value(),
			Size:  condition.Size.Value(),
		},
	}

	return uPage, nil
}

// 不返回密码
func hidePassword(users []entity.User) {
	for i := range users {
		users[i].Password = ""
	}
}

// 根据邮箱验证密码
func VerifyByEmail(u entity.User) (string, error) {
	password := u.Password

	// 查询用户
	u, err := dao.SelectUserByEmail(u.Email.String())
	if err != nil {
		log.Println(err)
		return "", errors.New("用户不存在")
	}

	// 验证密码
	err = u.VerifyByPassword(password)
	if err != nil {
		log.Println(err)
		return "", errors.New("用户名或密码错误")
	}

	// 生成token
	token, err := utils.GenerateToken(u.Id)
	if err != nil {
		log.Println(err)
		return "", errors.New("生成token失败")
	}

	return token, nil
}
