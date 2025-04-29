package comment

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/model"
)

func Update(req request.UpdateCommentReq, reqUser model.ReqUser) error {
	cm1 := comment.NewComment(
		comment.WithContent(req.Content),
	)

	return cm1.Update()
}
