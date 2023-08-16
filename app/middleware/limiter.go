package middleware

import (
	"gin-rest-tool/pkg/app"
	"gin-rest-tool/pkg/errSys"
	"gin-rest-tool/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.Error(errSys.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
