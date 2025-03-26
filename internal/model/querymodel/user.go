package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type UserQuery struct {
	Id        model.FieldList[uint64]
	Username  model.Field[string]
	Role      model.FieldList[uint8]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      model.QueryPage
	Sort      model.QuerySort
}

func (query *UserQuery) Parse(c *gin.Context) {
	query.Username.Parse(c, "username")
	query.Id.Parse(c, "id")
	query.Role.Parse(c, "role")
	timePreiod := &model.Period{}
	err := timePreiod.GetPeriod(c)
	if err == nil {
		query.StartTime.Set(timePreiod.StartTime)
		query.EndTime.Set(timePreiod.EndTime)
	}
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *UserQuery) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.UserId, option.OpIn, query.Id.Value())
	}
	if query.Username.Exist() {
		options.Filters.Add(field.UserUsername, option.OpLike, query.Username.Value())
	}
	if query.Role.Exist() {
		options.Filters.Add(field.UserRole, option.OpIn, query.Role.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.UserCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.UserCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
