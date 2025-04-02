package comment

import (
	"fmt"
	"time"

	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/comment/valueobject"
	"STUOJ/internal/errors"
)

type Comment struct {
	Id         uint64
	UserId     uint64
	BlogId     uint64
	Content    valueobject.Content
	Status     entity.CommentStatus
	CreateTime time.Time
	UpdateTime time.Time
}

func (c *Comment) verify() error {
	if c.UserId == 0 {
		return fmt.Errorf("用户ID不能为空")
	}
	if c.BlogId == 0 {
		return fmt.Errorf("博客ID不能为空")
	}
	if !entity.CommentStatus(c.Status).IsValid() {
		return fmt.Errorf("评论状态不合法")
	}
	if err := c.Content.Verify(); err != nil {
		return err
	}
	return nil
}

func (c *Comment) toEntity() entity.Comment {
	return entity.Comment{
		Id:         c.Id,
		UserId:     c.UserId,
		BlogId:     c.BlogId,
		Content:    c.Content.String(),
		Status:     c.Status,
		CreateTime: c.CreateTime,
		UpdateTime: c.UpdateTime,
	}
}

func (c *Comment) fromEntity(comment entity.Comment) *Comment {
	c.Id = comment.Id
	c.UserId = comment.UserId
	c.BlogId = comment.BlogId
	c.Content = valueobject.NewContent(comment.Content)
	c.Status = comment.Status
	c.CreateTime = comment.CreateTime
	c.UpdateTime = comment.UpdateTime
	return c
}

func (c *Comment) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.CommentId, option.OpEqual, c.Id)
	return options
}

func (c *Comment) Create() (uint64, error) {
	c.CreateTime = time.Now()
	c.UpdateTime = time.Now()
	if err := c.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	comment, err := dao.CommentStore.Insert(c.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return comment.Id, &errors.NoError
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
	return &errors.NoError
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
	return &errors.NoError
}

type Option func(*Comment)

func NewComment(option ...Option) *Comment {
	c := &Comment{
		Status: entity.CommentPublic,
	}
	for _, opt := range option {
		opt(c)
	}
	return c
}

func WithId(id uint64) Option {
	return func(c *Comment) {
		c.Id = id
	}
}

func WithUserId(userId uint64) Option {
	return func(c *Comment) {
		c.UserId = userId
	}
}

func WithBlogId(blogId uint64) Option {
	return func(c *Comment) {
		c.BlogId = blogId
	}
}

func WithContent(content string) Option {
	return func(c *Comment) {
		c.Content = valueobject.NewContent(content)
	}
}

func WithStatus(status entity.CommentStatus) Option {
	return func(c *Comment) {
		c.Status = status
	}
}
