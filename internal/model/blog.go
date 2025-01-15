package model

import (
	"STUOJ/internal/entity"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogWhere struct {
	Id        Field[uint64]
	UserId    FieldList[uint64]
	ProblemId FieldList[uint64]
	Title     Field[string]
	Status    Field[entity.BlogStatus]
	StartTime Field[time.Time]
	EndTime   Field[time.Time]
	Page      Field[uint64]
	Size      Field[uint64]
	OrderBy   Field[string]
	Order     Field[string]
}

func (con *BlogWhere) Parse(c *gin.Context) {
	if c.Query("title") != "" {
		con.Title.Set(c.Query("title"))
	}
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			con.Status.Set(entity.BlogStatus(status))
		}
	}
	if c.Query("problem") != "" {
		problemQuery := c.Query("problem")
		problems := strings.Split(problemQuery, ",")
		var problemsInt []uint64
		for _, problem := range problems {
			problemInt, err := strconv.Atoi(problem)
			if err != nil {
				log.Println(err)
			} else {
				problemsInt = append(problemsInt, uint64(problemInt))
			}
		}
		con.ProblemId.Set(problemsInt)
	}
	if c.Query("user") != "" {
		userQuery := c.Query("user")
		users := strings.Split(userQuery, ",")
		var usersInt []uint64
		for _, user := range users {
			userInt, err := strconv.Atoi(user)
			if err != nil {
				log.Println(err)
			} else {
				usersInt = append(usersInt, uint64(userInt))
			}
		}
		con.UserId.Set(usersInt)
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

func (con *BlogWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		if con.Id.Exist() {
			whereClause["tbl_blog.id"] = con.Id.Value()
		}
		if con.Status.Exist() {
			whereClause["tbl_blog.status"] = con.Status.Value()
		}
		where := db.Where(whereClause)

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
		return where
	}
}

func (con *BlogWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
