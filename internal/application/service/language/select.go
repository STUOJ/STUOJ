package language

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/language"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/repository/queryfield"
)

func Select(params request.QueryLanguageParams, reqUser request.ReqUser) ([]response.LanguageData, error) {
	var res []response.LanguageData
	languageQuery := params2Model(params)
	languageQuery.Field = *queryfield.LanguageSimpleField
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

func Statistics(params request.LanguageStatisticsParams, reqUser request.ReqUser) (response.StatisticsRes, error) {
	languageQuery := params2Model(params.QueryLanguageParams)
	languageQuery.GroupBy = params.GroupBy
	resp, err := language.Query.GroupCount(languageQuery)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
