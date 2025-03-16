package middleware

import "github.com/gin-gonic/gin"

func SecurityHeader() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ginCtx.Header("X-Frame-Options", "DENY")
		ginCtx.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		ginCtx.Header("X-XSS-Protection", "1; mode=block")
		ginCtx.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		ginCtx.Header("Referrer-Policy", "origin-when-cross-origin")
		ginCtx.Header("X-Content-Type-Options", "nosniff")
		ginCtx.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		ginCtx.Header("Server", "Go")

		ginCtx.Next()
	}
}
