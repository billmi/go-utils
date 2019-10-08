package savelog

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/**
	自定义写日志
    author Bill
*/
var sufferFix = ".log"

func WriteLog(dir string, logFileName string, logData map[string]interface{}) {
	var logFilePath = dir + "/" + logFileName + sufferFix
	var Mkdir = dir
	if IsExist(Mkdir) == false {
		os.Mkdir(Mkdir, 0775)
	}
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0775)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer logFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	Log := log.New(logFile, "", log.LstdFlags)

	Msg := fmt.Sprintf(" %s ", _toJsonString(logData))
	Log.Println(Msg)
}

/**
  JSON (map转json)
  @author Bill
*/

func _toJsonString(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
