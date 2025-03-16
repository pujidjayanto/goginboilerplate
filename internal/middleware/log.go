package middleware

import (
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/config"
	"github.com/pujidjayanto/goginboilerplate/pkg/logger"
)

func LogRequest() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		start := time.Now()
		path := ginCtx.Request.URL.Path
		raw := ginCtx.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		// Process request
		ginCtx.Next()

		// Log details after request is processed
		latency := time.Since(start)
		clientIP := ginCtx.ClientIP()
		method := ginCtx.Request.Method
		statusCode := ginCtx.Writer.Status()
		requestId := requestid.Get(ginCtx)

		// log everything when in development
		// other than that log only internal server error to prevent excessive log
		if config.GetEnv() == "development" {
			logger.Info("incoming request",
				"request_id", requestId,
				"status", statusCode,
				"method", method,
				"path", path,
				"ip", clientIP,
				"latency", latency,
				"user-agent", ginCtx.Request.UserAgent(),
			)
		}

		if statusCode >= 500 {
			logger.Error("server error",
				"request_id", requestId,
				"status", statusCode,
				"method", ginCtx.Request.Method,
				"path", path,
				"duration", time.Since(start),
				"ip", ginCtx.ClientIP(),
				"errors", ginCtx.Errors.String(),
			)
		}
	}
}
