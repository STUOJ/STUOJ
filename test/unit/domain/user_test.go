package domain_test

import (
	"STUOJ/internal/infrastructure/repository"
	"STUOJ/pkg/config"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"STUOJ/internal/domain/user"
	"STUOJ/internal/domain/user/valueobject"
	"STUOJ/internal/infrastructure/repository/entity"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := config.InitConfig()
	if err != nil {
		panic("配置文件加载失败：" + err.Error())
	}
	err = repository.InitDatabase()
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	code := m.Run()

	os.Exit(code)
}

func randId() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randId := strconv.Itoa(random.Intn(1 << 16))
	return randId
}

func TestUser_Create(t *testing.T) {
	randId := randId()

	u := user.User{
		Username:  valueobject.NewUsername("test_" + randId),
		Password:  valueobject.NewPassword("password123"),
		Role:      entity.RoleUser,
		Email:     valueobject.NewEmail(randId + "@example.com"),
		Avatar:    valueobject.NewAvatar("http://example.com/avatar.png"),
		Signature: valueobject.NewSignature("Hello World"),
	}

	id, err := u.Create()
	assert.NoError(t, err)
	assert.NotZero(t, id)
}

func TestUser_Update(t *testing.T) {
	randId := randId()

	u := user.User{
		Id:         1,
		Username:   valueobject.NewUsername("test_" + randId),
		Password:   valueobject.NewPassword("newpassword123"),
		Role:       entity.RoleAdmin,
		Email:      valueobject.NewEmail(randId + "@example.com"),
		Avatar:     valueobject.NewAvatar("http://example.com/newavatar.png"),
		Signature:  valueobject.NewSignature("Updated Signature"),
		UpdateTime: time.Now(),
	}

	err := u.Update(true)
	assert.NoError(t, err)
}
