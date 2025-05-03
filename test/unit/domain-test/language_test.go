package domain_test

import (
	"STUOJ/internal/domain/language"
	"STUOJ/internal/domain/language/valueobject"
	"STUOJ/internal/infrastructure/repository/entity"
	"math/rand"
	"testing"
)

// 生成随机语言名
func randomLangName() valueobject.Name {
	return valueobject.NewName("Lang_" + string(rune(rand.Intn(26)+'A')))
}

// 测试语言创建成功
func TestLanguageCreate_Success(t *testing.T) {
	l := &language.Language{
		Name:   randomLangName(),
		Serial: uint16(rand.Intn(100)),
		MapId:  uint32(rand.Intn(1000)),
		Status: entity.LanguageEnabled,
	}
	id, err := l.Create()
	if err != nil {
		t.Fatalf("创建语言失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建语言返回的ID无效")
	}
}

// 测试语言名为空时创建失败
func TestLanguageCreate_EmptyName(t *testing.T) {
	l := &language.Language{
		Name:   valueobject.NewName(""),
		Serial: 1,
		MapId:  1,
		Status: entity.LanguageEnabled,
	}
	_, err := l.Create()
	if err == nil {
		t.Fatalf("语言名为空时应创建失败")
	}
}

// 测试语言更新成功
func TestLanguageUpdate_Success(t *testing.T) {
	l := &language.Language{
		Name:   randomLangName(),
		Serial: uint16(rand.Intn(100)),
		MapId:  uint32(rand.Intn(1000)),
		Status: entity.LanguageEnabled,
	}
	id, err := l.Create()
	if err != nil {
		t.Fatalf("创建语言失败: %v", err)
	}
	l.Id = id
	l.Status = entity.LanguageDisabled
	err = l.Update()
	if err != nil {
		t.Fatalf("更新语言失败: %v", err)
	}
}

// 测试更新不存在的语言
func TestLanguageUpdate_NotFound(t *testing.T) {
	l := &language.Language{
		Id:     99999999,
		Name:   randomLangName(),
		Serial: 1,
		MapId:  1,
		Status: entity.LanguageEnabled,
	}
	err := l.Update()
	if err == nil {
		t.Fatalf("更新不存在的语言应失败")
	}
}

// 测试语言删除成功
func TestLanguageDelete_Success(t *testing.T) {
	l := &language.Language{
		Name:   randomLangName(),
		Serial: uint16(rand.Intn(100)),
		MapId:  uint32(rand.Intn(1000)),
		Status: entity.LanguageEnabled,
	}
	id, err := l.Create()
	if err != nil {
		t.Fatalf("创建语言失败: %v", err)
	}
	l.Id = id
	err = l.Delete()
	if err != nil {
		t.Fatalf("删除语言失败: %v", err)
	}
}

// 测试删除不存在的语言
func TestLanguageDelete_NotFound(t *testing.T) {
	l := &language.Language{
		Id:     99999999,
		Name:   randomLangName(),
		Serial: 1,
		MapId:  1,
		Status: entity.LanguageEnabled,
	}
	err := l.Delete()
	if err == nil {
		t.Fatalf("删除不存在的语言应失败")
	}
}
