package entity_test

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"testing"
)

func TestEnum(t *testing.T) {
	var status entity.LanguageStatus
	status = 2
	// 直接调用接口方法（无需断言，因 LanguageStatus 已实现接口）
	if !status.IsValid() {
		t.Error("status is invalid")
	}

	// 若需显式断言（可选）
	s, ok := any(status).(model.ValidatableStatus)
	if !ok {
		t.Error("type assertion failed")
	}
	t.Logf("Valid: %v", s.IsValid())
}
