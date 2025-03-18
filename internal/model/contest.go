package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContestWhere struct {
	Id        Field[uint64]
	UserId    Field[uint64]
	Status    FieldList[uint64]
	Format    FieldList[uint64]
	TeamSize  FieldList[uint8]
	StartTime Field[time.Time]
	EndTime   Field[time.Time]
	Page      Field[uint64]
	Size      Field[uint64]
	OrderBy   Field[string]
	Order     Field[string]
}

func (con *ContestWhere) Parse(c *gin.Context) {
	con.UserId.Parse(c, "user")
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
			where.Where("tbl_contest.user in ?", con.UserId.Value())
		}
		if con.Format.Exist() {
			where.Where("tbl_contest.format in ?", con.Format.Value())
		}
		if con.TeamSize.Exist() {
			where.Where("tbl_contest.team_size in ?", con.TeamSize.Value())
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
		where = briefUserJoins(where, "tbl_contest")

		return where.Select(query)
	}
}

func (con *ContestWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		if con.Page.Exist() && con.Size.Exist() {
			return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
		}
		return where(db).Offset(0).Limit(1)
	}
}
