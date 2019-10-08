package hbase64

import "encoding/base64"

/**
  base64 decode
  @author Bill
*/

func DeBase64ToString(data string) (string, bool) {
	res, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", false
	}
	return string(res), true
}
