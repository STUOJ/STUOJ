package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BriefUser struct {
	Username string      `gorm:"column:user_username"`
	Role     entity.Role `gorm:"column:user_role"`
	Avatar   string      `gorm:"column:user_avatar"`
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

	tx := db.Db.Where(&entity.User{Id: id}).First(&user)
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
func SelectUsers(condition model.UserWhere) ([]entity.User, error) {
	var users []entity.User
	where := condition.GenerateWhere()

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
func CountUsers(condition model.UserWhere) (uint64, error) {
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

func briefUserSelect() []string {
	return []string{
		"tbl_user.username as user_username",
		"tbl_user.role as user_role",
		"tbl_user.avatar as user_avatar",
	}
}

func briefUserJoins(db *gorm.DB, tbl string) *gorm.DB {
	db = db.Joins(fmt.Sprintf("LEFT JOIN tbl_user ON %s.user_id = tbl_user.id", tbl))
	return db
}
