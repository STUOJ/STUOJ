package model

import (
	"STUOJ/utils"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Page struct {
	Page  uint64 `json:"page,omitempty"`
	Size  uint64 `json:"size,omitempty"`
	Total uint64 `json:"total"`
}

type Field[T any] struct {
	exist bool
	value T
}

func (f *Field[T]) Exist() bool {
	return f.exist
}

func (f *Field[T]) Value() T {
	return f.value
}

func (f *Field[T]) Set(value interface{}) error {
	if v, ok := value.(T); ok {
		f.exist = true
		f.value = v
		return nil
	}
	return fmt.Errorf("cannot set value of type %T to field of type %T", value, f.value)
}

func (f *Field[T]) Parse(c *gin.Context, name string) error {
	query := c.Query(name)
	if query == "" {
		return nil
	}

	var tmp T
	var ptr interface{} = tmp

	if err := utils.ConvertStringToType[T](query, &ptr); err != nil {
		return fmt.Errorf("failed to parse value for field %s: %w", name, err)
	}

	f.Set(ptr)

	f.exist = true
	return nil
}

type FieldList[T any] struct {
	exist bool
	value []T
}

func (f *FieldList[T]) Exist() bool {
	return f.exist
}

func (f *FieldList[T]) Value() []T {
	return f.value
}

func (f *FieldList[T]) Set(value interface{}) error {
	if v, ok := value.([]T); ok {
		f.exist = true
		f.value = v
		return nil
	}
	return fmt.Errorf("cannot set value of type %T to field of type %T", value, f.value)
}

func (f *FieldList[T]) Parse(c *gin.Context, name string) error {
	query := c.Query(name)
	if query == "" {
		return nil
	}
	splQuerys := strings.Split(query, ",")

	var tmp []T

	for _, splQuery := range splQuerys {
		var tmpT T
		var ptr interface{} = tmpT
		if err := utils.ConvertStringToType[T](splQuery, &ptr); err != nil {
			return fmt.Errorf("failed to parse value for field %s: %w", name, err)
		}
		if v, ok := ptr.(T); ok {
			tmp = append(tmp, v)
		} else {
			return fmt.Errorf("type assertion failed: expected %T, got %T", tmpT, ptr)
		}
	}
	f.Set(tmp)
	f.exist = true
	return nil
}

type ConditonWhere interface {
	Parse(c gin.Context)
	GenerateWhere() func(gorm.DB) gorm.DB
	GenerateWhereWithNoPage() func(gorm.DB) gorm.DB
}
