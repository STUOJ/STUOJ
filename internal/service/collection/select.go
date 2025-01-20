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
func SelectById(id uint64) (entity.Collection, error) {
	t, err := dao.SelectCollectionById(id)
	if err != nil {
		return entity.Collection{}, err
	}

	return t, nil
}

// 查询所有题单
func Select(condition model.CollectionWhere) (CollectionPage, error) {
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
