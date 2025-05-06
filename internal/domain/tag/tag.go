package tag

//go:generate go run ../../../dev/gen/query_gen.go tag
//go:generate go run ../../../dev/gen/domain.go tag

import (
	"STUOJ/internal/domain/tag/valueobject"
)

type Tag struct {
	Id   valueobject.Id
	Name valueobject.Name
}
