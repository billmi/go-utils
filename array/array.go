package array

import (
	"strconv"
	"strings"
)

//合并数组
func MergeArray(dest []interface{}, src []interface{}) (result []interface{}) {
	result = make([]interface{}, len(dest)+len(src))
	copy(result, dest)
	copy(result[len(dest):], src)
	return
}

//删除数组
func DeleteArray(src []interface{}, index int) (result []interface{}) {
	result = append(src[:index], src[(index+1):]...)
	return
}

// []string => []int
func ArrayStr2Int(data []string) []int {
	var (
		arr = make([]int, 0)
	)
	if len(data) == 0 {
		return arr
	}
	for i, _ := range data {
		var num, _ = strconv.Atoi(data[i])
		arr = append(arr, num)
	}
	return arr
}

// []int => []string
func ArrayInt2Str(data []int) []string {
	var (
		arr = make([]string, 0)
	)
	if len(data) == 0 {
		return arr
	}
	for i, _ := range data {
		arr = append(arr, strconv.Itoa(data[i]))
	}
	return arr
}

// str in string list
func StrInArray(str string, data []string) bool {
	if len(data) > 0 {
		for _, row := range data {
			if str == strings.TrimSpace(row) {
				return true
			}
		}
	}
	return false
}

// str in int list
func IntInArray(num int, data []int) bool {
	if len(data) > 0 {
		for _, row := range data {
			if num == row {
				return true
			}
		}
	}
	return false
}
