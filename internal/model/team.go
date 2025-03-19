package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeamWhere struct {
	Id        Field[uint64]
	UserId    Field[uint64]
	ContestId Field[uint64]
	Name      Field[string]
	Status    FieldList[uint64]
	Page      Field[uint64]
	Size      Field[uint64]
	OrderBy   Field[string]
	Order     Field[string]
}

func (con *TeamWhere) Parse(c *gin.Context) {
	con.Id.Parse(c, "id")
	con.Name.Parse(c, "name")
	con.UserId.Parse(c, "user")
	con.ContestId.Parse(c, "contest")
	con.Status.Parse(c, "status")
	con.Page.Parse(c, "page")
	con.Size.Parse(c, "size")
	con.OrderBy.Parse(c, "order_by")
	con.Order.Parse(c, "order")
}

func (con *TeamWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		if con.Id.Exist() {
			whereClause["tbl_team.id"] = con.Id.Value()
		}
		where := db.Where(whereClause)
		if con.Status.Exist() {
			where.Where("tbl_team.status in ?", con.Status.Value())
		}
		if con.UserId.Exist() {
			where.Joins("JOIN tbl_team_user ON tbl_team.id = tbl_team_user.team_id").
				Where("tbl_team.user_id = ? OR tbl_team_user.user_id = ?", con.UserId.Value(), con.UserId.Value())
		}
		if con.Name.Exist() {
			where = where.Where("tbl_team.name LIKE ?", "%"+con.Name.Value()+"%")
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
		query := []string{"tbl_team.*"}
		query = append(query, briefUserSelect()...)
		query = append(query, "(SELECT GROUP_CONCAT(DISTINCT user_id) FROM tbl_team_user WHERE team_id = tbl_team.id) AS team_user_id")
		where = briefUserJoins(where, "tbl_team")

		return where.Select(query)
	}
}

func (con *TeamWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		if con.Page.Exist() && con.Size.Exist() {
			return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
		}
		return where(db).Offset(0).Limit(1)
	}
}

type BriefTeam struct {
	Name        string `gorm:"column:team_name"`
	Description string `gorm:"column:team_description"`
}

func briefTeamSelect() []string {
	return []string{
		"tbl_team.name as team_name",
		"tbl_team.description as team_description",
	}
}

func briefTeamJoins(db *gorm.DB, tbl string) *gorm.DB {
	db = db.Joins(fmt.Sprintf("LEFT JOIN tbl_team ON %s.team_id = tbl_team.id", tbl))
	return db
}
