package bootstrap

import (
	"STUOJ/internal/infrastructure/persistence/repository"
	"log"
	"time"
)

// InitDatabase initializes the database connection with retry mechanism.
// It attempts to connect to the database three times with increasing delays (5s, 10s, 15s) before panicking.
func InitDatabase() {
	var err error
	retryDelays := []time.Duration{5 * time.Second, 10 * time.Second, 15 * time.Second}

	for i := 0; i <= len(retryDelays); i++ {
		err = repository.InitDatabase()
		if err == nil {
			log.Println("初始化数据库成功")
			return
		}

		log.Printf("初始化数据库失败，尝试第 %d 次: %v", i+1, err)

		if i < len(retryDelays) {
			log.Printf("等待 %s 后重试...", retryDelays[i])
			time.Sleep(retryDelays[i])
		} else {
			log.Println("数据库连接尝试已达最大次数，程序终止。")
			panic(err)
		}
	}
}
