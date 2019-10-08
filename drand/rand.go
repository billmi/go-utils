package drand

import (
	"bytes"
	"math/rand"
	"sync"
	"time"
)

/**
	随机数,随机字符串生成
    author Bill
*/

var (
	randSeek = int64(1)
	l        sync.Mutex
)

func GetRandomSring(num int, str ...string) string {
	s := "0123456789"
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

func GetRandomChars(num int) string {
	ss := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return GetRandomSring(num, ss)
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
