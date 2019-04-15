package json

import (
	"github.com/gpmgo/gopm/modules/log"
	"encoding/json"
)

/**
    JSON (map转json)
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

/*
	泛型比较麻烦,单独做一个
 */
func Struct2Json(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}

/**
    JSON (json转map)
    @author Bill
*/
func StringToJson(data string) map[string]interface{} {
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(data), &jsonData)
	return jsonData
}

/**
	JSONstring (json转IntList)
    @author Bill
 */
func ToIntList(data string) []int {
	var tmp = make([]int, 0)
	json.Unmarshal([]byte(data), &tmp)
	return tmp
}
