package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 插入题单
func Insert(c entity.Collection) (uint64, error) {
	var err error

	updateTime := time.Now()
	c.UpdateTime = updateTime
	c.CreateTime = updateTime

	// 插入题单
	c.Id, err = dao.InsertCollection(c)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入失败")
	}

	return c.Id, nil
}

// 给题单添加题目
func InsertProblem(cid uint64, pid uint64, uid uint64, role entity.Role) error {
	// 初始化题单题目
	cp := entity.CollectionProblem{
		CollectionId: cid,
		ProblemId:    pid,
	}

	// 读取题目
	_, err := dao.SelectProblemById(pid)
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

	if role < entity.RoleAdmin {
		if c0.UserId != uid {
			return errors.New("没有权限，只能操作自己的题单")
		}
	}

	// 更新题单更新时间
	err = dao.UpdateCollectionUpdateTimeById(pid)
	if err != nil {
		log.Println(err)
		return errors.New("更新题单更新时间失败")
	}

	// 插入题单题目
	err = dao.InsertCollectionProblem(cp)
	if err != nil {
		log.Println(err)
		return errors.New("添加题目失败")
	}

	return nil
}
