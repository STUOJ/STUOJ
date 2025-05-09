package utils_test

import (
	"STUOJ/pkg/utils"
	"testing"
)

// BenchmarkSafeTypeAssert 对 SafeTypeAssert 函数进行基准测试。
func BenchmarkSafeTypeAssert(b *testing.B) {
	// 测试用例定义
	testCases := []struct {
		name          string
		sourceValue   interface{}
		targetPointer interface{}
	}{
		{
			name:          "string to []int",
			sourceValue:   "1,2,3,4,5",
			targetPointer: new([]int),
		},
		{
			name:          "string to []int8",
			sourceValue:   "1,2,3",
			targetPointer: new([]int8),
		},
		{
			name:          "string to []uint",
			sourceValue:   "10,20,30",
			targetPointer: new([]uint),
		},
		{
			name:          "string to []float32",
			sourceValue:   "1.1,2.2,3.3",
			targetPointer: new([]float32),
		},
		{
			name:          "empty string to []int",
			sourceValue:   "",
			targetPointer: new([]int),
		},
		{
			name:          "string with spaces to []int",
			sourceValue:   " 1,  2 ,3 ",
			targetPointer: new([]int),
		},
		{
			name:          "invalid string to []int",
			sourceValue:   "1,abc,3",
			targetPointer: new([]int),
		},
		{
			name:          "int to bool (non-zero)",
			sourceValue:   1,
			targetPointer: new(bool),
		},
		{
			name:          "int to bool (zero)",
			sourceValue:   0,
			targetPointer: new(bool),
		},
		{
			name:          "uint to bool (non-zero)",
			sourceValue:   uint(1),
			targetPointer: new(bool),
		},
		{
			name:          "float64 to bool (non-zero)",
			sourceValue:   1.23,
			targetPointer: new(bool),
		},
		{
			name:          "direct assign int to int",
			sourceValue:   123,
			targetPointer: new(int),
		},
		{
			name:          "convertible assign int to int64",
			sourceValue:   123,
			targetPointer: new(int64),
		},
		{
			name:          "nil to *int (should be false, set to zero)",
			sourceValue:   nil,
			targetPointer: new(int),
		},
		{
			name:          "nil to *[]int (should be true, set to nil slice)",
			sourceValue:   nil,
			targetPointer: new([]int),
		},
		{
			name:          "int to string (not assignable/convertible)",
			sourceValue:   123,
			targetPointer: new(string),
		},
		{
			name:          "string to int (not a slice, direct assignable if types match)",
			sourceValue:   "hello",
			targetPointer: new(string),
		},
	}

	var globalDummy interface{}
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				utils.SafeTypeAssert(tc.sourceValue, tc.targetPointer)
				globalDummy = tc.targetPointer
			}
		})
	}
	// 确保全局变量不会被编译器优化掉
	_ = globalDummy
}
