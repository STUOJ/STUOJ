package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
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

	if c.Status != entity.CollectionPublic && role < entity.RoleAdmin {
		userIdsMap := make(map[uint64]struct{})
		for _, uid := range c.UserIds {
			userIdsMap[uid] = struct{}{}
		}
		if _, exists := userIdsMap[userId]; !exists {
			return entity.Collection{}, errors.New("没有该题权限")
		}
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
