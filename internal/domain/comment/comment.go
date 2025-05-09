package comment

//go:generate go run ../../../dev/gen/query_gen.go comment
//go:generate go run ../../../dev/gen/domain.go comment

import (
	"time"

	"STUOJ/internal/domain/comment/valueobject"
)

type Comment struct {
	Id         valueobject.Id
	UserId     valueobject.UserID
	BlogId     valueobject.BlogID
	Content    valueobject.Content
	Status     valueobject.Status
	CreateTime time.Time
	UpdateTime time.Time
}
