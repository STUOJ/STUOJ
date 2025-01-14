package model

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SubmissionWhere struct {
	ProblemId  FieldList[uint64]
	UserId     FieldList[uint64]
	LanguageId Field[uint64]
	StartTime  Field[time.Time]
	EndTime    Field[time.Time]
	Status     Field[uint64]
	Page       Field[uint64]
	Size       Field[uint64]
}

func (con *SubmissionWhere) Parse(c *gin.Context) {
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
	if c.Query("language") != "" {
		language, err := strconv.Atoi(c.Query("language"))
		if err != nil {
			log.Println(err)
		} else {
			con.LanguageId.Set(uint64(language))
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
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			con.Status.Set(uint64(status))
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
}

func (con *SubmissionWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		if con.LanguageId.Exist() {
			whereClause["tbl_submission.language_id"] = con.LanguageId.Value()
		}
		if con.Status.Exist() {
			whereClause["tbl_submission.status"] = con.Status.Value()
		}
		where := db.Where(whereClause)
		if con.UserId.Exist() {
			where.Where("tbl_submission.user_id in ?", con.UserId.Value())
		}
		if con.ProblemId.Exist() {
			where.Where("tbl_submission.problem_id in ?", con.ProblemId.Value())
		}
		if con.StartTime.Exist() {
			where.Where("tbl_submission.create_time >= ?", con.StartTime.Value())
		}
		if con.EndTime.Exist() {
			where.Where("tbl_submission.create_time <= ?", con.EndTime.Value())
		}
		return where
	}
}

func (con *SubmissionWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
