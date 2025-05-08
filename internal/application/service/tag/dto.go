package tag

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/infrastructure/repository/querycontext"
)

// domain2Resp 将领域模型转换为响应模型
func domain2Resp(t tag.Tag) response.TagData {
	return response.TagData{
		Id:   t.Id.Value(),
		Name: t.Name.Value(),
	}
}

// params2Query 将请求参数转换为查询上下文
func params2Query(params request.QueryTagParams) querycontext.TagQueryContext {
	qc := querycontext.TagQueryContext{}

	// 设置分页
	if params.Page != nil && params.Size != nil {
		qc.Page.Page = *params.Page
		qc.Page.PageSize = *params.Size
	}

	// 设置名称查询
	if params.Name != nil && *params.Name != "" {
		qc.Name.Add(*params.Name)
	}

	return qc
}
