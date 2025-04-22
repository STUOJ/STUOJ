package tag

import "STUOJ/internal/domain/tag/valueobject"

type Option func(*Tag)

func NewTag(options ...Option) *Tag {
	tag := &Tag{}
	for _, option := range options {
		option(tag)
	}
	return tag
}

func WithId(id int64) Option {
	return func(t *Tag) {
		t.Id = id
	}
}

func WithName(name string) Option {
	return func(t *Tag) {
		t.Name = valueobject.NewName(name)
	}
}
