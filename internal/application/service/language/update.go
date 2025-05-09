package language

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/language"
	entity "STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/pkg/errors"
)

func Update(req request.UpdateLanguageReq, reqUser request.ReqUser) error {
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
