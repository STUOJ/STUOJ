package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID删除题单
func Delete(id uint64, uid uint64, role entity.Role) error {
	// 查询题单
	c0, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("题单不存在")
	}

	if role < entity.RoleAdmin {
		if c0.UserId != uid {
			return errors.New("没有权限，只能删除自己的题单")
		}
	}

	// 删除题单
	err = dao.DeleteCollectionById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除题单失败")
	}

	return nil
}

// 删除题单的某个题目
func DeleteProblem(cid uint64, pid uint64, uid uint64, role entity.Role) error {
	// 初始化题目标签
	cp := entity.CollectionProblem{
		CollectionId: cid,
		ProblemId:    pid,
	}

	// 读取题目
	p0, err := dao.SelectProblemById(pid)
	if err != nil {
		log.Println(err)
		return errors.New("题目不存在")
	}

	if role < entity.RoleAdmin {
		userIdsMap := make(map[uint64]struct{})
		for _, uid := range p0.UserIds {
			userIdsMap[uid] = struct{}{}
		}
		if _, exists := userIdsMap[uid]; !exists {
			return errors.New("没有该题权限")
		}
	}

	// 读取标签
	_, err = dao.SelectTagById(cid)
	if err != nil {
		log.Println(err)
		return errors.New("标签不存在")
	}

	// 检查题目标签关系是否存在
	count, err := dao.CountProblemTag(cp)
	if err != nil || count < 1 {
		if err != nil {
			log.Println(err)
		}
		return errors.New("该题目不存在该标签")
	}

	// 更新题目更新时间
	err = dao.UpdateProblemUpdateTimeById(pid)
	if err != nil {
		log.Println(err)
		return errors.New("更新题目更新时间失败")
	}

	// 删除题目标签
	err = dao.DeleteProblemTag(cp)
	if err != nil {
		log.Println(err)
		return errors.New("删除失败")
	}

	return nil
}
