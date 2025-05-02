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
