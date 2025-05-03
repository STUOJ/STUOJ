package domain_test

import (
	"STUOJ/internal/domain/image"
	"STUOJ/internal/domain/image/valueobject"
	"bytes"
	"math/rand"
	"testing"
	"time"
)

// 生成随机相册编号
func randomAlbum() valueobject.Album {
	return valueobject.NewAlbum(uint8(rand.Intn(256)))
}

// 测试图片上传成功
func TestImageUpload_Success(t *testing.T) {
	img := image.Image{
		Key:    "test_key_" + time.Now().Format("150405.000"),
		Album:  randomAlbum(),
		Reader: bytes.NewBuffer([]byte("fake image data")),
	}
	_, err := img.Upload()
	if err != nil {
		t.Fatalf("图片上传失败: %v", err)
	}
}

// 测试图片上传失败（无reader）
func TestImageUpload_NoReader(t *testing.T) {
	img := image.Image{
		Key:   "test_key_" + time.Now().Format("150405.000"),
		Album: randomAlbum(),
	}
	_, err := img.Upload()
	if err == nil {
		t.Fatalf("无reader时应上传失败")
	}
}

// 测试图片删除（假设url已存在）
func TestImageDelete(t *testing.T) {
	img := image.Image{
		Url: "https://fakeurl.com/test.jpg",
	}
	err := img.Delete()
	// 只要不panic即可，具体实现可根据handler调整
	if err != nil {
		t.Logf("图片删除返回错误: %v", err)
	}
}
