package language

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/language"
	"STUOJ/internal/model/option"
	"STUOJ/utils"
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
		Id:     languageDomain.Id,
		Serial: languageDomain.Serial,
		Status: uint8(languageDomain.Status),
		Name:   languageDomain.Name.String(),
		MapId:  languageDomain.MapId,
	}
	return
}
