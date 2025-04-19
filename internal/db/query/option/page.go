package option

import "gorm.io/gorm"

type Pagination struct {
	Page     int64 // 当前页码
	PageSize int64 // 每页数量
}

func NewPagination(page, pageSize int64) Pagination {
	return Pagination{Page: page, PageSize: pageSize}
}

func (p Pagination) Limit() int64 {
	return p.PageSize
}

func (p Pagination) Offset() int64 {
	return (p.Page - 1) * p.PageSize
}

func (p *Pagination) GeneratePage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.PageSize > 0 {
			return db.Offset(int(p.Offset())).Limit(int(p.Limit()))
		}
		return db.Offset(0).Limit(1)
	}
}
