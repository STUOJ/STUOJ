package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

//go:generate go run ../../../utils/gen/querymodel_gen.go LanguageQueryModel
type LanguageQueryModel struct {
	Id     model.FieldList[int64]
	Name   model.Field[string]
	Serial model.FieldList[int16]
	MapId  model.FieldList[int64]
	Status model.FieldList[int8]
	Page   option.Pagination
	Sort   option.Sort
	Field  field.LanguageField
}

func (query *LanguageQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.LanguageId, option.OpIn, query.Id.Value())
	}
	if query.Name.Exist() {
		options.Filters.Add(field.LanguageName, option.OpLike, query.Name.Value())
	}
	if query.Serial.Exist() {
		options.Filters.Add(field.LanguageSerial, option.OpIn, query.Serial.Value())
	}
	if query.MapId.Exist() {
		options.Filters.Add(field.LanguageMapId, option.OpIn, query.MapId.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.LanguageStatus, option.OpIn, query.Status.Value())
	}
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
