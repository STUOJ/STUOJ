package tag

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/tag"
)

// Insert 插入标签
func Insert(req request.CreateTagReq, reqUser request.ReqUser) (int64, error) {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return 0, err
	}

	t1 := tag.NewTag(
		tag.WithName(req.Name),
	)

	return t1.Create()
}
