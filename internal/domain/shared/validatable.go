package shared

type ValidatableStatus interface {
	IsValid() bool
	ValidValues() []int // 返回允许的值列表
}
