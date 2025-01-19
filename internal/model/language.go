package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LanguageWhere struct {
	Status  FieldList[uint64]
	OrderBy Field[string]
	Order   Field[string]
}

func (con *LanguageWhere) Parse(c *gin.Context) {
	con.Status.Parse(c, "status")
	con.OrderBy.Parse(c, "order_by")
	con.Order.Parse(c, "order")
}
func (con *LanguageWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		where := db.Where(whereClause)
		if con.Status.Exist() {
			where.Where("tbl_language.status in ?", con.Status.Value())
		}
		if con.OrderBy.Exist() {
			orderBy := con.OrderBy.Value()
			order := con.Order.Value()
			if order == "desc" {
				order = "DESC"
			} else {
				order = "ASC"
			}
			where = where.Order(orderBy + " " + order)
		}
		return where
	}
}
