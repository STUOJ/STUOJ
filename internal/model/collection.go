package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CollectionWhere struct {
	Id        Field[uint64]
	Title     Field[string]
	UserId    FieldList[uint64]
	ProblemId FieldList[uint64]
	Status    FieldList[uint64]
	StartTime Field[time.Time]
	EndTime   Field[time.Time]
	Page      Field[uint64]
	Size      Field[uint64]
	OrderBy   Field[string]
	Order     Field[string]
}

func (con *CollectionWhere) Parse(c *gin.Context) {
	con.Id.Parse(c, "id")
	con.Title.Parse(c, "title")
	con.UserId.Parse(c, "user")
	con.ProblemId.Parse(c, "problem")
	con.Status.Parse(c, "status")
	timePreiod := &Period{}
	err := timePreiod.GetPeriod(c)
	if err == nil {
		con.StartTime.Set(timePreiod.StartTime)
		con.EndTime.Set(timePreiod.EndTime)
	}
	con.Page.Parse(c, "page")
	con.Size.Parse(c, "size")
	con.OrderBy.Parse(c, "order_by")
	con.Order.Parse(c, "order")
}

func (con *CollectionWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		if con.Id.Exist() {
			whereClause["tbl_collection.id"] = con.Id.Value()
		}
		where := db.Where(whereClause)
		if con.Status.Exist() {
			where.Where("tbl_collection.status in ?", con.Status.Value())
		}
		if con.UserId.Exist() {
			where.Where("tbl_collection.user_id in ?", con.UserId.Value())
		}
		if con.ProblemId.Exist() {
			where = where.Joins("JOIN tbl_collection_problem ON tbl_collection.id = tbl_collection_problem.collection_id").
				Where("tbl_collection_problem.problem_id IN(?) GROUP BY collection_id HAVING COUNT(DISTINCT problem_id) =?", con.ProblemId.Value(), len(con.ProblemId.Value()))
		}
		if con.Title.Exist() {
			where = where.Where("tbl_collection.title LIKE ?", "%"+con.Title.Value()+"%")
		}
		if con.StartTime.Exist() {
			where = where.Where("tbl_collection.create_time >= ?", con.StartTime.Value())
		}
		if con.EndTime.Exist() {
			where = where.Where("tbl_collection.create_time <= ?", con.EndTime.Value())
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

func (con *CollectionWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
