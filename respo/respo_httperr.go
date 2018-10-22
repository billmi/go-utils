package respo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpErr(c *gin.Context, msg string, data interface{}, url string) {
	httpResponse(c, msg, data, url, ErrorCode, http.StatusOK)
}

func HttpHErr(c *gin.Context, msg string, headers interface{}, data interface{}, url string) {
	httpHeaderResponse(c, msg, headers, data, url, ErrorCode, http.StatusOK)
}
