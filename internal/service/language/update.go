package language

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
)

func Update(lang entity.Language) error {
	err := dao.UpdateLanguage(lang)
	if err != nil {
		return err
	}
	return nil
}
