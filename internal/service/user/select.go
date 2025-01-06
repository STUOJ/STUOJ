package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
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
func Select(condition dao.UserWhere) (UserPage, error) {
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
