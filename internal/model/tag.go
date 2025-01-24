package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TagWhere struct {
	Name Field[string]
	Page Field[uint64]
	Size Field[uint64]
}

func (con *TagWhere) Parse(c *gin.Context) {
	con.Name.Parse(c, "name")
	con.Page.Parse(c, "page")
	con.Size.Parse(c, "size")
}

func (con *TagWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		where := db.Where(whereClause)
		if con.Name.Exist() {
			where = where.Where("tbl_tag.title LIKE ?", "%"+con.Name.Value()+"%")
		}
		return where
	}
}

func (con *TagWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		if con.Page.Exist() && con.Size.Exist() {
			return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
		}
		return where(db).Offset(0).Limit(1)
	}
}
