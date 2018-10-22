package respo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpSucc(c *gin.Context, msg string, data interface{}, url string) {
	httpResponse(c, msg, data, url, SuccessCode, http.StatusOK)
}

func HttpHSucc(c *gin.Context, msg string, headers interface{}, data interface{}, url string) {
	httpHeaderResponse(c, msg, headers, data, url, SuccessCode, http.StatusOK)
}