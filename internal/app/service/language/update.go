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
		language.WithId(req.Id),
		language.WithStatus(entity.LanguageStatus(req.Status)),
		language.WithMapId(uint32(req.MapId)),
		language.WithSerial(uint16(req.Serial)),
	)
	return languageDomain.Update()
}
