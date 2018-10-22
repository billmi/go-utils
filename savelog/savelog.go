package savelog

import (
	"fmt"
	"demo/utils/json"
	"log"
	"os"
	"time"
)


/**
	自定义写日志
    author Bill
 */
var sufferFix = ".log"
var nowDateTime = time.Now().Format("2006-01-02 15:04:05")
var nowDate = time.Now().Format("2006-01-02")

func writeLog(logFileName string, custPrefix string, message string, logData interface{}) {
	if IsExist(nowDate) == false {
		os.Mkdir(nowDate, 0775)
	}
	var logFilePath = nowDate + "/" + logFileName + "_" + nowDate + sufferFix
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND, 0775)
	if err != nil {
		log.Panicln("error : %s ,fileName : %s", err, logFileName)
		fmt.Print("%s",err)
	}
	defer logFile.Close()
	if err != nil {
		log.Panicln("open file error")
		fmt.Print("%s",err)
	}
	perfix := fmt.Sprintf("[%s]", custPrefix)
	ErrorLog := log.New(logFile, perfix, log.Llongfile)
	var dataLog string
	if logData != "" {
		dataLog, _ = json.ToJsonString(logData)
	}
	errMsg := fmt.Sprintf("Msg : %s , Data: %s , Date : %s", message, dataLog, nowDateTime) //时间点无法修改
	ErrorLog.Println(errMsg)
}

func WriteErrorLog(logFileName string, message string, logData interface{}) {
	writeLog(logFileName, "Error", message, logData)
}

func WriteInfoLog(logFileName string, message string, logData interface{}) {
	writeLog(logFileName, "Info", message, logData)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
