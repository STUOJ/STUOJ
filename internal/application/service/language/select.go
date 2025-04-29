package language

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/language"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/query"
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
