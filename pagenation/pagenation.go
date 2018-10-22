package pagenation

import (
	"math"
	"fmt"
)

/**
	分页页面信息
    author Bill
 */
func Paginator(page int, listRow int, total int64) map[string]interface{} {
	totalpages := int(math.Ceil(float64(total) / float64(listRow)))
	if page <= 0 {
		page = 1
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["total"] = totalpages
	paginatorMap["currpage"] = page
	paginatorMap["listRow"] = listRow
	return paginatorMap
}

/**
	页面计算(用于sql解析)
	@author Bill
 */
func PagenationParse(page int,limitRow int) string{
	pageInfo := map[string]int{
		"starRow":0,
		"limitRow":limitRow,
	}
	switch page{
	case 1:
		return fmt.Sprintf("LIMIT %d,%d",pageInfo["starRow"],pageInfo["limitRow"])
	default:
		pageInfo["starRow"] = (page - 1) * limitRow
	}
	return fmt.Sprintf("LIMIT %d,%d",pageInfo["starRow"],pageInfo["limitRow"])
}