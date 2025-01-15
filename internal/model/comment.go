package model

import (
	"STUOJ/internal/entity"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentWhere struct {
	UserId    Field[uint64]
	BlogId    Field[uint64]
	Status    Field[entity.CommentStatus]
	StartTime Field[time.Time]
	EndTime   Field[time.Time]
	Page      Field[uint64]
	Size      Field[uint64]
	OrderBy   Field[string]
	Order     Field[string]
}

func (con *CommentWhere) Parse(c *gin.Context) {

	if c.Query("user") != "" {
		user, err := strconv.Atoi(c.Query("user"))
		if err != nil {
			log.Println(err)
		} else {
			con.UserId.Set(uint64(user))
		}
	}
	if c.Query("blog") != "" {
		blog, err := strconv.Atoi(c.Query("blog"))
		if err != nil {
			log.Println(err)
		} else {
			con.BlogId.Set(uint64(blog))
		}
	}
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			con.Status.Set(entity.CommentStatus(status))
		}
	}
	timePreiod := &Period{}
	err := timePreiod.GetPeriod(c)
	if err != nil {
		log.Println(err)
	} else {
		con.StartTime.Set(timePreiod.StartTime)
		con.EndTime.Set(timePreiod.EndTime)
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

func (con *CommentWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		if con.UserId.Exist() {
			whereClause["tbl_comment.user_id"] = con.UserId.Value()
		}
		if con.BlogId.Exist() {
			whereClause["tbl_comment.blog_id"] = con.BlogId.Value()
		}
		if con.Status.Exist() {
			whereClause["tbl_comment.status"] = con.Status.Value()
		}
		where := db.Where(whereClause)
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
		return where
	}
}

func (con *CommentWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
