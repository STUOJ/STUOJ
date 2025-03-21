package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"errors"
	"fmt"
	"log"
	"slices"
)

type CollectionPage struct {
	Collections []entity.Collection `json:"collections"`
	model.Page
}

// 根据ID查询题单
func SelectById(id uint64, userId uint64, role entity.Role) (entity.Collection, error) {
	// 获取题目信息
	c, err := dao.SelectCollectionById(id)

	if err != nil {
		return entity.Collection{}, errors.New("获取题单失败")
	}
	flag := false
	if c.Status != entity.CollectionPublic && role < entity.RoleAdmin && c.UserId != userId {
		if slices.Contains(c.UserIds, userId) {
			flag = true
		}
	} else {
		flag = true
	}
	if !flag {
		return entity.Collection{}, errors.New("没有权限查看该题单")
	}

	pCondition := model.ProblemWhere{}
	pCondition.Id.Set(c.ProblemIds)
	pCondition.Page.Set(uint64(1))
	pCondition.Size.Set(uint64(100))
	pCondition.ScoreUserId.Set(userId)
	pCondition.OrderBy.Set(fmt.Sprintf("FIELD(tbl_problem.id,%s)", utils.Uint64SliceToString(c.ProblemIds)))
	c.Problem, err = dao.SelectProblems(pCondition)
	if err != nil {
		return c, errors.New("获取题目信息失败")
	}

	return c, nil
}

// 查询题单
func Select(condition model.CollectionWhere, uid uint64, role entity.Role) (CollectionPage, error) {
	if !condition.Status.Exist() {
		condition.Status.Set([]uint64{uint64(entity.CollectionPublic)})
	} else {
		for _, v := range condition.Status.Value() {
			if entity.CollectionStatus(v) < entity.CollectionPublic {
				if role < entity.RoleAdmin {
					condition.UserId.Set([]uint64{uid})
				}
			}
		}
	}
	if !condition.Page.Exist() {
		condition.Page.Set(1)
	}
	if !condition.Size.Exist() {
		condition.Size.Set(10)
	}
	users, err := dao.SelectCollections(condition)
	if err != nil {
		log.Println(err)
		return CollectionPage{}, errors.New("查询题单失败")
	}

	count, err := dao.CountCollections(condition)
	if err != nil {
		log.Println(err)
		return CollectionPage{}, errors.New("查询统计失败")
	}
	uPage := CollectionPage{
		Collections: users,
		Page: model.Page{
			Total: count,
			Page:  condition.Page.Value(),
			Size:  condition.Size.Value(),
		},
	}

	return uPage, nil
}
