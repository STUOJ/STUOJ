package comment

//go:generate go run ../../../dev/gen/dto_gen.go comment
//go:generate go run ../../../dev/gen/query_gen.go comment
//go:generate go run ../../../dev/gen/builder.go comment

import (
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
	"fmt"
	"time"

	"STUOJ/internal/domain/comment/valueobject"
)

type Comment struct {
	Id         int64
	UserId     int64
	BlogId     int64
	Content    valueobject.Content
	Status     entity.CommentStatus
	CreateTime time.Time
	UpdateTime time.Time
}

func (c *Comment) verify() error {
	if c.UserId == 0 {
		return fmt.Errorf("用户Id不能为空")
	}
	if c.BlogId == 0 {
		return fmt.Errorf("博客Id不能为空")
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
		Id:         uint64(c.Id),
		UserId:     uint64(c.UserId),
		BlogId:     uint64(c.BlogId),
		Content:    c.Content.String(),
		Status:     c.Status,
		CreateTime: c.CreateTime,
		UpdateTime: c.UpdateTime,
	}
}

func (c *Comment) fromEntity(comment entity.Comment) *Comment {
	c.Id = int64(comment.Id)
	c.UserId = int64(comment.UserId)
	c.BlogId = int64(comment.BlogId)
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
