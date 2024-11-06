package db

import (
	"STUOJ/model"
	"time"
)

// 插入题解
func InsertSolution(s model.Solution) (uint64, error) {
	updateTime := time.Now()
	s.UpdateTime = updateTime
	s.CreateTime = updateTime
	tx := Db.Create(&s)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return s.Id, nil
}

// 根据ID查询题解
func SelectSolutionById(id uint64) (model.Solution, error) {
	var s model.Solution

	tx := Db.Where("id = ?", id).First(&s)
	if tx.Error != nil {
		return model.Solution{}, tx.Error
	}

	return s, nil
}

// 查询所有题解
func SelectAllSolutions() ([]model.Solution, error) {
	var solutions []model.Solution

	tx := Db.Find(&solutions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return solutions, nil
}

// 根据题目ID查询题解
func SelectSolutionsByProblemId(pid uint64) ([]model.Solution, error) {
	var solutions []model.Solution

	tx := Db.Where("problem_id = ?", pid).Find(&solutions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return solutions, nil
}

// 根据ID更新题解
func UpdateSolutionById(s model.Solution) error {
	updateTime := time.Now()
	tx := Db.Model(&model.Solution{}).Where("id = ?", s.Id).Updates(map[string]interface{}{
		"problem_id":  s.ProblemId,
		"language_id": s.LanguageId,
		"source_code": s.SourceCode,
		"update_time": updateTime,
	})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除题解
func DeleteSolutionById(id uint64) error {
	tx := Db.Where("id = ?", id).Delete(&model.Solution{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}