package tag

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/model"
)

// Update 根据ID更新标签数据
func Update(req request.UpdateTagReq, reqUser model.ReqUser) error {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return err
	}

	// 查询
	qc := querycontext.TagQueryContext{}
	qc.Id.Add(req.Id)
	qc.Field.SelectAll()
	t0, _, err := tag.Query.SelectOne(qc)
	if err != nil {
		return err
	}

	t1 := tag.NewTag(
		tag.WithId(t0.Id),
		tag.WithName(req.Name),
	)

	return t1.Update()
}
