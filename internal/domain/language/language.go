package language

//go:generate go run ../../../dev/gen/query_gen.go language
//go:generate go run ../../../dev/gen/domain.go language

import (
	"STUOJ/internal/domain/language/valueobject"
)

type Language struct {
	Id     valueobject.Id
	Name   valueobject.Name
	Serial valueobject.Serial
	MapId  valueobject.MapId
	Status valueobject.Status
}
