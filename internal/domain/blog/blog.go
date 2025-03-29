package blog

import (
	"time"

	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/blog/valueobject"
	"STUOJ/internal/errors"
)

type Blog struct {
	ID         uint64
	UserID     uint64
	ProblemID  uint64
	Title      valueobject.Title
	Content    valueobject.Content
	Status     entity.BlogStatus
	CreateTime time.Time
	UpdateTime time.Time
}

func (b *Blog) verify() error {
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
		Id:         b.ID,
		UserId:     b.UserID,
		ProblemId:  b.ProblemID,
		Title:      b.Title.String(),
		Content:    b.Content.String(),
		Status:     b.Status,
		CreateTime: b.CreateTime,
		UpdateTime: b.UpdateTime,
	}
}

func (b *Blog) fromEntity(blog entity.Blog) *Blog {
	b.ID = blog.Id
	b.UserID = blog.UserId
	b.ProblemID = blog.ProblemId
	b.Title = valueobject.NewTitle(blog.Title)
	b.Content = valueobject.NewContent(blog.Content)
	b.Status = blog.Status
	b.CreateTime = blog.CreateTime
	b.UpdateTime = blog.UpdateTime
	return b
}

func (b *Blog) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.BlogId, option.OpEqual, b.ID)
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

func WithID(id uint64) Option {
	return func(b *Blog) {
		b.ID = id
	}
}

func WithUserID(userID uint64) Option {
	return func(b *Blog) {
		b.UserID = userID
	}
}

func WithProblemID(problemID uint64) Option {
	return func(b *Blog) {
		b.ProblemID = problemID
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
