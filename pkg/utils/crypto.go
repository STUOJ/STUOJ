package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

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
