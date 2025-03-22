package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/model"
)

type userStore struct{}

var UserStore = new(userStore)

// 插入用户
func (store *userStore) Insert(u entity.User) (uint64, error) {
	tx := db.Db.Create(&u)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return u.Id, nil
}

// 根据ID查询用户
func (store *userStore) SelectById(id uint64) (entity.User, error) {
	var u entity.User
	conditin := model.UserWhere{}

	tx := db.Db.Where(&entity.User{Id: id})
	tx = conditin.GenerateWhere()(tx)
	tx = tx.First(&u)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	return u, nil
}

// 根据邮箱查询用户
func (store *userStore) SelectByEmail(e string) (entity.User, error) {
	var user entity.User

	tx := db.Db.Where("email = ?", e).First(&user)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	return user, nil
}

// 查询用户
func (store *userStore) Select(condition model.UserWhere) ([]entity.User, error) {
	var u []entity.User
	where := condition.GenerateWhere()

	tx := db.Db.Model(&entity.User{})
	tx = where(tx)
	tx = tx.Find(&u)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return u, nil
}

// 根据ID更新用户
func (store *userStore) UpdateById(u entity.User) error {
	tx := db.Db.Model(&u).Where("id = ?", u.Id).Updates(u)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除用户
func (store *userStore) DeleteById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.User{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计用户数量
func (store *userStore) Count(condition model.UserWhere) (uint64, error) {
	var count int64
	where := condition.GenerateWhereWithNoPage()
	tx := db.Db.Model(&entity.User{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}
