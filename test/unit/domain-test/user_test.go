package domain_test

import (
	"STUOJ/internal/domain/user"
	"STUOJ/internal/domain/user/valueobject"
	"STUOJ/internal/infrastructure/persistence/entity"
	"math/rand"
	"strconv"
	"testing"
)

// 生成随机用户名
func randomUsername() string {
	return "user_" + strconv.Itoa(rand.Intn(100000))
}

// 生成随机邮箱
func randomEmail() string {
	return "user" + strconv.Itoa(rand.Intn(100000)) + "@example.com"
}

// 测试用户创建成功
func TestUserCreate_Success(t *testing.T) {
	u := user.NewUser(
		user.WithUsername(randomUsername()),
		user.WithPasswordPlaintext("Password123!"),
		user.WithRole(entity.RoleUser),
		user.WithEmail(randomEmail()),
		user.WithAvatar("https://avatar.com/1.png"),
		user.WithSignature("签名"),
	)
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
	u := user.NewUser(
		user.WithUsername(""),
		user.WithPasswordPlaintext("Password123!"),
		user.WithRole(entity.RoleUser),
		user.WithEmail(randomEmail()),
		user.WithAvatar("https://avatar.com/1.png"),
		user.WithSignature("签名"),
	)
	_, err := u.Create()
	if err == nil {
		t.Fatalf("用户名为空时应创建失败")
	}
}

// 测试邮箱格式错误时创建失败
func TestUserCreate_InvalidEmail(t *testing.T) {
	u := user.NewUser(
		user.WithUsername(randomUsername()),
		user.WithPasswordPlaintext("Password123!"),
		user.WithRole(entity.RoleUser),
		user.WithEmail("invalid-email"),
		user.WithAvatar("https://avatar.com/1.png"),
		user.WithSignature("签名"),
	)
	_, err := u.Create()
	if err == nil {
		t.Fatalf("邮箱格式错误时应创建失败")
	}
}

// 测试用户更新成功
func TestUserUpdate_Success(t *testing.T) {
	u := user.NewUser(
		user.WithUsername(randomUsername()),
		user.WithPasswordPlaintext("Password123!"),
		user.WithRole(entity.RoleUser),
		user.WithEmail(randomEmail()),
		user.WithAvatar("https://avatar.com/1.png"),
		user.WithSignature("签名"),
	)
	id, err := u.Create()
	if err != nil {
		t.Fatalf("创建用户失败: %v", err)
	}
	u.Id = id
	u.Signature = valueobject.NewSignature("新签名")
	err = u.Update()
	if err != nil {
		t.Fatalf("更新用户失败: %v", err)
	}
}

// 测试更新不存在的用户
func TestUserUpdate_NotFound(t *testing.T) {
	u := user.NewUser(
		user.WithId(99999999),
		user.WithUsername(randomUsername()),
		user.WithPasswordPlaintext("Password123!"),
		user.WithRole(entity.RoleUser),
		user.WithEmail(randomEmail()),
		user.WithAvatar("https://avatar.com/1.png"),
		user.WithSignature("签名"),
	)
	err := u.Update()
	if err == nil {
		t.Fatalf("更新不存在的用户应失败")
	}
}

// 测试用户删除成功
func TestUserDelete_Success(t *testing.T) {
	u := user.NewUser(
		user.WithUsername(randomUsername()),
		user.WithPasswordPlaintext("Password123!"),
		user.WithRole(entity.RoleUser),
		user.WithEmail(randomEmail()),
		user.WithAvatar("https://avatar.com/1.png"),
		user.WithSignature("签名"),
	)
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
	u := user.NewUser(
		user.WithId(99999999),
		user.WithUsername(randomUsername()),
		user.WithPasswordPlaintext("Password123!"),
		user.WithRole(entity.RoleUser),
		user.WithEmail(randomEmail()),
		user.WithAvatar("https://avatar.com/1.png"),
		user.WithSignature("签名"),
	)
	err := u.Delete()
	if err == nil {
		t.Fatalf("删除不存在的用户应失败")
	}
}

// 测试用户修改密码成功
func TestUserUpdatePassword_Success(t *testing.T) {
	u := user.NewUser(
		user.WithUsername(randomUsername()),
		user.WithPasswordPlaintext("OldPassword123!"),
		user.WithRole(entity.RoleUser),
		user.WithEmail(randomEmail()),
		user.WithAvatar("https://avatar.com/1.png"),
		user.WithSignature("签名"),
	)
	id, err := u.Create()
	if err != nil {
		t.Fatalf("创建用户失败: %v", err)
	}
	u.Id = id

	// 修改密码
	u.Password = valueobject.NewPasswordPlaintext("NewPassword456!")
	err = u.Update()
	if err != nil {
		t.Fatalf("修改密码失败: %v", err)
	}
}
