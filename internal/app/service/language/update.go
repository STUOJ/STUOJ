package language

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/language"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

func Update(req request.UpdateLanguageReq, reqUser model.ReqUser) error {
	if reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	languageDomain := language.NewLanguage(
		language.WithId(uint64(req.ID)),
		language.WithStatus(entity.LanguageStatus(req.Status)),
		language.WithMapId(uint32(req.MapID)),
		language.WithSerial(uint16(req.Serial)),
	)
	return languageDomain.Update()
}
