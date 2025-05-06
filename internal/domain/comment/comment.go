package comment

//go:generate go run ../../../dev/gen/query_gen.go comment
//go:generate go run ../../../dev/gen/builder.go comment

import (
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
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

func (c *Comment) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.CommentId, option.OpEqual, c.Id)
	return options
}

func (c *Comment) Create() (int64, error) {
	c.CreateTime = time.Now()
	c.UpdateTime = time.Now()
	if err := c.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	comment, err := dao.CommentStore.Insert(c.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(comment.Id), nil
}

func (c *Comment) Update() error {
	var err error
	options := c.toOption()
	_, err = dao.CommentStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	c.UpdateTime = time.Now()
	if err := c.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.CommentStore.Updates(c.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (c *Comment) Delete() error {
	options := c.toOption()
	_, err := dao.CommentStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.CommentStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}
