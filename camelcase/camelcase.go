package camelcase

import (
	"bytes"
)

/**
作者链接 : https://www.jianshu.com/p/de3782af0a3a
收集by Bill
*/

/**
大写驼峰转小写,使用 "_" 隔开
log.Println(CamelCase("AAAA"))
log.Println(CamelCase("IconUrl"))
log.Println(CamelCase("iconUrl"))
log.Println(CamelCase("parentId"))
log.Println(CamelCase("a9b9Ba"))
log.Println(CamelCase("_An"))

res : demo
	 a_a_a_a
	 icon_url
	 icon_url
	 parent_id
	 a9b9ba
	 Xan
*/
func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		t = append(t, 'X')
		i++
	}
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIIUpper(s[i+1]) {
			continue
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}

		if isASCIIUpper(c) {
			c ^= ' '
		}
		t = append(t, c)

		for i+1 < len(s) && isASCIIUpper(s[i+1]) {
			i++
			t = append(t, '_')
			t = append(t, bytes.ToLower([]byte{s[i]})[0])
		}
	}
	return string(t)
}
func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
