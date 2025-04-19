package request

type QueryCollectionParams struct {
	EndTime   *string `form:"end-time,omitempty"`
	Order     *string `form:"order,omitempty"`
	OrderBy   *string `form:"order_by,omitempty"`
	Page      *int64  `form:"page,omitempty"`
	Problem   *string `form:"problem,omitempty"`
	Size      *int64  `form:"size,omitempty"`
	StartTime *string `form:"start-time,omitempty"`
	Status    *string `form:"status,omitempty"`
	Title     *string `form:"title,omitempty"`
	User      *string `form:"user,omitempty"`
}

type CreateCollectionReq struct {
	Description string `json:"description"`
	Status      int64  `json:"status"`
	Title       string `json:"title"`
}

type UpdateCollectionReq struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Status      int64  `json:"status"`
	Title       string `json:"title"`
}

type UpdateCollectionProblemReq struct {
	CollectionId int64 `json:"collection_id"`
	Problem      []struct {
		ProblemId int64 `json:"problem_id"`
		Serial    int64 `json:"serial"`
	} `json:"problem"`
}

type UpdateCollectionUserReq struct {
	CollectionId int64   `json:"collection_id"`
	UserIds      []int64 `json:"user_ids"`
}
