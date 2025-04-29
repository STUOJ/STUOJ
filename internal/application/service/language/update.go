package language

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/language"
	entity2 "STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

func Update(req request.UpdateLanguageReq, reqUser model.ReqUser) error {
	if reqUser.Role < entity2.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	languageDomain := language.NewLanguage(
		language.WithId(req.Id),
		language.WithStatus(entity2.LanguageStatus(req.Status)),
		language.WithMapId(uint32(req.MapId)),
		language.WithSerial(uint16(req.Serial)),
	)
	return languageDomain.Update()
}
