package model

import "gorm.io/gorm"

type QueryOptions struct {
	Filters Filters    // 过滤条件列表
	Sort    Sort       // 排序条件
	Page    Pagination // 分页条件
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
		return db
	}
}
