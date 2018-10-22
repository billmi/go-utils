package json

import (
	"github.com/gpmgo/gopm/modules/log"
	"encoding/json"
)

/**
    JSON
    @author Bill
*/

func ToJsonString(data interface{}) (string,error){
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Error("errro : %s",err)
		return "",err
	}
	return string(jsonData),err
}

func StringToJson(data string) map[string]interface{}  {
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(data),&jsonData)
	return jsonData
}
