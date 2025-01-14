package comment

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

type CommentPage struct {
	Comments []entity.Comment `json:"comments"`
	model.Page
}

func Select(condition model.CommentWhere, userId uint64, admin ...bool) (CommentPage, error) {
	if !condition.Page.Exist() {
		condition.Page.Set(1)
	}
	if !condition.Size.Exist() {
		condition.Size.Set(10)
	}
	comments, err := dao.SelectComments(condition)
	if err != nil {
		log.Println(err)
		return CommentPage{}, errors.New("获取评论失败")
	}
	if len(admin) == 0 || !admin[0] {
		var publicComment []entity.Comment
		for _, comment := range comments {
			if comment.Status >= entity.CommentPublic || comment.UserId == userId {
				publicComment = append(publicComment, comment)
			}
		}
		comments = publicComment
	}
	count, err := dao.CountComments(condition)
	if err != nil {
		log.Println(err)
		return CommentPage{}, errors.New("获取统计失败")
	}
	cPage := CommentPage{
		Comments: comments,
		Page: model.Page{
			Page:  condition.Page.Value(),
			Size:  condition.Size.Value(),
			Total: count,
		},
	}

	return cPage, nil
}
