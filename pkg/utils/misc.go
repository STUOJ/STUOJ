package utils

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// StringToInt64Slice 将逗号分隔的字符串转换为int64数组
func StringToInt64Slice(str string) ([]int64, error) {
	if str == "" {
		return []int64{}, nil
	}
	strArr := strings.Split(str, ",")
	result := make([]int64, 0, len(strArr))
	for _, s := range strArr {
		num, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value %s to int64: %w", s, err)
		}
		result = append(result, num)
	}
	return result, nil
}

// StringToUint64Slice 将逗号分隔的字符串转换为int64数组
func StringToUint64Slice(str string) ([]uint64, error) {
	if str == "" {
		return []uint64{}, nil
	}
	strArr := strings.Split(str, ",")
	result := make([]uint64, 0, len(strArr))
	for _, s := range strArr {
		num, err := strconv.ParseUint(strings.TrimSpace(s), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value %s to uint64: %w", s, err)
		}
		result = append(result, num)
	}
	return result, nil
}

func StringToInt8Slice(str string) ([]int8, error) {
	if str == "" {
		return []int8{}, nil
	}
	strArr := strings.Split(str, ",")
	result := make([]int8, 0, len(strArr))
	for _, s := range strArr {
		num, err := strconv.ParseInt(strings.TrimSpace(s), 10, 8)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value %s to int8: %w", s, err)
		}
		result = append(result, int8(num))
	}
	return result, nil
}

// StringToUint8Slice 将逗号分隔的字符串转换为int64数组
func StringToUint8Slice(str string) ([]uint8, error) {
	if str == "" {
		return []uint8{}, nil
	}
	strArr := strings.Split(str, ",")
	result := make([]uint8, 0, len(strArr))
	for _, s := range strArr {
		num, err := strconv.ParseUint(strings.TrimSpace(s), 10, 8)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value %s to uint8: %w", s, err)
		}
		result = append(result, uint8(num))
	}
	return result, nil
}

func Int64ToUint64Slice(src []int64) []uint64 {
	dst := make([]uint64, len(src))
	for i, v := range src {
		dst[i] = uint64(v)
	}
	return dst
}

func Uint64SliceToString(ids []uint64) string {
	strs := make([]string, len(ids))
	for i, id := range ids {
		strs[i] = strconv.FormatUint(id, 10)
	}
	return strings.Join(strs, ",")
}

// SafeTypeAssert 安全类型断言函数，如果转换失败则返回期望类型的零值
// value: 要转换的interface{}值
// expectedType: 期望的类型，如"int", "string"等
// 返回值: 转换后的值或期望类型的零值
func SafeTypeAssert(value interface{}, expectedType string) interface{} {
	switch expectedType {
	case "int":
		if v, ok := value.(int); ok {
			return v
		}
		return 0
	case "int8":
		if v, ok := value.(int8); ok {
			return v
		}
		return int8(0)
	case "int16":
		if v, ok := value.(int16); ok {
			return v
		}
		return int16(0)
	case "int32":
		if v, ok := value.(int32); ok {
			return v
		}
		return int32(0)
	case "int64":
		if v, ok := value.(int64); ok {
			return v
		}
		return int64(0)
	case "uint":
		if v, ok := value.(uint); ok {
			return v
		}
		return uint(0)
	case "uint8":
		if v, ok := value.(uint8); ok {
			return v
		}
		return uint8(0)
	case "uint16":
		if v, ok := value.(uint16); ok {
			return v
		}
		return uint16(0)
	case "uint32":
		if v, ok := value.(uint32); ok {
			return v
		}
		return uint32(0)
	case "uint64":
		if v, ok := value.(uint64); ok {
			return v
		}
		return uint64(0)
	case "float32":
		if v, ok := value.(float32); ok {
			return v
		}
		return float32(0)
	case "float64":
		if v, ok := value.(float64); ok {
			return v
		}
		return float64(0)
	case "bool":
		if v, ok := value.(bool); ok {
			return v
		}
		return false
	case "string":
		if v, ok := value.(string); ok {
			return v
		}
		return ""
	default:
		return nil
	}
}

func ConvertStringToType[T any](str string, result *interface{}) error {
	var tmp T
	switch any(tmp).(type) {
	case int:
		parsed, err := strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case int8:
		parsed, err := strconv.ParseInt(str, 10, 8)
		if err != nil || parsed < math.MinInt8 || parsed > math.MaxInt8 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case entity.BlogStatus:
		parsed, err := strconv.ParseUint(str, 10, 8)
		if err != nil || parsed > math.MaxUint8 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case int16:
		parsed, err := strconv.ParseInt(str, 10, 16)
		if err != nil || parsed < math.MinInt16 || parsed > math.MaxInt16 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case int32:
		parsed, err := strconv.ParseInt(str, 10, 32)
		if err != nil || parsed < math.MinInt32 || parsed > math.MaxInt32 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case int64:
		parsed, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case uint:
		parsed, err := strconv.ParseUint(str, 10, 0)
		if err != nil || parsed > math.MaxUint {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case uint8:
		parsed, err := strconv.ParseUint(str, 10, 8)
		if err != nil || parsed > math.MaxUint8 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case uint16:
		parsed, err := strconv.ParseUint(str, 10, 16)
		if err != nil || parsed > math.MaxUint16 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case uint32:
		parsed, err := strconv.ParseUint(str, 10, 32)
		if err != nil || parsed > math.MaxUint32 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case uint64:
		parsed, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case float32:
		parsed, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case float64:
		parsed, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case bool:
		parsed, err := strconv.ParseBool(str)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case string:
		*result = str // 直接赋值字符串
	default:
		return fmt.Errorf("unsupported type: %v", reflect.TypeOf(tmp))
	}
	return nil
}
