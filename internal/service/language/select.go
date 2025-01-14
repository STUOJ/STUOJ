package language

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

// 查询所有语言
func Select(condition model.LanguageWhere, role entity.Role) ([]entity.Language, error) {
	var languages []entity.Language

	languages, err := dao.SelectLanguage(condition)
	if err != nil {
		log.Println(err)
		return nil, errors.New("查询语言失败")
	}
	if role < entity.RoleRoot {
		for i := range languages {
			languages[i].MapId = 0
		}
	}

	return languages, nil
}

func SelectById(id uint64) (entity.Language, error) {
	var l entity.Language

	l, err := dao.SelectLanguageById(id)
	if err != nil {
		log.Println(err)
		return entity.Language{}, errors.New("查询语言失败")
	}

	return l, nil
}

// 根据名字模糊查询语言
func SelectLikeName(name string) (entity.Language, error) {
	var l entity.Language

	l, err := dao.SelectLanguageLikeName(name)
	if err != nil {
		log.Println(err)
		return entity.Language{}, errors.New("查询语言失败")
	}

	return l, nil
}
