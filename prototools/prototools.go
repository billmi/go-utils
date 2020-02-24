package prototools

import (
	"reflect"
	"strings"
)

var (
	UnUse_Field = "-"
)

/**
Use by proto3
ProtoBuf => params
By Bill
*/
func Proto_To_Map(inStructPtr interface{}) map[string]interface{} {
	var resData = make(map[string]interface{}, 0)
	rType := reflect.TypeOf(inStructPtr)
	rVal := reflect.ValueOf(inStructPtr)
	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		return resData
	}
	for i := 0; i < rType.NumField(); i++ {
		if rType.Field(i).Tag.Get("json") != UnUse_Field {
			resData[strings.Split(rType.Field(i).Tag.Get("json"), ",")[0]] = rVal.Field(i).Interface()
		}
	}
	for i, row := range resData {
		kid := reflect.TypeOf(row).Kind()
		switch kid {
		case reflect.String:
			resData[i] = strings.TrimSpace(row.(string))
			break
		case reflect.Int64:
			resData[i] = int(row.(int64))
			break
		case reflect.Int32:
			resData[i] = int(row.(int32))
			break
		case reflect.Float32:
			resData[i] = float32(row.(float32))
			break
		case reflect.Float64:
			resData[i] = float64(row.(float64))
			break
		}
	}
	return resData
}
