package utils

import (
	"STUOJ/internal/db/entity"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"crypto/rand"
	randv2 "math/rand/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func GetRandKey(n int) string {
	key := make([]rune, n)
	for i := range key {
		key[i] = letters[randv2.IntN(len(letters))]
	}
	return string(key)
}

func IsFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // 文件不存在，返回false和不为nil的error
		}
		return false, err // 其他错误，返回false和错误
	}
	return true, nil // 文件存在，返回true和nil的error
}

func Uint64SliceToString(ids []uint64) string {
	strs := make([]string, len(ids))
	for i, id := range ids {
		strs[i] = strconv.FormatUint(id, 10)
	}
	return strings.Join(strs, ",")
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

func AESEncrypt(str, token string) (string, error) {
	// 生成32字节的SHA256哈希，取前16字节作为AES-128密钥
	hash := sha256.Sum256([]byte(token))
	key := hash[:16]

	// 创建加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher block: %w", err)
	}

	// 创建随机IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", fmt.Errorf("failed to generate IV: %w", err)
	}

	// 创建CFB加密器
	stream := cipher.NewCFBEncrypter(block, iv)

	// 加密数据
	plaintext := []byte(str)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	// 组合IV和密文并进行base64编码
	combined := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(combined), nil
}

// AESDecrypt 解密AES-128 CFB加密的base64字符串
// encryptedStr: base64编码的密文（包含IV）
// token: 用于生成解密密钥的种子字符串
func AESDecrypt(encryptedStr, token string) (string, error) {
	// 生成密钥（与加密过程相同）
	hash := sha256.Sum256([]byte(token))
	key := hash[:16]

	// 解码base64字符串
	combined, err := base64.StdEncoding.DecodeString(encryptedStr)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	// 检查最小长度（IV长度）
	if len(combined) < aes.BlockSize {
		return "", errors.New("invalid encrypted string length")
	}

	// 分离IV和密文
	iv := combined[:aes.BlockSize]
	ciphertext := combined[aes.BlockSize:]

	// 创建解密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher block: %w", err)
	}

	// 创建CFB解密器（解密器与加密器相同）
	stream := cipher.NewCFBDecrypter(block, iv)

	// 解密数据
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}

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

func ToSnakeCase(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
