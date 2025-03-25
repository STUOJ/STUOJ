package query

import "STUOJ/internal/db/entity/field"

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
