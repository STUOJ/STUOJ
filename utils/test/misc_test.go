package utils_test

import (
	"STUOJ/utils"
	"math/rand/v2"
	"encoding/base64"
	"testing"
)

func TestRandKey(t *testing.T) {
	for range 10 {
		l := rand.IntN(256)
		str := utils.GetRandKey(l)
		if len(str) != l {
			t.Errorf("length of string is not equal to length of key")
		} else {
			t.Logf("%s", str)
		}
func TestAESEncryptDecrypt(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		name      string
		plaintext string
		token     string
	}{
		{"normal case", "hello world", "secret-token"},
		{"empty string", "", "token"},
		{"special chars", "!@#$%^&*()", "special-token"},
		{"chinese text", "你好世界", "中文密钥"},
		{"invalid token", "test", "wrong-token"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 加密测试
			encrypted, err := utils.AESEncrypt(tc.plaintext, tc.token)
			if err != nil {
				t.Fatalf("AESEncrypt failed: %v", err)
			}

			// 解密测试（使用正确token）
			decrypted, err := utils.AESDecrypt(encrypted, tc.token)
			if err != nil {
				t.Fatalf("AESDecrypt failed: %v", err)
			}

			if decrypted != tc.plaintext {
				t.Errorf("Expected %q, got %q", tc.plaintext, decrypted)
			} else {
				t.Logf("Expected %q, got %q", tc.plaintext, decrypted)
			}

			// 错误token测试
			_, err = utils.AESDecrypt(encrypted, "wrong-token")
			if err == nil {
				t.Logf("Expected error with wrong token")
			} else {
				t.Errorf("Error token decrypted successfully")
			}
		})
	}
}

func TestAESDecrypt_InvalidInput(t *testing.T) {
	// 无效base64测试
	_, err := utils.AESDecrypt("invalid-base64", "token")
	if err == nil {
		t.Error("Expected error for invalid base64")
	}

	// 短密文测试（小于IV长度）
	shortCipher := base64.StdEncoding.EncodeToString([]byte("short"))
	_, err = utils.AESDecrypt(shortCipher, "token")
	if err == nil {
		t.Error("Expected error for short ciphertext")
	}
}
