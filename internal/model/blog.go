package model

import (
	"STUOJ/internal/entity"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogWhere struct {
	Id        Field[uint64]
	UserId    FieldList[uint64]
	ProblemId FieldList[uint64]
	Title     Field[string]
	Status    FieldList[uint64]
	StartTime Field[time.Time]
	EndTime   Field[time.Time]
	Page      Field[uint64]
	Size      Field[uint64]
	OrderBy   Field[string]
	Order     Field[string]
}

func (con *BlogWhere) Parse(c *gin.Context) {
	con.Title.Parse(c, "title")
	con.Status.Parse(c, "status")
	con.ProblemId.Parse(c, "problem")
	con.UserId.Parse(c, "user")
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

func (con *BlogWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		if con.Id.Exist() {
			whereClause["tbl_blog.id"] = con.Id.Value()
		}
		where := db.Where(whereClause)

		if con.Status.Exist() {
			where.Where("tbl_blog.status in ?", con.Status.Value())
		}
		if con.UserId.Exist() {
			where.Where("tbl_blog.user_id in ?", con.UserId.Value())
		}
		if con.ProblemId.Exist() {
			where.Where("tbl_blog.problem_id in ?", con.ProblemId.Value())
		}
		if con.Title.Exist() {
			where = where.Where("tbl_blog.title LIKE ?", "%"+con.Title.Value()+"%")
		}
		if con.StartTime.Exist() {
			where = where.Where("tbl_blog.create_time >= ?", con.StartTime.Value())
		}
		if con.EndTime.Exist() {
			where = where.Where("tbl_blog.create_time <= ?", con.EndTime.Value())
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
		query := []string{"tbl_blog.*"}
		query = append(query, briefUserSelect()...)
		query = append(query, briefProblemSelect()...)
		where = briefProblemJoins(where, "tbl_blog")
		where = briefUserJoins(where, "tbl_blog")
		return where.Select(query)
	}
}

func (con *BlogWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		if con.Page.Exist() && con.Size.Exist() {
			return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
		}
		return where(db).Offset(0).Limit(1)
	}
}

type BriefBlog struct {
	BlogTitle  string            `gorm:"column:blog_title"`
	BlogStatus entity.BlogStatus `gorm:"column:blog_status"`
}

func briefBlogSelect() []string {
	return []string{
		"tbl_blog.title as blog_title",
		"tbl_blog.status as blog_status",
	}
}

func briefBlogJoins(db *gorm.DB, tbl string) *gorm.DB {
	return db.Joins(fmt.Sprintf("LEFT JOIN tbl_blog ON %s.blog_id = tbl_blog.id", tbl))
}

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
