package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"

	"gorm.io/gorm"
)

type UserWhere struct {
	Role model.Field[entity.Role]
	Page model.Field[uint64]
	Size model.Field[uint64]
}

// 插入用户
func InsertUser(u entity.User) (uint64, error) {
	tx := db.Db.Create(&u)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return u.Id, nil
}

// 根据ID查询用户
func SelectUserById(id uint64) (entity.User, error) {
	var user entity.User

	tx := db.Db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	return user, nil
}

// 根据邮箱查询用户
func SelectUserByEmail(e string) (entity.User, error) {
	var user entity.User

	tx := db.Db.Where("email = ?", e).First(&user)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	return user, nil
}

// 查询用户
func SelectUsers(condition UserWhere) ([]entity.User, error) {
	var users []entity.User
	where := generateUserWhereCondition(condition)

	tx := db.Db.Model(&entity.User{})
	tx = where(tx)
	tx = tx.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

// 根据ID更新用户
func UpdateUserById(u entity.User) error {
	tx := db.Db.Model(&u).Where("id = ?", u.Id).Updates(u)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除用户
func DeleteUserById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.User{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计用户数量
func CountUsers(condition UserWhere) (uint64, error) {
	var count int64
	where := generateUserWhereConditionWithNoPage(condition)
	tx := db.Db.Model(&entity.User{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

// 根据角色统计用户数量
func CountUsersGroupByRole() ([]model.CountByRole, error) {
	var counts []model.CountByRole

	tx := db.Db.Model(&entity.User{}).Select("role, count(*) as count").Group("role").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
}

// 根据创建时间统计用户数量
func CountUsersBetweenCreateTime(startTime time.Time, endTime time.Time) ([]model.CountByDate, error) {
	var counts []model.CountByDate

	tx := db.Db.Model(&entity.User{}).Where("create_time between ? and ?", startTime, endTime).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
}

func generateUserWhereConditionWithNoPage(con UserWhere) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		if con.Role.Exist() {
			whereClause["role"] = con.Role.Value()
		}
		return db.Where(whereClause)
	}
}

func generateUserWhereCondition(con UserWhere) func(*gorm.DB) *gorm.DB {
	where := generateUserWhereConditionWithNoPage(con)
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
