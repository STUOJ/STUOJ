package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SubmissionWhere struct {
	ProblemId      FieldList[uint64]
	UserId         FieldList[uint64]
	LanguageId     Field[uint64]
	StartTime      Field[time.Time]
	EndTime        Field[time.Time]
	Status         FieldList[uint64]
	ExcludeHistory Field[bool]
	Distinct       Field[string]
	Page           Field[uint64]
	Size           Field[uint64]
	OrderBy        Field[string]
	Order          Field[string]
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
	con.Distinct.Parse(c, "distinct")
	con.Status.Parse(c, "status")
	con.ExcludeHistory.Parse(c, "exclude_history")
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

		query := []string{"tbl_submission.*"}
		query = append(query, briefUserSelect()...)
		query = append(query, briefProblemSelect()...)
		where = briefProblemJoins(where, "tbl_submission")
		where = briefUserJoins(where, "tbl_submission")

		if con.ExcludeHistory.Exist() && con.ExcludeHistory.Value() {
			where = where.Joins("LEFT JOIN tbl_history ON tbl_submission.problem_id = tbl_history.problem_id AND tbl_history.user_id = tbl_submission.user_id")
			where = where.Where("tbl_history.user_id IS NULL")
		}

		return where.Select(query)
	}
}

func (con *SubmissionWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		if con.Page.Exist() && con.Size.Exist() {
			return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
		}
		return where(db).Offset(0).Limit(1)
	}
}
