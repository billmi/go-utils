package hsort

import "sort"


/**
	这里主要实现sort库 sort string接口进行排序
	@author Bill
 */

//排序
func SortMapStringKey(data map[string]interface{}) {
	var (
		strs []string
	)
	for k := range data {
		strs = append(strs, k)
	}
	sort.Strings(strs)
}
