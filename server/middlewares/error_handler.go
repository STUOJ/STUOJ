package middlewares

import (
	"net/http"

	"STUOJ/internal/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			if e, ok := err.(*errors.Error); ok {
				c.JSON(e.HttpStatus, gin.H{
					"code":    e.Code,
					"message": e.Message,
					"data":    e.Data,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "服务器内部错误",
				})
			}
		}
	}
}
