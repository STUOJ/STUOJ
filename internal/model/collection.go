package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CollectionWhere struct {
	Id        Field[uint64]
	Title     Field[string]
	UserId    Field[uint64]
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
			where.Joins("JOIN tbl_collection_user ON tbl_collection.id = tbl_collection_user.collection_id").
				Where("tbl_collection.user_id = ? OR tbl_collection_user.user_id = ?", con.UserId.Value(), con.UserId.Value())
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
		query := []string{"tbl_collection.*"}
		query = append(query, briefUserSelect()...)
		query = append(query, "(SELECT GROUP_CONCAT(DISTINCT user_id) FROM tbl_collection_user WHERE collection_id = tbl_collection.id) AS collection_user_id",
			"(SELECT GROUP_CONCAT(DISTINCT problem_id ORDER BY serial ASC) FROM tbl_collection_problem WHERE collection_id = tbl_collection.id) AS collection_problem_id",
		)
		where = briefUserJoins(where, "tbl_collection")

		return where.Select(query)
	}
}

func (con *CollectionWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		if con.Page.Exist() && con.Size.Exist() {
			return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
		}
		return where(db).Offset(0).Limit(1)
	}
}
