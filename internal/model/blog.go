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
	Status    FieldList[entity.BlogStatus]
	StartTime Field[time.Time]
	EndTime   Field[time.Time]
	Page      Field[uint64]
	Size      Field[uint64]
	OrderBy   Field[string]
	Order     Field[string]
}

func (con *BlogWhere) Parse(c *gin.Context) {
	if titleQuery := c.Query("title"); titleQuery != "" {
		con.Title.Set(titleQuery)
	}
	if statusQuery := c.Query("status"); statusQuery != "" {
		statuses := strings.Split(statusQuery, ",")
		var statusesInt []entity.BlogStatus
		for _, status := range statuses {
			statusInt, err := strconv.Atoi(status)
			if err != nil {
				log.Println(err)
			} else {
				statusesInt = append(statusesInt, entity.BlogStatus(statusInt))
			}
		}
		con.Status.Set(statusesInt)
	}
	if problemQuery := c.Query("problem"); problemQuery != "" {
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
	if userQuery := c.Query("user"); userQuery != "" {
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
	if err == nil {
		con.StartTime.Set(timePreiod.StartTime)
		con.EndTime.Set(timePreiod.EndTime)
	}
	if pageQuery := c.Query("page"); pageQuery != "" {
		page, err := strconv.Atoi(pageQuery)
		if err != nil {
			log.Println(err)
		} else {
			con.Page.Set(uint64(page))
		}
	}
	if sizeQuery := c.Query("size"); sizeQuery != "" {
		size, err := strconv.Atoi(sizeQuery)
		if err != nil {
			log.Println(err)
		} else {
			con.Size.Set(uint64(size))
		}
	}
	if order := c.Query("order"); order != "" {
		if orderBy := c.Query("order_by"); orderBy != "" {
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
		return where
	}
}

func (con *BlogWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
