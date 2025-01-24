package model

import "gorm.io/gorm"

type HistoryWhere struct {
}

func (con *HistoryWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		where := db.Where(whereClause)
		query := []string{"tbl_history.*"}
		query = append(query, briefUserSelect()...)
		where = briefUserJoins(where, "tbl_history")
		return where.Select(query)
	}
}
