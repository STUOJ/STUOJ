package tag

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/tag"
)

func Update(req request.UpdateTagReq, reqUser request.ReqUser) error {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return err
	}

	t1 := tag.NewTag(
		tag.WithId(req.Id),
		tag.WithName(req.Name),
	)

	return t1.Update()
}
