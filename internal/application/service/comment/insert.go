package comment

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/infrastructure/persistence/entity"
)

// Insert 插入评论
func Insert(req request.CreateCommentReq, reqUser request.ReqUser) (int64, error) {
	cm1 := comment.NewComment(
		comment.WithBlogId(req.BlogId),
		comment.WithUserId(reqUser.Id),
		comment.WithContent(req.Content),
		comment.WithStatus(entity.CommentPublic),
	)

	return cm1.Create()
}
