package model

import "gorm.io/gorm"

type SolutionWhere struct {
}

func (con *SolutionWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		where := db.Where(whereClause)
		query := []string{"tbl_solution.*"}
		query = append(query, briefProblemSelect()...)
		where = briefProblemJoins(where, "tbl_solution")
		return where.Select(query)
	}
}
