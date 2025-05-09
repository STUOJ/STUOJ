package tag

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/tag"
)

func Delete(id int64, reqUser request.ReqUser) error {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return err
	}

	t1 := tag.NewTag(
		tag.WithId(id),
	)

	return t1.Delete()
}
