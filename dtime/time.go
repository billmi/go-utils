package dtime

import "time"

/**
    时间获取
    @author Bill
*/

func GetNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}
