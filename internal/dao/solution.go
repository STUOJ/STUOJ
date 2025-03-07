package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
)

// 插入题解
func InsertSolution(s entity.Solution) (uint64, error) {
	tx := db.Db.Create(&s)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return s.Id, nil
}

// 根据ID查询题解
func SelectSolutionById(id uint64) (entity.Solution, error) {
	var s entity.Solution

	condition := model.SolutionWhere{}

	tx := db.Db.Where(&entity.Solution{Id: id})
	tx = condition.GenerateWhere()(tx)
	tx = tx.First(&s)
	if tx.Error != nil {
		return entity.Solution{}, tx.Error
	}

	return s, nil
}

// 查询所有题解
func SelectAllSolutions() ([]entity.Solution, error) {
	var solutions []entity.Solution

	condition := model.SolutionWhere{}

	tx := db.Db.Model(&entity.Solution{})
	tx = condition.GenerateWhere()(tx)
	tx = tx.Find(&solutions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return solutions, nil
}

// 根据题目ID查询题解
func SelectSolutionsByProblemId(pid uint64) ([]entity.Solution, error) {
	var solutions []entity.Solution

	tx := db.Db.Where("problem_id = ?", pid).Find(&solutions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return solutions, nil
}

// 根据ID更新题解
func UpdateSolutionById(s entity.Solution) error {
	tx := db.Db.Model(&s).Updates(s)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除题解
func DeleteSolutionById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Solution{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计题解数量
func CountSolutions() (int64, error) {
	var count int64

	tx := db.Db.Model(&entity.Solution{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}
