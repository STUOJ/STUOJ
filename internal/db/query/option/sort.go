package option

import "gorm.io/gorm"

type SortOrder string

const (
	OrderAsc  SortOrder = "ASC"
	OrderDesc SortOrder = "DESC"
)

type Sort interface {
	GenerateSort() func(*gorm.DB) *gorm.DB
}

type SortQuery struct {
	OrderBy string
	Order   SortOrder
}

func NewSortQuery(orderBy string, order SortOrder) Sort {
	return &SortQuery{OrderBy: orderBy, Order: order}
}

func (con *SortQuery) GenerateSort() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if con.OrderBy != "" && con.Order != "" {
			db = db.Order(con.OrderBy + " " + string(con.Order))
		}
		return db
	}
}
