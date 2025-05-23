package problem

//go:generate go run ../../../dev/gen/query_gen.go problem
//go:generate go run ../../../dev/gen/domain.go problem

import (
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/pkg/errors"
	"time"

	"STUOJ/internal/domain/problem/valueobject"
)

type Problem struct {
	Id           valueobject.Id
	Title        valueobject.Title
	Source       valueobject.Source
	Difficulty   valueobject.Difficulty
	TimeLimit    valueobject.TimeLimit
	MemoryLimit  valueobject.MemoryLimit
	Description  valueobject.Description
	Input        valueobject.Input
	Output       valueobject.Output
	SampleInput  valueobject.Input
	SampleOutput valueobject.Output
	Hint         valueobject.Description
	Status       valueobject.Status
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (p *Problem) UpdateTags(tagIds []int64) error {
	var err error
	options := querycontext.ProblemTagQueryContext{}
	options.ProblemId.Set(p.Id.Value())
	err = dao.ProblemTagStore.Delete(options.GenerateOptions())
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	var errs []error
	for _, id := range tagIds {
		_, err = dao.ProblemTagStore.Insert(entity.ProblemTag{
			ProblemId: uint64(p.Id.Value()),
			TagId:     uint64(id),
		})
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.ErrInternalServer.WithErrors(errs)
	}
	return nil
}
