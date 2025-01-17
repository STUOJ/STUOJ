package model

import (
	"STUOJ/internal/entity"
	"log"
	"strconv"
	"strings"

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
	Difficulty Field[entity.Difficulty]
	Status     Field[entity.ProblemStatus]
	Tag        FieldList[uint64]
	UserId     Field[uint64]
	Page       Field[uint64]
	Size       Field[uint64]
	OrderBy    Field[string]
	Order      Field[string]
}

func (con *ProblemWhere) Parse(c *gin.Context) {
	if c.Query("title") != "" {
		con.Title.Set(c.Query("title"))
	}
	if c.Query("difficulty") != "" {
		difficulty, err := strconv.Atoi(c.Query("difficulty"))
		if err != nil {
			log.Println(err)
		} else {
			con.Difficulty.Set(entity.Difficulty(difficulty))
		}
	}
	if c.Query("tag") != "" {
		tagsQuery := c.Query("tag")           // 获取URL参数 "ids"
		tags := strings.Split(tagsQuery, ",") // 将字符串分割成字符串切片

		// 假设我们需要将字符串切片转换为int切片
		var tagsInt []uint64
		for _, tagStr := range tags {
			id, err := strconv.Atoi(tagStr)
			if err != nil {
				continue
			}
			tagsInt = append(tagsInt, uint64(id))
		}
		con.Tag.Set(tagsInt)
	}
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			con.Status.Set(entity.ProblemStatus(status))
		}
	}
	if c.Query("user") != "" {
		user, err := strconv.Atoi(c.Query("user"))
		if err != nil {
			log.Println(err)
		} else {
			con.UserId.Set(uint64(user))
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
	if c.Query("order") != "" {
		order := c.Query("order")
		if c.Query("order_by") != "" {
			orderBy := c.Query("order_by")
			con.OrderBy.Set(orderBy)
			con.Order.Set(order)
		}
	}
}

func (con *ProblemWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		if con.Id.Exist() {
			whereClause["tbl_problem.id"] = con.Id.Value()
		}
		if con.Status.Exist() {
			whereClause["tbl_problem.status"] = con.Status.Value()
		}
		if con.Difficulty.Exist() {
			whereClause["tbl_problem.difficulty"] = con.Difficulty.Value()
		}

		where := db.Where(whereClause)
		if con.Tag.Exist() {
			where = where.Where("id IN (SELECT problem_id FROM tbl_problem_tag WHERE tag_id In(?) GROUP BY problem_id HAVING COUNT(DISTINCT tag_id) =?)", con.Tag.Value(), len(con.Tag.Value()))
		}
		if con.Title.Exist() {
			where = where.Where("tbl_problem.title LIKE ?", "%"+con.Title.Value()+"%")
		}
		if con.UserId.Exist() {
			where = where.Where("tbl_problem.id IN (SELECT problem_id FROM tbl_history WHERE user_id = ?)", con.UserId.Value())
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
