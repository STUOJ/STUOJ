package blog

//go:generate go run ../../../utils/gen/dto_gen.go blog
//go:generate go run ../../../utils/gen/query_gen.go blog

import (
	"fmt"
	"time"

	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/blog/valueobject"
	"STUOJ/internal/errors"
)

type Blog struct {
	Id         uint64
	UserId     uint64
	ProblemId  uint64
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
		Id:         b.Id,
		UserId:     b.UserId,
		ProblemId:  b.ProblemId,
		Title:      b.Title.String(),
		Content:    b.Content.String(),
		Status:     b.Status,
		CreateTime: b.CreateTime,
		UpdateTime: b.UpdateTime,
	}
}

func (b *Blog) fromEntity(blog entity.Blog) *Blog {
	b.Id = blog.Id
	b.UserId = blog.UserId
	b.ProblemId = blog.ProblemId
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

func (b *Blog) Create() (uint64, error) {
	b.CreateTime = time.Now()
	b.UpdateTime = time.Now()
	if err := b.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	blog, err := dao.BlogStore.Insert(b.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return blog.Id, &errors.NoError
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

type Option func(*Blog)

func NewBlog(option ...Option) *Blog {
	b := &Blog{
		Status: entity.BlogDraft,
	}
	for _, opt := range option {
		opt(b)
	}
	return b
}

func WithId(id uint64) Option {
	return func(b *Blog) {
		b.Id = id
	}
}

func WithUserId(userId uint64) Option {
	return func(b *Blog) {
		b.UserId = userId
	}
}

func WithProblemId(problemId uint64) Option {
	return func(b *Blog) {
		b.ProblemId = problemId
	}
}

func WithTitle(title string) Option {
	return func(b *Blog) {
		b.Title = valueobject.NewTitle(title)
	}
}

func WithContent(content string) Option {
	return func(b *Blog) {
		b.Content = valueobject.NewContent(content)
	}
}

func WithStatus(status entity.BlogStatus) Option {
	return func(b *Blog) {
		b.Status = status
	}
}
