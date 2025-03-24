package model

import "gorm.io/gorm"

type Pagination struct {
	Page     int // 当前页码
	PageSize int // 每页数量
}

func NewPagination(page, pageSize int) Pagination {
	return Pagination{Page: page, PageSize: pageSize}
}

func (p Pagination) Limit() int {
	return p.PageSize
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *Pagination) GeneratePage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.PageSize > 0 {
			return db.Offset(p.Offset()).Limit(p.Limit())
		}
		return db.Offset(0).Limit(1)
	}
}
