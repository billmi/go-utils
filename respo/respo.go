package respo

import (
	"encoding/json"
	"time"
)

const (
	ErrorCode           = 0
	SuccessCode         = 1
	DefaultHttpErrorMsg = "请求失败!"
	DefaultHttpSuccMsg  = "请求成功!"
)

type responseEmpty struct{}

/**
此包按需求修改使用
@author Bill
*/
func ResToJsonString(msg string, data interface{}, errorCode int) string {
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
	if resultData["data"] == nil {
		resultData["data"] = responseEmpty{}
	}
	resultData["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	res := _toJsonString(resultData)
	return res
}

func _toJsonString(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
