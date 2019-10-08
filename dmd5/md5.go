package dmd5

import (
	"crypto/md5"
	"encoding/hex"
)

/**
  md5
  @author Bill
*/

func Md5EnCode(string string) string {
	h := md5.New()
	h.Write([]byte(string)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}
