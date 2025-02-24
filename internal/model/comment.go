package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentWhere struct {
	UserId    Field[uint64]
	BlogId    Field[uint64]
	Status    FieldList[uint64]
	StartTime Field[time.Time]
	EndTime   Field[time.Time]
	Page      Field[uint64]
	Size      Field[uint64]
	OrderBy   Field[string]
	Order     Field[string]
}

func (con *CommentWhere) Parse(c *gin.Context) {
	con.UserId.Parse(c, "user")
	con.BlogId.Parse(c, "blog")
	con.Status.Parse(c, "status")
	timePreiod := &Period{}
	err := timePreiod.GetPeriod(c)
	if err == nil {
		con.StartTime.Set(timePreiod.StartTime)
		con.EndTime.Set(timePreiod.EndTime)
	}
	con.Page.Parse(c, "page")
	con.Size.Parse(c, "size")
	con.OrderBy.Parse(c, "order_by")
	con.Order.Parse(c, "order")
}

func (con *CommentWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		if con.UserId.Exist() {
			whereClause["tbl_comment.user_id"] = con.UserId.Value()
		}
		if con.BlogId.Exist() {
			whereClause["tbl_comment.blog_id"] = con.BlogId.Value()
		}
		where := db.Where(whereClause)
		if con.Status.Exist() {
			where.Where("tbl_comment.status in ?", con.Status.Value())
		}
		if con.StartTime.Exist() {
			where.Where("tbl_comment.create_time >= ?", con.StartTime.Value())
		}
		if con.EndTime.Exist() {
			where.Where("tbl_comment.create_time <= ?", con.EndTime.Value())
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
		query := []string{"tbl_comment.*"}
		query = append(query, briefUserSelect()...)
		query = append(query, briefBlogSelect()...)
		where = briefUserJoins(where, "tbl_comment")
		where = briefBlogJoins(where, "tbl_comment")
		return where.Select(query)
	}
}

func (con *CommentWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		if con.Page.Exist() && con.Size.Exist() {
			return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
		}
		return where(db).Offset(0).Limit(1)
	}
}
