package respo

import (
	"time"
	"fmt"
	json2 "commaai.cn/commams-oauth/extend/tools/json"
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
	res, err := ToJsonString(resultData)
	if err != nil {
		fmt.Print(err)
	}
	return res
}


/**
    JSON (map转json,包无依赖性，直接copy过来了)
    @author Bill
*/

func ToJsonString(data map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Error("errro : %s", err)
		return "", err
	}
	return string(jsonData), err
}
