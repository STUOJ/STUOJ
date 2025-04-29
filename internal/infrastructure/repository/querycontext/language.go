package querycontext

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go LanguageQueryContext
type LanguageQueryContext struct {
	Id     model.FieldList[int64]
	Name   model.Field[string]
	Serial model.FieldList[uint16]
	MapId  model.FieldList[uint32]
	Status model.FieldList[uint8]
	option.QueryParams
	Field field.LanguageField
}

// applyFilter 应用编程语言查询过滤条件
// 根据查询参数设置过滤条件，并返回更新后的options对象
func (query *LanguageQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.LanguageId, option.OpIn, query.Id.Value())
	}
	if query.Name.Exist() {
		filters.Add(field.LanguageName, option.OpLike, query.Name.Value())
	}
	if query.Serial.Exist() {
		filters.Add(field.LanguageSerial, option.OpIn, query.Serial.Value())
	}
	if query.MapId.Exist() {
		filters.Add(field.LanguageMapId, option.OpIn, query.MapId.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.LanguageStatus, option.OpIn, query.Status.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
