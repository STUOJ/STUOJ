package judge

import (
	"STUOJ/internal/domain/language"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

func SelectLanguageMapId(languageId int64) (int64, error) {
	languageQuery := querycontext.LanguageQueryContext{}
	languageQuery.Id.Add(languageId)
	languageQuery.Field.SelectMapId()
	languageDomain, _, err := language.Query.SelectOne(languageQuery)
	if err != nil {
		return 0, err
	}
	return int64(languageDomain.MapId.Value()), nil
}
