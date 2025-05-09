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

// SafeTypeAssert 尝试安全地将sourceValue赋值给targetPointer指向的变量
// 处理直接赋值和Go reflect包支持的类型转换（例如基础类型、数值类型等）
// 参数:
//
//	sourceValue: 要赋值的值，可以为nil
//	targetPointer: 指向目标变量的指针，必须是一个非nil且可设置的可寻址变量
//
// 返回值:
//
//	bool: 如果赋值或转换成功返回true，否则返回false
//	      如果返回false且targetPointer是一个有效的可设置指针
//	      则将其指向的变量设置为零值
//	      如果targetPointer不是有效指针或指向不可设置的变量，则立即返回false
func SafeTypeAssert(sourceValue interface{}, targetPointer interface{}) bool {
	targetVal := reflect.ValueOf(targetPointer)

	// 检查targetPointer是否为指针且不为nil
	if targetVal.Kind() != reflect.Ptr || targetVal.IsNil() {
		return false
	}

	targetElem := targetVal.Elem()

	// 检查指向的元素是否可设置
	if !targetElem.CanSet() {
		return false
	}

	targetType := targetElem.Type()

	// 处理sourceValue为nil的情况
	if sourceValue == nil {
		// 检查目标类型是否可以接受nil值
		switch targetType.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			targetElem.Set(reflect.Zero(targetType)) // 对于可nil类型设置为nil
			return true
		default:
			// 目标类型不可nil(如int, struct)，无法赋值为nil
			targetElem.Set(reflect.Zero(targetType))
			return false
		}
	}

	sourceVal := reflect.ValueOf(sourceValue)
	sourceType := sourceVal.Type()

	// 特殊处理：数字类型到布尔类型的转换
	if targetType.Kind() == reflect.Bool {
		switch sourceType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			targetElem.SetBool(sourceVal.Int() != 0)
			return true
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			targetElem.SetBool(sourceVal.Uint() != 0)
			return true
		case reflect.Float32, reflect.Float64:
			targetElem.SetBool(sourceVal.Float() != 0)
			return true
		}
	}

	// 1. 直接赋值
	if sourceType.AssignableTo(targetType) {
		targetElem.Set(sourceVal)
		return true
	}

	// 2. 可转换赋值(处理数值类型、底层类型等)
	if sourceType.ConvertibleTo(targetType) {
		convertedValue := sourceVal.Convert(targetType)
		targetElem.Set(convertedValue)
		return true
	}

	// 如果无法赋值或转换，将目标设置为零值并返回false
	targetElem.Set(reflect.Zero(targetType))
	return false
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
