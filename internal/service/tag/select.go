package tag

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

type TagPage struct {
	Tags []entity.Tag `json:"tags"`
	model.Page
}

// 查询所有标签
func Select(condition model.TagWhere) (TagPage, error) {
	if !condition.Page.Exist() {
		condition.Page.Set(1)
	}
	if !condition.Size.Exist() {
		condition.Size.Set(10)
	}
	tags, err := dao.SelectTags(condition)
	if err != nil {
		log.Println(err)
		return TagPage{}, errors.New("查询题单失败")
	}

	count, err := dao.CountTags(condition)
	if err != nil {
		log.Println(err)
		return TagPage{}, errors.New("查询统计失败")
	}
	uPage := TagPage{
		Tags: tags,
		Page: model.Page{
			Total: count,
			Page:  condition.Page.Value(),
			Size:  condition.Size.Value(),
		},
	}

	return uPage, nil
}
