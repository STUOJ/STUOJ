package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"html"
	"log"
	"strings"
	"time"
)

// 插入用户
func Insert(u entity.User) (uint64, error) {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	err := u.HashPassword()
	if err != nil {
		return 0, errors.New("密码加密失败")
	}

	updateTime := time.Now()
	u.CreateTime = updateTime
	u.UpdateTime = updateTime

	u.Id, err = dao.InsertUser(u)
	if err != nil {
		return 0, errors.New("插入用户失败，用户名或邮箱已存在")
	}

	return u.Id, nil
}

// 插入用户（注册）
func Register(u entity.User) (uint64, error) {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	err := u.HashPassword()
	if err != nil {
		return 0, err
	}

	updateTime := time.Now()
	u.CreateTime = updateTime
	u.UpdateTime = updateTime

	u.Id, err = dao.InsertUser(u)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入用户失败，用户名或邮箱已存在")
	}

	return u.Id, nil
}
