package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	CommentAllField    = field.NewCommentField().SelectAll()
	CommentSimpleField = field.NewCommentField().
				SelectId().
				SelectUserId().
				SelectBlogId().
				SelectStatus().
				SelectCreateTime().
				SelectUpdateTime()
)
