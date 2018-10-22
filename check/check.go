package check

import "regexp"

/**
    手机号码检测
    @author Bill
*/
func CheckIsMobile(mobileNum string) bool {
	var  regular = "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}