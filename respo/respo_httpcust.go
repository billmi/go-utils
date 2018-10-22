package respo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpCust(c *gin.Context, msg string, data interface{}, url string, httpCode int) {
	httpResponse(c, msg, data, url, ErrorCode, httpCode)
}

func HttpHCust(c *gin.Context, msg string, headers interface{}, data interface{}, url string) {
	httpHeaderResponse(c, msg, headers, data, url, ErrorCode, http.StatusOK)
}
