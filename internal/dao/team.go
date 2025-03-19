package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type auxiliaryTeam struct {
	entity.Team
	model.BriefUser
	TeamUserIds string `gorm:"column:team_user_id"`
}

func InsertTeam(c entity.Team) (uint64, error) {
	err := db.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Team{}).Create(&c).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return c.Id, nil
}

func SelectTeamById(id uint64) (entity.Team, error) {
	var c auxiliaryTeam

	condition := model.TeamWhere{}
	condition.Id.Set(id)

	tx := db.Db.Model(&entity.Team{})
	where := condition.GenerateWhere()
	tx = tx.Where(&entity.Team{Id: id})
	tx = where(tx)
	tx = tx.Scan(&c)

	if tx.Error != nil {
		return entity.Team{}, tx.Error
	}
	userIds := make([]uint64, 0)
	if c.TeamUserIds != "" {
		for _, idStr := range strings.Split(c.TeamUserIds, ",") {
			if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
				userIds = append(userIds, id)
			}
		}
	}
	c.Team.UserIds = userIds

	c.Team.User = entity.User{
		Id:       c.UserId,
		Username: c.Username,
		Role:     c.Role,
		Avatar:   c.Avatar,
	}

	return c.Team, nil
}

func SelectTeams(condition model.TeamWhere) ([]entity.Team, error) {
	var auxiliaryTeams []auxiliaryTeam
	var teams []entity.Team
	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Team{})
	tx = where(tx)
	tx = tx.Scan(&auxiliaryTeams)
	if tx.Error != nil {
		return nil, tx.Error
	}

	for i := range auxiliaryTeams {
		userIds := make([]uint64, 0)
		if auxiliaryTeams[i].TeamUserIds != "" {
			for _, idStr := range strings.Split(auxiliaryTeams[i].TeamUserIds, ",") {
				if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
					userIds = append(userIds, id)
				}
			}
		}
		auxiliaryTeams[i].Team.UserIds = userIds
		auxiliaryTeams[i].Team.User = entity.User{
			Id:       auxiliaryTeams[i].UserId,
			Username: auxiliaryTeams[i].Username,
			Role:     auxiliaryTeams[i].Role,
			Avatar:   auxiliaryTeams[i].Avatar,
		}
		teams = append(teams, auxiliaryTeams[i].Team)
	}

	return teams, nil
}

func UpdateTeamById(c entity.Team) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&c).Updates(c).Error
	})
}

func DeleteTeamById(id uint64) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&entity.Team{}).Error
	})
}

func CountTeams(condition model.TeamWhere) (uint64, error) {
	var count int64
	where := condition.GenerateWhereWithNoPage()

	tx := db.Db.Model(&entity.Team{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

func InsertTeamUser(tu entity.TeamUser) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.TeamUser{}).Create(&tu).Error
	})
}

func DeleteTeamUser(tu entity.TeamUser) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.TeamUser{}).Delete(&tu).Error
	})
}
