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
}

func (b *BlogWhere) Parse(c *gin.Context) {
	if c.Query("title") != "" {
		b.Title.Set(c.Query("title"))
	}
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			b.Status.Set(entity.BlogStatus(status))
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
		b.ProblemId.Set(problemsInt)
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
		b.UserId.Set(usersInt)
	}
	timePreiod := &Period{}
	err := timePreiod.GetPeriod(c)
	if err != nil {
		log.Println(err)
	} else {
		b.StartTime.Set(timePreiod.StartTime)
		b.EndTime.Set(timePreiod.EndTime)
	}
	if c.Query("page") != "" {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			log.Println(err)
		} else {
			b.Page.Set(uint64(page))
		}
	}
	if c.Query("size") != "" {
		size, err := strconv.Atoi(c.Query("size"))
		if err != nil {
			log.Println(err)
		} else {
			b.Size.Set(uint64(size))
		}
	}
}

func (b *BlogWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		if b.Id.Exist() {
			whereClause["tbl_blog.id"] = b.Id.Value()
		}
		if b.Status.Exist() {
			whereClause["tbl_blog.status"] = b.Status.Value()
		}
		where := db.Where(whereClause)

		if b.UserId.Exist() {
			where.Where("tbl_blog.user_id in ?", b.UserId.Value())
		}
		if b.ProblemId.Exist() {
			where.Where("tbl_blog.problem_id in ?", b.ProblemId.Value())
		}
		if b.Title.Exist() {
			where = where.Where("tbl_blog.title LIKE ?", "%"+b.Title.Value()+"%")
		}
		if b.StartTime.Exist() {
			where = where.Where("tbl_blog.create_time >= ?", b.StartTime.Value())
		}
		if b.EndTime.Exist() {
			where = where.Where("tbl_blog.create_time <= ?", b.EndTime.Value())
		}
		return where
	}
}

func (b *BlogWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := b.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((b.Page.Value() - 1) * b.Size.Value())).Limit(int(b.Size.Value()))
	}
}
