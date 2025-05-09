package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
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
