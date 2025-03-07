package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
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

	if role < entity.RoleAdmin && c0.UserId != uid {
		return errors.New("没有权限，只能删除自己的题单")
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
	_, err := dao.SelectProblemById(pid, model.ProblemWhere{})
	if err != nil {
		log.Println(err)
		return errors.New("题目不存在")
	}

	// 读取题单
	c0, err := dao.SelectCollectionById(cid)
	if err != nil {
		log.Println(err)
		return errors.New("题单不存在")
	}

	flag := false
	if role < entity.RoleAdmin && c0.UserId != uid {
		for _, userId := range c0.UserIds {
			if userId == uid {
				flag = true
			}
		}
	} else {
		flag = true
	}
	if !flag {
		return errors.New("没有权限")
	}

	// 删除题单题目
	err = dao.DeleteCollectionProblem(cp)
	if err != nil {
		log.Println(err)
		return errors.New("删除失败")
	}

	// 更新题单更新时间
	err = dao.UpdateCollectionUpdateTimeById(pid)
	if err != nil {
		log.Println(err)
		return errors.New("更新题单更新时间失败")
	}

	return nil
}

func DeleteUser(cid uint64, delUid uint64, uid uint64, role entity.Role) error {
	cu := entity.CollectionUser{
		CollectionId: cid,
		UserId:       delUid,
	}
	_, err := dao.SelectUserById(delUid)
	if err != nil {
		log.Println(err)
		return errors.New("用户不存在")
	}
	c0, err := dao.SelectCollectionById(cid)
	if err != nil {
		log.Println(err)
		return errors.New("题单不存在")
	}
	if role < entity.RoleAdmin {
		if c0.UserId != uid {
			return errors.New("没有权限，只能操作自己的题单")
		}
	}
	if c0.UserId == delUid {
		return errors.New("题单创建者不能删除")
	}
	for _, i := range c0.UserIds {
		if i == delUid {
			err = dao.DeleteCollectionUser(cu)
			if err != nil {
				log.Println(err)
				return errors.New("删除用户失败")
			}
			return nil
		}
	}
	return errors.New("用户不在题单中")
}
