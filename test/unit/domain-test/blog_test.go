package domain_test

import (
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/domain/blog/valueobject"
	"STUOJ/internal/infrastructure/persistence/entity"
	"testing"
	"time"
)

// 生成随机标题
func randomTitle() string {
	return "博客标题" + time.Now().Format("150405.000")
}

// 生成随机内容
func randomContent() string {
	return "博客内容" + time.Now().Format("150405.000")
}

// 测试博客创建成功
func TestBlogCreate_Success(t *testing.T) {
	b := blog.NewBlog(
		blog.WithUserId(1),
		blog.WithProblemId(0),
		blog.WithTitle(randomTitle()),
		blog.WithContent(randomContent()),
		blog.WithStatus(entity.BlogPublic),
	)
	id, err := b.Create()
	if err != nil {
		t.Fatalf("创建博客失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建博客返回的ID无效")
	}
}

// 测试用户ID为空时创建失败
func TestBlogCreate_EmptyUserId(t *testing.T) {
	b := blog.NewBlog(
		blog.WithUserId(0),
		blog.WithProblemId(0),
		blog.WithTitle(randomTitle()),
		blog.WithContent(randomContent()),
		blog.WithStatus(entity.BlogPublic),
	)
	_, err := b.Create()
	if err == nil {
		t.Fatalf("用户ID为空时应创建失败")
	}
}

// 测试博客更新成功
func TestBlogUpdate_Success(t *testing.T) {
	b := blog.NewBlog(
		blog.WithUserId(1),
		blog.WithProblemId(0),
		blog.WithTitle(randomTitle()),
		blog.WithContent(randomContent()),
		blog.WithStatus(entity.BlogPublic),
	)
	id, err := b.Create()
	if err != nil {
		t.Fatalf("创建博客失败: %v", err)
	}
	b.Id.Set(id) // 使用Valueobject设置ID
	b.Title = valueobject.NewTitle("更新后的标题")
	err = b.Update()
	if err != nil {
		t.Fatalf("更新博客失败: %v", err)
	}
}

// 测试更新不存在的博客
func TestBlogUpdate_NotFound(t *testing.T) {
	b := blog.NewBlog(
		blog.WithId(99999999),
		blog.WithUserId(1),
		blog.WithProblemId(0),
		blog.WithTitle(randomTitle()),
		blog.WithContent(randomContent()),
		blog.WithStatus(entity.BlogPublic),
	)
	err := b.Update()
	if err == nil {
		t.Fatalf("更新不存在的博客应失败")
	}
}

// 测试博客删除成功
func TestBlogDelete_Success(t *testing.T) {
	b := blog.NewBlog(
		blog.WithUserId(1),
		blog.WithProblemId(0),
		blog.WithTitle(randomTitle()),
		blog.WithContent(randomContent()),
		blog.WithStatus(entity.BlogPublic),
	)
	id, err := b.Create()
	if err != nil {
		t.Fatalf("创建博客失败: %v", err)
	}
	b.Id = valueobject.NewId(id)
	err = b.Delete()
	if err != nil {
		t.Fatalf("删除博客失败: %v", err)
	}
}

// 测试删除不存在的博客
func TestBlogDelete_NotFound(t *testing.T) {
	b := blog.NewBlog(
		blog.WithId(99999999),
		blog.WithUserId(1),
		blog.WithProblemId(0),
		blog.WithTitle(randomTitle()),
		blog.WithContent(randomContent()),
		blog.WithStatus(entity.BlogPublic),
	)
	err := b.Delete()
	if err == nil {
		t.Fatalf("删除不存在的博客应失败")
	}
}
