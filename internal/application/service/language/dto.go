package language

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/language"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/utils"
)

func params2Model(params request.QueryLanguageParams) (query querycontext.LanguageQueryContext) {
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.OrderBy, *params.Order)
	}
	if params.Status != nil {
		if statusIds, err := utils.StringToUint8Slice(*params.Status); err == nil {
			query.Status.Set(statusIds)
		}
	}
	return query
}

func domain2response(languageDomain language.Language) (res response.LanguageData) {
	res = response.LanguageData{
		Id:     languageDomain.Id.Value(),
		Serial: languageDomain.Serial.Value(),
		Status: uint8(languageDomain.Status.Value()),
		Name:   languageDomain.Name.Value(),
		MapId:  languageDomain.MapId.Value(),
	}
	return
}
