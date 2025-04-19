package blog

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/blog/valueobject"
)

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

func WithId(id int64) Option {
	return func(b *Blog) {
		b.Id = id
	}
}

func WithUserId(userId int64) Option {
	return func(b *Blog) {
		b.UserId = userId
	}
}

func WithProblemId(problemId int64) Option {
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
