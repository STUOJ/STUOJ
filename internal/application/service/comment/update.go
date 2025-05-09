package comment

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/comment"
)

func Update(req request.UpdateCommentReq, reqUser request.ReqUser) error {
	cm1 := comment.NewComment(
		comment.WithId(req.Id),
		comment.WithContent(req.Content),
	)

	return cm1.Update()
}
