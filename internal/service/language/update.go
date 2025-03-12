package language

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
)

func Update(lang entity.Language, role entity.Role) error {
	if role < entity.RoleAdmin {
		return errors.New("无权限")
	}

	lang_, err := SelectById(lang.Id)
	if err != nil {
		return err
	}

	if lang_.Name != lang.Name && lang.Name != "" {
		return errors.New("语言名不可修改")
	}
	if role < entity.RoleRoot && lang_.MapId != lang.MapId && lang.MapId != 0 {
		return errors.New("无权限更改映射ID")
	}

	err = dao.UpdateLanguage(lang)
	if err != nil {
		return err
	}
	return nil
}
