package respo

import (
	"github.com/gin-gonic/gin"
	"time"
	"demo/modules/common/helper"
	"strconv"
)



/**
	http响应返回
    author Bill
 */

const (
	ErrorCode           = 0
	SuccessCode         = 1
	DefaultHttpErrorMsg = "请求失败!"
	DefaultHttpSuccMsg  = "请求成功!"
)

var (
	token     = "token"
	memberId  = "member_id"
)

type responseEmpty struct{}

type Headers struct {
	MemberId int `json:"member_id"`
	Token    string `json:"token"`
}

/**
	依赖gin包
	@author Bill
 */
func httpResponse(c *gin.Context, msg string, data interface{}, url string, errorCode int, code int) {
	message := msg
	if message == "" {
		if errorCode == ErrorCode {
			message = DefaultHttpErrorMsg
		} else {
			message = DefaultHttpSuccMsg
		}
	}

	resultData := map[string]interface{}{}
	resultData["code"] = errorCode
	resultData["msg"] = message
	resultData["data"] = data
	resultData["url"] = url
	if resultData["data"] == nil {
		resultData["data"] = responseEmpty{}
	}
	resultData["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	c.JSON(code, resultData)
}

func httpHeaderResponse(c *gin.Context, msg string, headers interface{},data interface{}, url string, errorCode int, code int) {
	message := msg
	if message == "" {
		if errorCode == ErrorCode {
			message = DefaultHttpErrorMsg
		} else {
			message = DefaultHttpSuccMsg
		}
	}

	resultData := map[string]interface{}{}
	resultData["code"] = errorCode
	resultData["msg"] = message
	resultData["headers"] = headers
	resultData["data"] = data
	resultData["url"] = url
	if resultData["headers"] == nil {
		resultData["headers"] = getHeaders(c)
	}
	if resultData["data"] == nil {
		resultData["data"] = responseEmpty{}
	}
	resultData["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	c.JSON(code, resultData)
}


func getHeaders(c *gin.Context) interface{}{
	requestUser := helper.GetRequestHeader(c)
	if len(requestUser) > 0{
		token,_  := requestUser[token].(string)
		transUserId,_ := strconv.Atoi(requestUser[memberId].(string))
		return &Headers{MemberId:transUserId,Token:token}
	}else{
		return responseEmpty{}
	}
}
