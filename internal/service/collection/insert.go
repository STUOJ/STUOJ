package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
	"time"

	iproblem "STUOJ/internal/service/problem"
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
	_, err := iproblem.SelectById(pid, uid, role, model.ProblemWhere{})
	if err != nil {
		return err
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

func InsertUser(cid uint64, addUid uint64, uid uint64, role entity.Role) error {
	cu := entity.CollectionUser{
		CollectionId: cid,
		UserId:       addUid,
	}

	_, err := dao.SelectUserById(addUid)
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
	if c0.UserId == addUid {
		return errors.New("不能添加自己为题单成员")
	}
	for _, i := range c0.UserIds {
		if i == addUid {
			return errors.New("用户已存在题单中")
		}
	}
	err = dao.InsertCollectionUser(cu)
	if err != nil {
		log.Println(err)
		return errors.New("添加用户失败")
	}
	return nil
}
