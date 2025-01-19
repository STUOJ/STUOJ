package model

import (
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
	Status     FieldList[uint64]
	Page       Field[uint64]
	Size       Field[uint64]
	OrderBy    Field[string]
	Order      Field[string]
}

func (con *SubmissionWhere) Parse(c *gin.Context) {
	con.ProblemId.Parse(c, "problem")
	con.UserId.Parse(c, "user")
	con.LanguageId.Parse(c, "language")
	timePreiod := &Period{}
	err := timePreiod.GetPeriod(c)
	if err == nil {
		con.StartTime.Set(timePreiod.StartTime)
		con.EndTime.Set(timePreiod.EndTime)
	}
	con.Status.Parse(c, "status")
	con.Page.Parse(c, "page")
	con.Size.Parse(c, "size")
	con.OrderBy.Parse(c, "order_by")
	con.Order.Parse(c, "order")
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

func (con *SubmissionWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
