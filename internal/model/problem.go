package model

import (
	"STUOJ/internal/entity"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProblemWhere struct {
	Id          FieldList[uint64] // 只提供在service层供其他模块使用，不提供直接从用户层解析，避免垮权限查询
	Title       Field[string]
	Difficulty  FieldList[uint64]
	Status      FieldList[uint64]
	Tag         FieldList[uint64]
	UserId      Field[uint64]
	ScoreUserId Field[uint64]
	Page        Field[uint64]
	Size        Field[uint64]
	OrderBy     Field[string]
	Order       Field[string]
	StartTime   Field[time.Time]
	EndTime     Field[time.Time]
	Detail      Field[bool]
	Testcases   Field[bool]
	Solutions   Field[bool]
}

func (con *ProblemWhere) Parse(c *gin.Context) {
	con.Title.Parse(c, "title")
	con.Difficulty.Parse(c, "difficulty")
	con.Tag.Parse(c, "tag")
	con.Status.Parse(c, "status")
	con.ScoreUserId.Parse(c, "score_user_id")
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
	con.Detail.Parse(c, "detail")
	con.Testcases.Parse(c, "testcases")
	con.Solutions.Parse(c, "solutions")
}

func (con *ProblemWhere) GenerateWhereWithNoPage() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		where := db.Where(whereClause)
		if con.Id.Exist() {
			where.Where("tbl_problem.id in ?", con.Id.Value())
		}
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

		query := []string{"tbl_problem.id", "tbl_problem.title", "tbl_problem.source", "tbl_problem.difficulty", "tbl_problem.time_limit", "tbl_problem.memory_limit", "tbl_problem.status", "tbl_problem.create_time", "tbl_problem.update_time"}
		if con.Detail.Exist() && con.Detail.Value() {
			query = append(query,
				"tbl_problem.description", "tbl_problem.input", "tbl_problem.output", "tbl_problem.sample_input", "tbl_problem.sample_output", "tbl_problem.hint",
			)
		}
		query = append(query,
			"(SELECT GROUP_CONCAT(DISTINCT tbl_problem_tag.tag_id) FROM tbl_problem_tag WHERE problem_id = tbl_problem.id) AS problem_tag_id",
			"(SELECT GROUP_CONCAT(DISTINCT tbl_history.user_id) FROM tbl_history WHERE problem_id = tbl_problem.id) AS problem_user_id",
			"(SELECT GROUP_CONCAT(DISTINCT tbl_collection_problem.collection_id) FROM tbl_collection_problem WHERE problem_id = tbl_problem.id) AS problem_collection_id",
			"(SELECT GROUP_CONCAT(DISTINCT tbl_collection_user.user_id) FROM tbl_collection_problem JOIN tbl_collection_user ON tbl_collection_problem.collection_id = tbl_collection_user.collection_id WHERE tbl_collection_problem.problem_id = tbl_problem.id AND EXISTS (SELECT 1 FROM tbl_collection WHERE tbl_collection.id = tbl_collection_user.collection_id AND tbl_collection.user_id IN (SELECT DISTINCT user_id FROM tbl_history WHERE problem_id = tbl_problem.id))) AS problem_collection_user_id",
		)
		if con.ScoreUserId.Exist() {
			query = append(query, fmt.Sprintf(
				"(SELECT MAX(tbl_submission.score) FROM tbl_submission WHERE tbl_submission.problem_id = tbl_problem.id AND tbl_submission.user_id = %d) AS problem_user_score", con.ScoreUserId.Value()),
				fmt.Sprintf("EXISTS (SELECT 1 FROM tbl_submission WHERE tbl_submission.problem_id = tbl_problem.id AND tbl_submission.user_id = %d) AS has_user_submission", con.ScoreUserId.Value()),
			)
		}
		queryStr := strings.Join(query, ", ")
		return where.Select(queryStr)
	}
}

func (con *ProblemWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	where := con.GenerateWhereWithNoPage()
	return func(db *gorm.DB) *gorm.DB {
		if con.Page.Exist() && con.Size.Exist() {
			return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
		}
		return where(db).Offset(0).Limit(1)
	}
}

type BriefProblem struct {
	ProblemTitle      string               `gorm:"column:problem_title"`
	ProblemStatus     entity.ProblemStatus `gorm:"column:problem_status"`
	ProblemDifficulty entity.Difficulty    `gorm:"column:problem_difficulty"`
}

func briefProblemSelect() []string {
	return []string{
		"tbl_problem.title as problem_title",
		"tbl_problem.status as problem_status",
		"tbl_problem.difficulty as problem_difficulty",
	}
}

func briefProblemJoins(db *gorm.DB, tbl string) *gorm.DB {
	db = db.Joins(fmt.Sprintf("LEFT JOIN tbl_problem ON %s.problem_id = tbl_problem.id", tbl))
	return db
}
