package option

import "gorm.io/gorm"

type Pagination struct {
	Page     uint64 // 当前页码
	PageSize uint64 // 每页数量
}

func NewPagination(page, pageSize uint64) Pagination {
	return Pagination{Page: page, PageSize: pageSize}
}

func (p Pagination) Limit() uint64 {
	return p.PageSize
}

func (p Pagination) Offset() uint64 {
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
