package domain_test

import (
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/domain/collection/valueobject"
	"STUOJ/internal/infrastructure/persistence/entity"
	"testing"
	"time"
)

// 生成随机标题
func randomCollectionTitle() string {
	return "题单标题" + time.Now().Format("150405.000")
}

// 生成随机描述
func randomCollectionDescription() string {
	return "题单描述" + time.Now().Format("150405.000")
}

// 测试题单创建成功
func TestCollectionCreate_Success(t *testing.T) {
	c := collection.NewCollection(
		collection.WithUserId(1),
		collection.WithTitle(randomCollectionTitle()),
		collection.WithDescription(randomCollectionDescription()),
		collection.WithStatus(entity.CollectionPublic),
	)
	id, err := c.Create()
	if err != nil {
		t.Fatalf("创建题单失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建题单返回的ID无效")
	}
}

// 测试用户ID为空时创建失败
func TestCollectionCreate_EmptyUserId(t *testing.T) {
	c := collection.NewCollection(
		collection.WithUserId(0),
		collection.WithTitle(randomCollectionTitle()),
		collection.WithDescription(randomCollectionDescription()),
		collection.WithStatus(entity.CollectionPublic),
	)
	_, err := c.Create()
	if err == nil {
		t.Fatalf("用户ID为空时应创建失败")
	}
}

// 测试题单更新成功
func TestCollectionUpdate_Success(t *testing.T) {
	c := collection.NewCollection(
		collection.WithUserId(1),
		collection.WithTitle(randomCollectionTitle()),
		collection.WithDescription(randomCollectionDescription()),
		collection.WithStatus(entity.CollectionPublic),
	)
	id, err := c.Create()
	if err != nil {
		t.Fatalf("创建题单失败: %v", err)
	}
	c.Id = id
	c.Description = valueobject.NewDescription("更新后的描述")
	err = c.Update()
	if err != nil {
		t.Fatalf("更新题单失败: %v", err)
	}
}

// 测试更新不存在的题单
func TestCollectionUpdate_NotFound(t *testing.T) {
	c := collection.NewCollection(
		collection.WithId(99999999),
		collection.WithUserId(1),
		collection.WithTitle(randomCollectionTitle()),
		collection.WithDescription(randomCollectionDescription()),
		collection.WithStatus(entity.CollectionPublic),
	)
	err := c.Update()
	if err == nil {
		t.Fatalf("更新不存在的题单应失败")
	}
}

// 测试题单删除成功
func TestCollectionDelete_Success(t *testing.T) {
	c := collection.NewCollection(
		collection.WithUserId(1),
		collection.WithTitle(randomCollectionTitle()),
		collection.WithDescription(randomCollectionDescription()),
		collection.WithStatus(entity.CollectionPublic),
	)
	id, err := c.Create()
	if err != nil {
		t.Fatalf("创建题单失败: %v", err)
	}
	c.Id = id
	err = c.Delete()
	if err != nil {
		t.Fatalf("删除题单失败: %v", err)
	}
}

// 测试删除不存在的题单
func TestCollectionDelete_NotFound(t *testing.T) {
	c := collection.NewCollection(
		collection.WithId(99999999),
		collection.WithUserId(1),
		collection.WithTitle(randomCollectionTitle()),
		collection.WithDescription(randomCollectionDescription()),
		collection.WithStatus(entity.CollectionPublic),
	)
	err := c.Delete()
	if err == nil {
		t.Fatalf("删除不存在的题单应失败")
	}
}
