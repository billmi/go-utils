package array

import "strconv"

//合并数组
func Copy(dest []interface{}, src []interface{}) (result []interface{}) {
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
func ArrayStr2Int(data []string)[]int{
	var (
		arr = make([]int,0)
	)
	if len(data) == 0{
		return arr
	}
	for i,_ := range data{
		var num,_ = strconv.Atoi(data[i])
		arr = append(arr, num)
	}
	return arr
}

// []int => []string
func ArrayInt2Str(data []int)[]string{
	var (
		arr = make([]string,0)
	)
	if len(data) == 0{
		return arr
	}
	for i,_ := range data{
		arr = append(arr, strconv.Itoa(data[i]))
	}
	return arr
}

