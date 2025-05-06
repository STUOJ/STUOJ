package blog

//go:generate go run ../../../dev/gen/query_gen.go blog
//go:generate go run ../../../dev/gen/builder.go blog

import (
	"time"

	"STUOJ/internal/domain/blog/valueobject"
)

type Blog struct {
	Id         valueobject.Id
	UserId     valueobject.UserId
	ProblemId  valueobject.ProblemId
	Title      valueobject.Title
	Content    valueobject.Content
	Status     valueobject.Status
	CreateTime time.Time
	UpdateTime time.Time
}
