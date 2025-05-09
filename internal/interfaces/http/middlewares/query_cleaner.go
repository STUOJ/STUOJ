package middlewares

import (
	"github.com/gin-gonic/gin"
)

// QueryCleaner 清理URL查询参数中的空值
// 当键存在但值为空(包括空字符串)时，移除该查询参数
func QueryCleaner() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()
		for key := range query {
			if query.Get(key) == "" {
				query.Del(key)
			}
		}
		c.Request.URL.RawQuery = query.Encode()
		c.Next()
	}
}
