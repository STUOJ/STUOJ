package tag

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/model"
)

type TagPage struct {
	Tags []response.TagData `json:"tags"`
	model.Page
}

// Select 分页查询标签列表
func Select(params request.QueryTagParams, reqUser model.ReqUser) (TagPage, error) {
	var res TagPage

	// 转换查询参数
	tagQueryContext := params2Query(params)

	// 设置查询字段
	tagQueryContext.Field.SelectId().SelectName()

	// 执行查询
	tagDomain, _, err := tag.Query.Select(tagQueryContext)
	if err != nil {
		return TagPage{}, err
	}

	// 转换结果
	res.Tags = make([]response.TagData, len(tagDomain))
	for i, v := range tagDomain {
		res.Tags[i] = domain2Resp(v)
	}

	// 设置分页信息
	res.Page = model.Page{
		Page: tagQueryContext.Page.Page,
		Size: tagQueryContext.Page.PageSize,
	}

	// 获取总数
	total, _ := tag.Query.Count(tagQueryContext)
	res.Page.Total = total

	return res, nil
}

// SelectById 根据ID查询单个标签
func SelectById(id int64, reqUser model.ReqUser) (response.TagData, error) {
	var resp response.TagData

	// 查询
	qc := querycontext.TagQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId().SelectName()
	tag0, _, err := tag.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	resp = domain2Resp(tag0)
	return resp, nil
}
