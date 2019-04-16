package pagenation

import (
	"math"
	"fmt"
)

const LIST_ROWS = 15

/**
	分页页面信息
    author Bill
 */
func CommaPaginator(page int, listRow int, total int64) map[string]int {
	totalpages := int(math.Ceil(float64(total) / float64(listRow)))
	if page <= 0 {
		page = 1
	}
	paginator := make(map[string]int)
	paginator["total_page"] = totalpages
	paginator["curr_page"] = page
	paginator["page_rows"] = listRow
	return paginator
}

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
	paginatorMap["totalPage"] = totalpages
	paginatorMap["currpage"] = page
	paginatorMap["listRow"] = listRow
	return paginatorMap
}

/**
	页面计算(用于sql解析)
	@author Bill
 */
func PagenationParse(page int, limitRow int) string {
	pageInfo := map[string]int{
		"starRow":  0,
		"limitRow": limitRow,
	}
	var _page = page
	if _page <= 0 {
		_page = 1
	}
	switch _page {
	case 1:
		return fmt.Sprintf("LIMIT %d,%d", pageInfo["starRow"], pageInfo["limitRow"])
	default:
		pageInfo["starRow"] = (page - 1) * limitRow
	}
	return fmt.Sprintf("LIMIT %d,%d", pageInfo["starRow"], pageInfo["limitRow"])
}

/**
	页面计算(开始于结束步长区间)
	@author Bill
 */
func PagenationStart(page int, limitRow int) int {
	var start = 0
	var _page = page
	if _page <= 0 {
		_page = 1
	}
	switch _page {
	case 1:
		start = 0
	default:
		start = (page - 1) * limitRow
	}
	return start
}
