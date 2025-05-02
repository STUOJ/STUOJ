package model

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

// Verify 验证值对象的值是否合法，需要在子类中实现
func (v *Valueobject[T]) Verify() error {
	return nil
}

// String 返回值的字符串表示，需要在子类中实现
func (v *Valueobject[T]) String() string {
	return ""
}
