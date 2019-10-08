package dtime

import (
	"strings"
	"time"
)

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

func GetIntTime() int {
	return int(time.Now().Unix())
}

//获取本地时间戳
func GetUtcTime(tm time.Time) int64 {
	return tm.Unix() //- 8*60*60
}

//当前时间向上取整点
func GetHour(timestamp int64) int {
	//	formaTime := time.Format("2006-01-02 15:04:05")
	tm := time.Unix(timestamp, 0)
	return tm.Hour()
}

//获取offset天的现在时间:注意时区
func GetLastDayCurrentTime(timestamp int64, offset int) time.Time {
	tm := time.Unix(timestamp, 0)
	yesDay := tm.AddDate(0, 0, 1*offset)
	return yesDay
}

//获取给定时间的星期
func GetTimeWeek(timestamp int64) int {
	tm := time.Unix(timestamp, 0)
	weekDay := tm.Weekday().String()
	var week int = 0
	switch weekDay {
	case "Monday":
		week = 1
	case "Tuesday":
		week = 2
	case "Wednesday":
		week = 3
	case "Thursday":
		week = 4
	case "Friday":
		week = 5
	case "Saturday":
		week = 6
	default:
		week = 0
	}
	return week
}

//获取向上整时时间
func GetHour0(timestamp int64) time.Time {
	tm := time.Unix(timestamp, 0)
	tStr := tm.Format("2006-01-02 15") + ":00:00"
	return StrToTime(tStr, "2006-01-02 15:04:05", nil)
}

//获取给定日期的零点时间
func GetDay0(timestamp int64) time.Time {
	tm := time.Unix(timestamp, 0)
	tStr := tm.Format("2006-01-02") + " 00:00:00"

	return StrToTime(tStr, "2006-01-02 15:04:05", nil)
}

//获取offset 0点时间
func GetUtcDay0(now time.Time, timeZone *time.Location) int64 {
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

//字符串转时间
func StrToTime(tStr, format string, timeZone *time.Location) time.Time {
	if len(format) == 0 {
		format = "2006-01-02 15:04:05"
	}
	if timeZone == nil {
		//chinaLocal, _ := time.LoadLocation("Local")
		timeZone = time.Local
	}
	ti, _ := time.ParseInLocation(format, tStr, timeZone)
	return ti
}

/*
	给定字符串时间转换成本地时间戳
*/
func StringTimetoUnix(timestr string) int64 {
	return StrToTime(timestr, "2006-01-02 15:04:05", time.Local).Unix()
}

//获取最近上个星期天的零点日期
func GetWeek0(timestamp int64) time.Time {
	weekday := GetTimeWeek(timestamp)
	tm0 := GetDay0(timestamp)
	return tm0.AddDate(0, 0, -1*weekday)
}

//获取最近上个星期天的零点日期
func GetUtcWeek0(timestamp int64) int64 {
	weekday := GetTimeWeek(timestamp)
	tm0 := GetDay0(timestamp)
	tm0 = tm0.AddDate(0, 0, -1*weekday)

	return tm0.Unix()
}

/*
	获取给定时间的当月1号零点时间
*/
func GetMonth0(timestamp int64) time.Time {
	tm0 := GetDay0(timestamp)
	month0 := tm0.Day() - 1
	tm0 = tm0.AddDate(0, 0, -1*month0) //这个月1号
	return tm0
}

//整点执行操作
func TimerByHour(f func()) {
	for {
		now := time.Now()
		// 计算下一个整点
		next := now.Add(time.Hour * 1)
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		//以下为定时执行的操作
		f()
	}
}

//获取系统时间的格式
func GetSysTimeLayout() string {
	t := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	strLayout := strings.Replace(t.String(), "+0000 UTC", "", -1)
	return strings.TrimSpace(strLayout)
}

func GetTimeStr(tm int) string {
	return time.Unix(int64(tm), 0).Format("2006-01-02 15:04:05")
}

func GetDayStr(tm int) string {
	return time.Unix(int64(tm), 0).Format("2006-01-02")
}
