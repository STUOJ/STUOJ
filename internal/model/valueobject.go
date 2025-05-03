package model

import "fmt"

type Valueobject[T any] struct {
	exist bool
	value T
}

func (v *Valueobject[T]) Exist() bool {
	return v.exist
}

func (v *Valueobject[T]) Value() T {
	return v.value
}

func (v *Valueobject[T]) Set(value T) {
	v.exist = true
	v.value = value
}

// Verify 验证值对象的值是否合法
func (v *Valueobject[T]) Verify() error {
	return nil
}

// String 返回值的字符串表示
func (v *Valueobject[T]) String() string {
	if v.Exist() {
		switch val := any(v.value).(type) {
		case string:
			return val
		default:
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}

/*
// Equals 判断两个值对象是否相等（必须为comparable的类型）
func (v *Valueobject[T]) Equals(u *Valueobject[T]) bool {
	if v.Exist() && u.Exist() {
		return v.Value() == u.Value()
	}
	return false
}
*/
