package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"slices"
	"time"
)

// 根据ID更新题单
func Update(coll entity.Collection, uid uint64, role entity.Role) error {
	// 查询题单
	c0, err := dao.SelectCollectionById(coll.Id)
	if err != nil {
		log.Println(err)
		return errors.New("题单不存在")
	}

	if role < entity.RoleAdmin && c0.UserId != uid {
		return errors.New("没有权限，只能操作自己的题单")

	}

	// 更新题单
	updateTime := time.Now()
	c0.Title = coll.Title
	c0.Description = coll.Description
	c0.Status = coll.Status
	c0.UpdateTime = updateTime

	// 更新题单
	err = dao.UpdateCollectionById(c0)
	if err != nil {
		log.Println(err)
		return errors.New("修改题单失败")
	}

	return nil
}

func UpdateProblem(cp entity.CollectionProblem, uid uint64, role entity.Role) error {
	// 查询题单
	c0, err := SelectById(cp.CollectionId, uid, role)
	if err != nil {
		return err
	}

	if role < entity.RoleAdmin && c0.UserId != uid && !slices.Contains(c0.UserIds, uid) {
		return errors.New("没有权限")
	}

	err = dao.UpdateCollectionProblem(cp)
	if err != nil {
		log.Println(err)
		return errors.New("更新题单失败")
	}
	err = dao.UpdateCollectionUpdateTimeById(cp.CollectionId)
	if err != nil {
		log.Println(err)
		return errors.New("更新题单更新时间失败")
	}
	return nil
}
