package tag

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/model"
)

type TagPage struct {
	Tags []response.TagData `json:"tags"`
	model.Page
}

// Select 查询所有用户
func Select(params request.QueryTagParams, reqTag model.ReqUser) (TagPage, error) {
	var resp TagPage

	// 查询
	qc := params2Query(params)
	qc.Field.SelectAll()
	tags, _, err := tag.Query.Select(qc)
	if err != nil {
		return resp, err
	}

	for _, t := range tags {
		respTag := domain2Resp(t)
		resp.Tags = append(resp.Tags, respTag)
	}

	resp.Page.Page = qc.Page.Page
	resp.Size = qc.Page.PageSize
	resp.Page.Total, err = Count(params)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
