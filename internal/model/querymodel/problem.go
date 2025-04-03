package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"
)

type ProblemQueryModel struct {
	Id        model.FieldList[int64]
	Title     model.Field[string]
	Source    model.Field[string]
	Status    model.FieldList[int8]
	Tag       model.FieldList[int64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      model.QueryPage
	Sort      model.QuerySort
	Field     field.ProblemField
}

func (query *ProblemQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.ProblemId, option.OpIn, query.Id.Value())
	}
	if query.Title.Exist() {
		options.Filters.Add(field.ProblemTitle, option.OpLike, query.Title.Value())
	}
	if query.Source.Exist() {
		options.Filters.Add(field.ProblemSource, option.OpLike, query.Source.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.ProblemStatus, option.OpIn, query.Status.Value())
	}
	if query.Tag.Exist() {
		options.Filters.Add(field.ProblemTag, option.OpHave, query.Tag.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.ProblemCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.ProblemCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	options.Field = &query.Field
	return options
}
