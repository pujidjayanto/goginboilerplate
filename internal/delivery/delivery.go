package delivery

import (
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SuccessNoContent(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusNoContent, nil)
}

func SuccessCreated(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusCreated, nil)
}

func SuccessWithMetadata(ginCtx *gin.Context, data, metadata any) {
	r := Response{
		Message: SuccessMessage,
		Data:    data,
	}

	if metadata != nil {
		r.Metadata = metadata
	}

	ginCtx.JSON(http.StatusOK, r)
}

func Success(ginCtx *gin.Context, response any) {
	ginCtx.JSON(http.StatusOK, Response{
		Message: SuccessMessage,
		Data:    response,
	})
}

func Failed(ginCtx *gin.Context, statusCode int, message string) {
	r := ErrorResponse{
		Error: message,
	}

	if isAppTrace, errParse := strconv.ParseBool(ginCtx.Request.Header.Get("X-Enable-Trace")); errParse == nil && isAppTrace {
		r.Trace = string(debug.Stack())
	}

	ginCtx.Header("Connection", "close")
	ginCtx.JSON(statusCode, r)
}
