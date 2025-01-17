package model

import (
	"STUOJ/internal/entity"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 题目数据（题目+标签+评测点数据+题解）
type ProblemData struct {
	Problem   entity.Problem    `json:"problem,omitempty"`
	Testcases []entity.Testcase `json:"testcases,omitempty"`
	Solutions []entity.Solution `json:"solutions,omitempty"`
}

type ProblemWhere struct {
	Id         Field[uint64]
	Title      Field[string]
	Difficulty FieldList[uint64]
	Status     FieldList[uint64]
	Tag        FieldList[uint64]
	UserId     Field[uint64]
	Page       Field[uint64]
	Size       Field[uint64]
	OrderBy    Field[string]
	Order      Field[string]
	StartTime  Field[time.Time]
	EndTime    Field[time.Time]
}

func (con *ProblemWhere) Parse(c *gin.Context) {
	con.Title.Parse(c, "title")
	con.Difficulty.Parse(c, "difficulty")
	con.Tag.Parse(c, "tag")
	con.Status.Parse(c, "status")
	con.UserId.Parse(c, "user")
	con.Page.Parse(c, "page")
	con.Size.Parse(c, "size")
	con.OrderBy.Parse(c, "order_by")
	con.Order.Parse(c, "order")
	timePreiod := &Period{}
	err := timePreiod.GetPeriod(c)
	if err == nil {
		con.StartTime.Set(timePreiod.StartTime)
		con.EndTime.Set(timePreiod.EndTime)
	}
}

func (con *ProblemWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		if con.Id.Exist() {
			whereClause["tbl_problem.id"] = con.Id.Value()
		}
		where := db.Where(whereClause)
		if con.Status.Exist() {
			where.Where("tbl_problem.status in ?", con.Status.Value())
		}
		if con.Difficulty.Exist() {
			where.Where("tbl_problem.difficulty in ?", con.Difficulty.Value())
		}
		if con.Tag.Exist() {
			where = where.Where("id IN (SELECT problem_id FROM tbl_problem_tag WHERE tag_id In(?) GROUP BY problem_id HAVING COUNT(DISTINCT tag_id) =?)", con.Tag.Value(), len(con.Tag.Value()))
		}
		if con.Title.Exist() {
			where = where.Where("tbl_problem.title LIKE ?", "%"+con.Title.Value()+"%")
		}
		if con.UserId.Exist() {
			where = where.Where("tbl_problem.id IN (SELECT problem_id FROM tbl_history WHERE user_id = ?)", con.UserId.Value())
		}
		if con.StartTime.Exist() {
			where = where.Where("tbl_problem.create_time >= ?", con.StartTime.Value())
		}
		if con.EndTime.Exist() {
			where = where.Where("tbl_problem.create_time <= ?", con.EndTime.Value())
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

func (con *ProblemWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
