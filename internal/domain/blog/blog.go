package blog

//go:generate go run ../../../dev/gen/dto_gen.go blog
//go:generate go run ../../../dev/gen/query_gen.go blog

import (
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
	"fmt"
	"time"

	"STUOJ/internal/domain/blog/valueobject"
)

type Blog struct {
	Id         int64
	UserId     int64
	ProblemId  int64
	Title      valueobject.Title
	Content    valueobject.Content
	Status     entity.BlogStatus
	CreateTime time.Time
	UpdateTime time.Time
}

func (b *Blog) verify() error {
	if b.UserId == 0 {
		return fmt.Errorf("用户Id不能为空")
	}
	if !entity.BlogStatus(b.Status).IsValid() {
		return fmt.Errorf("博客状态不合法")
	}
	if err := b.Title.Verify(); err != nil {
		return err
	}
	if err := b.Content.Verify(); err != nil {
		return err
	}
	return nil
}

func (b *Blog) toEntity() entity.Blog {
	return entity.Blog{
		Id:         uint64(b.Id),
		UserId:     uint64(b.UserId),
		ProblemId:  uint64(b.ProblemId),
		Title:      b.Title.String(),
		Content:    b.Content.String(),
		Status:     b.Status,
		CreateTime: b.CreateTime,
		UpdateTime: b.UpdateTime,
	}
}

func (b *Blog) fromEntity(blog entity.Blog) *Blog {
	b.Id = int64(blog.Id)
	b.UserId = int64(blog.UserId)
	b.ProblemId = int64(blog.ProblemId)
	b.Title = valueobject.NewTitle(blog.Title)
	b.Content = valueobject.NewContent(blog.Content)
	b.Status = blog.Status
	b.CreateTime = blog.CreateTime
	b.UpdateTime = blog.UpdateTime
	return b
}

func (b *Blog) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.BlogId, option.OpEqual, b.Id)
	return options
}

func (b *Blog) Create() (int64, error) {
	b.CreateTime = time.Now()
	b.UpdateTime = time.Now()
	if err := b.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	blog, err := dao.BlogStore.Insert(b.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(blog.Id), &errors.NoError
}

func (b *Blog) Update() error {
	var err error
	options := b.toOption()
	_, err = dao.BlogStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	b.UpdateTime = time.Now()
	if err := b.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.BlogStore.Updates(b.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

func (b *Blog) Delete() error {
	options := b.toOption()
	_, err := dao.BlogStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.BlogStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}
