package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/config"
	"github.com/pujidjayanto/goginboilerplate/pkg/logger"
)

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		// Process request
		c.Next()

		// Log details after request is processed
		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		// log everything when in development
		// other than that log only internal server error to prevent excessive log
		if config.GetEnv() == "development" {
			logger.Info("incoming request",
				"status", statusCode,
				"method", method,
				"path", path,
				"ip", clientIP,
				"latency", latency,
				"user-agent", c.Request.UserAgent(),
			)
		}

		if statusCode >= 500 {
			logger.Error("server error",
				"status", statusCode,
				"method", c.Request.Method,
				"path", path,
				"duration", time.Since(start),
				"ip", c.ClientIP(),
				"errors", c.Errors.String(),
			)
		}
	}
}
