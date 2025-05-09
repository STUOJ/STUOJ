package valueobject

import (
	"STUOJ/internal/domain/shared"
	"fmt"
)

// TeamSize 表示比赛团队大小的值对象
type TeamSize struct {
	shared.Valueobject[uint8]
}

// Verify 验证团队大小是否有效
func (t TeamSize) Verify() error {
	if t.Value() < 1 {
		return fmt.Errorf("团队大小必须大于0")
	}
	if t.Value() > 10 {
		return fmt.Errorf("团队大小不能超过10人")
	}
	return nil
}

// NewTeamSize 创建一个新的团队大小值对象
func NewTeamSize(size uint8) TeamSize {
	var t TeamSize
	t.Set(size)
	return t
}
