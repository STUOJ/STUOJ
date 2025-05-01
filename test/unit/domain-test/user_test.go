package domain_test

import (
	"STUOJ/internal/domain/user"
	"STUOJ/internal/domain/user/valueobject"
	"STUOJ/internal/infrastructure/repository/entity"
	"math/rand"
	"strconv"
	"testing"
)

// 生成随机用户名
func randomUsername() valueobject.Username {
	return valueobject.NewUsername("user_" + strconv.Itoa(rand.Intn(100000)))
}

// 生成随机邮箱
func randomEmail() valueobject.Email {
	return valueobject.NewEmail("user" + strconv.Itoa(rand.Intn(100000)) + "@example.com")
}

// 测试用户创建成功
func TestUserCreate_Success(t *testing.T) {
	u := &user.User{
		Username:  randomUsername(),
		Password:  valueobject.NewPassword("Password123!"),
		Role:      entity.RoleUser,
		Email:     randomEmail(),
		Avatar:    valueobject.NewAvatar("https://avatar.com/1.png"),
		Signature: valueobject.NewSignature("签名"),
	}
	id, err := u.Create()
	if err != nil {
		t.Fatalf("创建用户失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建用户返回的ID无效")
	}
}

// 测试用户名为空时创建失败
func TestUserCreate_EmptyUsername(t *testing.T) {
	u := &user.User{
		Username:  valueobject.NewUsername(""),
		Password:  valueobject.NewPassword("Password123!"),
		Role:      entity.RoleUser,
		Email:     randomEmail(),
		Avatar:    valueobject.NewAvatar("https://avatar.com/1.png"),
		Signature: valueobject.NewSignature("签名"),
	}
	_, err := u.Create()
	if err == nil {
		t.Fatalf("用户名为空时应创建失败")
	}
}

// 测试邮箱格式错误时创建失败
func TestUserCreate_InvalidEmail(t *testing.T) {
	u := &user.User{
		Username:  randomUsername(),
		Password:  valueobject.NewPassword("Password123!"),
		Role:      entity.RoleUser,
		Email:     valueobject.NewEmail("invalid-email"),
		Avatar:    valueobject.NewAvatar("https://avatar.com/1.png"),
		Signature: valueobject.NewSignature("签名"),
	}
	_, err := u.Create()
	if err == nil {
		t.Fatalf("邮箱格式错误时应创建失败")
	}
}

// 测试用户更新成功
func TestUserUpdate_Success(t *testing.T) {
	u := &user.User{
		Username:  randomUsername(),
		Password:  valueobject.NewPassword("Password123!"),
		Role:      entity.RoleUser,
		Email:     randomEmail(),
		Avatar:    valueobject.NewAvatar("https://avatar.com/1.png"),
		Signature: valueobject.NewSignature("签名"),
	}
	id, err := u.Create()
	if err != nil {
		t.Fatalf("创建用户失败: %v", err)
	}
	u.Id = id
	u.Signature = valueobject.NewSignature("新签名")
	t.Logf("更新用户: %v", u)
	err = u.Update(true)
	if err != nil {
		t.Fatalf("更新用户失败: %v", err)
	}
}

// 测试更新不存在的用户
func TestUserUpdate_NotFound(t *testing.T) {
	u := &user.User{
		Id:        99999999,
		Username:  randomUsername(),
		Password:  valueobject.NewPassword("Password123!"),
		Role:      entity.RoleUser,
		Email:     randomEmail(),
		Avatar:    valueobject.NewAvatar("https://avatar.com/1.png"),
		Signature: valueobject.NewSignature("签名"),
	}
	err := u.Update(true)
	if err == nil {
		t.Fatalf("更新不存在的用户应失败")
	}
}

// 测试用户删除成功
func TestUserDelete_Success(t *testing.T) {
	u := &user.User{
		Username:  randomUsername(),
		Password:  valueobject.NewPassword("Password123!"),
		Role:      entity.RoleUser,
		Email:     randomEmail(),
		Avatar:    valueobject.NewAvatar("https://avatar.com/1.png"),
		Signature: valueobject.NewSignature("签名"),
	}
	id, err := u.Create()
	if err != nil {
		t.Fatalf("创建用户失败: %v", err)
	}
	u.Id = id
	err = u.Delete()
	if err != nil {
		t.Fatalf("删除用户失败: %v", err)
	}
}

// 测试删除不存在的用户
func TestUserDelete_NotFound(t *testing.T) {
	u := &user.User{
		Id:        99999999,
		Username:  randomUsername(),
		Password:  valueobject.NewPassword("Password123!"),
		Role:      entity.RoleUser,
		Email:     randomEmail(),
		Avatar:    valueobject.NewAvatar("https://avatar.com/1.png"),
		Signature: valueobject.NewSignature("签名"),
	}
	err := u.Delete()
	if err == nil {
		t.Fatalf("删除不存在的用户应失败")
	}
}
