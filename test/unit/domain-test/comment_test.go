package domain_test

import (
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/domain/comment/valueobject"
	"STUOJ/internal/infrastructure/repository/entity"
	"testing"
	"time"
)

// 生成随机内容
func randomCommentContent() valueobject.Content {
	return valueobject.NewContent("评论内容" + time.Now().Format("150405.000"))
}

// 测试评论创建成功
func TestCommentCreate_Success(t *testing.T) {
	c := &comment.Comment{
		UserId:  1,
		BlogId:  1,
		Content: randomCommentContent(),
		Status:  entity.CommentPublic,
	}
	id, err := c.Create()
	if err != nil {
		t.Fatalf("创建评论失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建评论返回的ID无效")
	}
}

// 测试用户ID为空时创建失败
func TestCommentCreate_EmptyUserId(t *testing.T) {
	c := &comment.Comment{
		UserId:  0,
		BlogId:  1,
		Content: randomCommentContent(),
		Status:  entity.CommentPublic,
	}
	_, err := c.Create()
	if err == nil {
		t.Fatalf("用户ID为空时应创建失败")
	}
}

// 测试博客ID为空时创建失败
func TestCommentCreate_EmptyBlogId(t *testing.T) {
	c := &comment.Comment{
		UserId:  1,
		BlogId:  0,
		Content: randomCommentContent(),
		Status:  entity.CommentPublic,
	}
	_, err := c.Create()
	if err == nil {
		t.Fatalf("博客ID为空时应创建失败")
	}
}

// 测试评论更新成功
func TestCommentUpdate_Success(t *testing.T) {
	c := &comment.Comment{
		UserId:  1,
		BlogId:  1,
		Content: randomCommentContent(),
		Status:  entity.CommentPublic,
	}
	id, err := c.Create()
	if err != nil {
		t.Fatalf("创建��论失败: %v", err)
	}
	c.Id = id
	c.Content = valueobject.NewContent("更新后的评论内容")
	err = c.Update()
	if err != nil {
		t.Fatalf("更新评论失败: %v", err)
	}
}

// 测试更新不存在的评论
func TestCommentUpdate_NotFound(t *testing.T) {
	c := &comment.Comment{
		Id:      99999999,
		UserId:  1,
		BlogId:  1,
		Content: randomCommentContent(),
		Status:  entity.CommentPublic,
	}
	err := c.Update()
	if err == nil {
		t.Fatalf("更新不存在的评论应失败")
	}
}

// 测试评论删除成功
func TestCommentDelete_Success(t *testing.T) {
	c := &comment.Comment{
		UserId:  1,
		BlogId:  1,
		Content: randomCommentContent(),
		Status:  entity.CommentPublic,
	}
	id, err := c.Create()
	if err != nil {
		t.Fatalf("创建评论失败: %v", err)
	}
	c.Id = id
	err = c.Delete()
	if err != nil {
		t.Fatalf("删除评论失败: %v", err)
	}
}

// 测试删除不存在的评论
func TestCommentDelete_NotFound(t *testing.T) {
	c := &comment.Comment{
		Id:      99999999,
		UserId:  1,
		BlogId:  1,
		Content: randomCommentContent(),
		Status:  entity.CommentPublic,
	}
	err := c.Delete()
	if err == nil {
		t.Fatalf("删除不存在的评论应失败")
	}
}
