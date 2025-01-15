package model

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LanguageWhere struct {
	Status  Field[uint64]
	OrderBy Field[string]
	Order   Field[string]
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
	if c.Query("order") != "" {
		order := c.Query("order")
		if c.Query("order_by") != "" {
			orderBy := c.Query("order_by")
			con.OrderBy.Set(orderBy)
			con.Order.Set(order)
		}
	}

}
func (con *LanguageWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		if con.Status.Exist() {
			whereClause["status"] = con.Status.Value()
		}
		where := db.Where(whereClause)
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
