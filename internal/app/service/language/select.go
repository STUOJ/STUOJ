package language

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/query"
	"STUOJ/internal/domain/language"
	"STUOJ/internal/model"
)

func Select(params request.QueryLanguageParams, reqUser model.ReqUser) ([]response.LanguageData, error) {
	var res []response.LanguageData
	languageQuery := params2Model(params)
	languageQuery.Field = *query.LanguageSimpleField
	if reqUser.Role >= entity.RoleAdmin {
		languageQuery.Field.SelectMapId()
	}
	languageDomain, _, err := language.Query.Select(languageQuery)
	if err != nil {
		return nil, err
	}
	for _, l := range languageDomain {
		res = append(res, domain2response(l))
	}
	return res, nil
}
