package dstring

import (
	"regexp"
	"unsafe"
)



//反转字符串
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//字节转字符串
func BytesToString(data []byte) string{
	return *(*string)(unsafe.Pointer(&data))
}

//字符串转字节数组
func StringToBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}

//判断字符串是否为中文[精确度需要反复试验]
func IsContainCN(str string) bool {
	var hzRegexp = regexp.MustCompile("[\u4e00-\u9fa5]+")
	return hzRegexp.MatchString(str)
}


