package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserWhere struct {
	Role    FieldList[uint64]
	Name    Field[string]
	Page    Field[uint64]
	Size    Field[uint64]
	OrderBy Field[string]
	Order   Field[string]
}

func (con *UserWhere) Parse(c *gin.Context) {
	con.Role.Parse(c, "role")
	con.Name.Parse(c, "name")
	con.Page.Parse(c, "page")
	con.Size.Parse(c, "size")
	con.OrderBy.Parse(c, "order_by")
	con.Order.Parse(c, "order")
}

func (con *UserWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		where := db.Where(whereClause)
		if con.Role.Exist() {
			where.Where("tbl_user.role in ?", con.Role.Value())
		}
		if con.Name.Exist() {
			where = where.Where("tbl_user.name LIKE ?", "%"+con.Name.Value()+"%")
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

func (con *UserWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
