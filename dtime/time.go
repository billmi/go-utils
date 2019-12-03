package dtime

import (
	"bytes"
	"math/rand"
	"sync"
	"time"
)

/**
  时间获取
  时间转换必须加入时区设置,请注意
  @author Bill
*/

var (
	randSeek = int64(1)
	l        sync.Mutex
	zone     = "CST" //时区
)

func TimeIntToDate(time_int int) string {
	var cstZone = time.FixedZone(zone, 8*3600)
	return time.Unix(int64(time_int), 0).In(cstZone).Format("2006-01-02 15:04:05")
}

func GetNowDateTime() string {
	var cstZone = time.FixedZone(zone, 8*3600)
	return time.Now().In(cstZone).Format("2006-01-02 15:04:05")
}

func GetDate() string {
	var cstZone = time.FixedZone(zone, 8*3600)
	return time.Now().In(cstZone).Format("2006-01-02")
}

//防时间间隔
func GetIntTime() int {
	var _t = int(time.Now().Unix())
	return _t
}

//暂时独立
func _getRandomSring(num int, str ...string) string {
	s := "123456789"
	if len(str) > 0 {
		s = str[0]
	}
	l := len(s)
	r := rand.New(rand.NewSource(getRandSeek()))
	var buf bytes.Buffer
	for i := 0; i < num; i++ {
		x := r.Intn(l)
		buf.WriteString(s[x : x+1])
	}
	return buf.String()
}

func getRandSeek() int64 {
	l.Lock()
	if randSeek >= 100000000 {
		randSeek = 1
	}
	randSeek++
	l.Unlock()
	return time.Now().UnixNano() + randSeek
}

//获取今天时间戳 Today => 00:00:00
func TodayTimeUnix() int {
	t := time.Now()
	tm1 := int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix())
	return tm1
}

//获取今天时间戳 Today => 23:59:59
func TodayNightUnix() int {
	tm1 := TodayTimeUnix() + 86400 - 1
	return tm1
}
