package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type ContestWhere struct {
	Id           Field[uint64]
	UserId       Field[uint64]
	CollectionId Field[uint64]
	Status       FieldList[uint64]
	Format       FieldList[uint64]
	TeamSize     Field[uint8]
	StartTime    Field[time.Time]
	EndTime      Field[time.Time]
	Page         Field[uint64]
	Size         Field[uint64]
	OrderBy      Field[string]
	Order        Field[string]
}

func (con *ContestWhere) Parse(c *gin.Context) {
	con.Id.Parse(c, "id")
	con.UserId.Parse(c, "user")
	con.CollectionId.Parse(c, "collection_id")
	con.Status.Parse(c, "status")
	con.Format.Parse(c, "format")
	con.TeamSize.Parse(c, "team_size")
	period := &Period{}
	err := period.GetPeriod(c)
	if err == nil {
		con.StartTime.Set(period.StartTime)
		con.EndTime.Set(period.EndTime)
	}
	con.Page.Parse(c, "page")
	con.Size.Parse(c, "size")
	con.OrderBy.Parse(c, "order_by")
	con.Order.Parse(c, "order")
}

// TODO 补全比赛查询条件
func (con *ContestWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		if con.Id.Exist() {
			whereClause["tbl_contest.id"] = con.Id.Value()
		}
		where := db.Where(whereClause)
		if con.Status.Exist() {
			where.Where("tbl_contest.status in ?", con.Status.Value())
		}
		if con.UserId.Exist() {
			where.Joins("JOIN tbl_contest_user ON tbl_contest.id = tbl_contest_user.contest_id").
				Where("tbl_contest.user_id = ? OR tbl_contest_user.user_id = ?", con.UserId.Value(), con.UserId.Value())
		}
		if con.StartTime.Exist() {
			where = where.Where("tbl_contest.create_time >= ?", con.StartTime.Value())
		}
		if con.EndTime.Exist() {
			where = where.Where("tbl_contest.create_time <= ?", con.EndTime.Value())
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
		query := []string{"tbl_contest.*"}
		query = append(query, briefUserSelect()...)
		query = append(query, "(SELECT GROUP_CONCAT(DISTINCT user_id) FROM tbl_contest_user WHERE contest_id = tbl_contest.id) AS contest_user_id",
			"(SELECT GROUP_CONCAT(DISTINCT problem_id) FROM tbl_contest_problem WHERE contest_id = tbl_contest.id) AS contest_problem_id",
		)
		where = briefUserJoins(where, "tbl_contest")

		return where.Select(query)
	}
}

func (con *ContestWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
