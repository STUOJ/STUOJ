package domain_test

import (
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/domain/tag/valueobject"
	"math/rand"
	"strconv"
	"testing"
)

// 生成随机标签名
func randomTagName() string {
	return "标签" + strconv.Itoa(rand.Intn(100000))
}

// 测试标签创建成功
func TestTagCreate_Success(t *testing.T) {
	tg := tag.NewTag(
		tag.WithName(randomTagName()),
	)
	id, err := tg.Create()
	if err != nil {
		t.Fatalf("创建标签失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建标签返回的ID无效")
	}
}

// 测试标签名为空时创建失败
func TestTagCreate_EmptyName(t *testing.T) {
	tg := tag.NewTag(
		tag.WithName(""),
	)
	_, err := tg.Create()
	if err == nil {
		t.Fatalf("标签名为空时应创建失败")
	}
}

// 测试标签更新成功
func TestTagUpdate_Success(t *testing.T) {
	tg := tag.NewTag(
		tag.WithName(randomTagName()),
	)
	id, err := tg.Create()
	if err != nil {
		t.Fatalf("创建标签失败: %v", err)
	}
	tg.Id = id
	tg.Name = valueobject.NewName("更新后的标签")
	err = tg.Update()
	if err != nil {
		t.Fatalf("更新标签失败: %v", err)
	}
}

// 测试更新不存在的标签
func TestTagUpdate_NotFound(t *testing.T) {
	tg := tag.NewTag(
		tag.WithId(99999999),
		tag.WithName(randomTagName()),
	)
	err := tg.Update()
	if err == nil {
		t.Fatalf("更新不存在的标签应失败")
	}
}

// 测试标签删除成功
func TestTagDelete_Success(t *testing.T) {
	tg := tag.NewTag(
		tag.WithName(randomTagName()),
	)
	id, err := tg.Create()
	if err != nil {
		t.Fatalf("创建标签失败: %v", err)
	}
	tg.Id = id
	err = tg.Delete()
	if err != nil {
		t.Fatalf("删除标签失败: %v", err)
	}
}

// 测试删除不存在的标签
func TestTagDelete_NotFound(t *testing.T) {
	tg := tag.NewTag(
		tag.WithId(99999999),
		tag.WithName(randomTagName()),
	)
	err := tg.Delete()
	if err == nil {
		t.Fatalf("删除不存在的标签应失败")
	}
}
