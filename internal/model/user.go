package model

import (
	"STUOJ/internal/entity"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserWhere struct {
	Role    Field[entity.Role]
	Page    Field[uint64]
	Size    Field[uint64]
	OrderBy Field[string]
	Order   Field[string]
}

func (con *UserWhere) Parse(c *gin.Context) {
	if c.Query("role") != "" {
		role, err := strconv.Atoi(c.Query("role"))
		if err != nil {
			log.Println(err)
		} else {
			con.Role.Set(entity.Role(role))
		}
	}
	if c.Query("page") != "" {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			log.Println(err)
		} else {
			con.Page.Set(uint64(page))
		}
	}
	if c.Query("size") != "" {
		size, err := strconv.Atoi(c.Query("size"))
		if err != nil {
			log.Println(err)
		} else {
			con.Size.Set(uint64(size))
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

func (con *UserWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		if con.Role.Exist() {
			whereClause["role"] = con.Role.Value()
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

func (con *UserWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
