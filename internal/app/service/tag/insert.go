package tag

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/model"
)

// Insert 插入标签
func Insert(req request.CreateTagReq, reqUser model.ReqUser) (int64, error) {
	err := isPermission(reqUser)
	if err != nil {
		return 0, err
	}

	t1 := tag.NewTag(
		tag.WithName(req.Name),
	)

	return t1.Create()
}
