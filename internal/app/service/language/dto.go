package language

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/language"
	"STUOJ/utils"
)

func params2Model(params request.QueryLanguageParams) (query querycontext.LanguageQueryContext) {
	if params.OrderBy != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.OrderBy, *params.Order)
	}
	if params.Status != nil {
		if statusIds, err := utils.StringToint8Slice(*params.Status); err == nil {
			query.Status.Set(statusIds)
		}
	}
	return query
}

func domain2response(languageDomain language.Language) (res response.LanguageData) {
	res = response.LanguageData{
		ID:     int64(languageDomain.Id),
		Serial: int64(languageDomain.Serial),
		Status: int64(languageDomain.Status),
		Name:   languageDomain.Name.String(),
		MapID:  int64(languageDomain.MapId),
	}
	return
}
