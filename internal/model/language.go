package model

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LanguageWhere struct {
	Status Field[uint64]
}

func (con *LanguageWhere) Parse(c *gin.Context) {
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			con.Status.Set(uint64(status))
		}
	}
}
func (con *LanguageWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		if con.Status.Exist() {
			whereClause["status"] = con.Status.Value()
		}
		return db.Where(whereClause)
	}
}
