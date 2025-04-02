package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

type LanguageQueryModel struct {
	Id     model.FieldList[uint64]
	Name   model.Field[string]
	Serial model.FieldList[uint16]
	MapId  model.FieldList[uint64]
	Status model.FieldList[uint8]
	Page   model.QueryPage
	Sort   model.QuerySort
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
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
