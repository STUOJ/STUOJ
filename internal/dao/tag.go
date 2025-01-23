package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
)

// 插入标签
func InsertTag(t entity.Tag) (uint64, error) {
	tx := db.Db.Create(&t)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return t.Id, nil
}

// 根据ID查询标签
func SelectTagById(id uint64) (entity.Tag, error) {
	var t entity.Tag
	tx := db.Db.Where("id = ?", id).First(&t)
	if tx.Error != nil {
		return entity.Tag{}, tx.Error
	}

	return t, nil
}

// 查询所有标签
func SelectAllTags() ([]entity.Tag, error) {
	var tags []entity.Tag
	tx := db.Db.Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

// 查询标签
func SelectTags(condition model.TagWhere) ([]entity.Tag, error) {
	var tags []entity.Tag
	where := condition.GenerateWhere()

	tx := db.Db.Model(&entity.Tag{})
	tx = where(tx)
	tx = tx.Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

// 根据ID更新标签
func UpdateTagById(t entity.Tag) error {
	tx := db.Db.Model(&t).Where("id = ?", t.Id).Updates(t)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除标签
func DeleteTagById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Tag{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计标签数量
func CountTags(condition model.TagWhere) (uint64, error) {
	var count int64
	where := condition.GenerateWhereWithNoPage()
	tx := db.Db.Model(&entity.Tag{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

// 给题目添加标签
func InsertProblemTag(pt entity.ProblemTag) error {
	tx := db.Db.Create(&pt)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func SelectTagsByProblemId(pid uint64) ([]entity.Tag, error) {
	var tags []entity.Tag

	tx := db.Db.Table("tbl_tag").Select("tbl_tag.id, tbl_tag.name").Joins("JOIN tbl_problem_tag ON tbl_tag.id = tbl_problem_tag.tag_id").Where("tbl_problem_tag.problem_id = ?", pid).Scan(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

// 查询题目标签关系是否存在
func CountProblemTag(pt entity.ProblemTag) (int64, error) {
	var count int64

	tx := db.Db.Model(&entity.ProblemTag{}).Where("problem_id = ? AND tag_id = ?", pt.ProblemId, pt.TagId).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}

// 删除题目的某个标签
func DeleteProblemTag(pt entity.ProblemTag) error {
	tx := db.Db.Where("problem_id = ? AND tag_id = ?", pt.ProblemId, pt.TagId).Delete(&entity.ProblemTag{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 删除题目的所有标签
func DeleteProblemTagsByProblemId(pid uint64) error {
	tx := db.Db.Where("problem_id = ?", pid).Delete(&entity.ProblemTag{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 按评测状态统计提交信息数量
func CountProblemsGroupByTag() ([]model.CountByTag, error) {
	var counts []model.CountByTag

	tx := db.Db.Model(&entity.ProblemTag{}).Select("tag_id, count(*) as count").Group("tag_id").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
}
