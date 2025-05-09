package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go LanguageQueryContext
type LanguageQueryContext struct {
	Id     dto.FieldList[int64]
	Name   dto.Field[string]
	Serial dto.FieldList[uint16]
	MapId  dto.FieldList[uint32]
	Status dto.FieldList[entity.LanguageStatus]
	option2.QueryParams
	Field field.LanguageField
}

// applyFilter 应用编程语言查询过滤条件
// 根据查询参数设置过滤条件，并返回更新后的options对象
func (query *LanguageQueryContext) applyFilter(options option2.Options) option2.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.LanguageId, option2.OpIn, query.Id.Value())
	}
	if query.Name.Exist() {
		filters.Add(field.LanguageName, option2.OpLike, query.Name.Value())
	}
	if query.Serial.Exist() {
		filters.Add(field.LanguageSerial, option2.OpIn, query.Serial.Value())
	}
	if query.MapId.Exist() {
		filters.Add(field.LanguageMapId, option2.OpIn, query.MapId.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.LanguageStatus, option2.OpIn, query.Status.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
