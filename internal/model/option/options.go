package option

import (
	"fmt"

	"gorm.io/gorm"
)

type QueryOptions struct {
	Filters Filters    // 过滤条件列表
	Sort    Sort       // 排序条件
	Page    Pagination // 分页条件
	Field   FieldSelector
	Errors  []error
}

func NewQueryOptions() *QueryOptions {
	return &QueryOptions{}
}

func (q *QueryOptions) GenerateQuery() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if q.Filters.Conditions != nil {
			db = q.Filters.GenerateWhere()(db)
			q.Errors = append(q.Errors, q.Filters.Errors...)
		}
		if q.Sort != nil {
			db = q.Sort.GenerateSort()(db)
		}
		if q.Page.PageSize > 0 {
			db = q.Page.GeneratePage()(db)
		}
		if q.Field != nil {
			db = db.Select(q.Field.SelectedColumns())
		}
		return db
	}
}

func (q *QueryOptions) GetFilters() *Filters {
	return &q.Filters
}

type GroupCountOptions struct {
	GroupBy GroupField
	Filters Filters
}

func NewGroupCountOptions(groupBy GroupField) *GroupCountOptions {
	return &GroupCountOptions{
		GroupBy: groupBy,
	}
}

func (g *GroupCountOptions) GenerateQuery() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if ok := g.GroupBy.Verify(); !ok {
			return db
		}
		if g.Filters.Conditions != nil {
			db = g.Filters.GenerateWhere()(db)
		}
		db.Select(fmt.Sprintf("%s AS field", g.GroupBy.Field()), "COUNT(*) AS count")
		return db.Group(g.GroupBy.Field())
	}
}

func (g *GroupCountOptions) SetGroupBy(field string) {
	g.GroupBy.SetField(field)
}

func (g *GroupCountOptions) Verify() bool {
	return g.GroupBy.Verify()
}

func (g *GroupCountOptions) GetFilters() *Filters {
	return &g.Filters
}

type GroupCountResult struct {
	Field string `gorm:"column:field"`
	Count int64  `gorm:"column:count"`
}

type Options interface {
	GetFilters() *Filters
}
