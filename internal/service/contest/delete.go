package contest

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID删除比赛
func Delete(id uint64, uid uint64, role entity.Role) error {
	// 查询比赛
	c0, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("比赛不存在")
	}

	if role < entity.RoleAdmin && c0.UserId != uid {
		return errors.New("没有权限，只能删除自己的比赛")
	}

	// 删除比赛
	err = dao.DeleteContestById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除比赛失败")
	}

	return nil
}
