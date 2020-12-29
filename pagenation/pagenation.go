package pagenation

import (
	"math"
)

const (
	TOTAL_PAGE_FIELD   = "total_page"
	PAGE_FIELD         = "page"
	ROWS_FIELD         = "rows"
	TOTAL_RECORD_FIELD = "total_record"
)

/**
	page 当前页
	listRow 每页行数
    total   数据总数

	分页数据填充 返回
		=> map[string]int
			["total_page"] => 1,
			["page"] => 1,
			["rows"] => 20,
			["total_record"] => 3,
    author Bill
*/
func CommaPaginator(page int, listRow int, total int) map[string]int {
	totalpages := int(math.Ceil(float64(total) / float64(listRow)))
	if page <= 0 {
		page = 1
	}
	paginator := make(map[string]int)
	paginator[TOTAL_PAGE_FIELD] = totalpages
	paginator[PAGE_FIELD] = page
	paginator[ROWS_FIELD] = listRow
	paginator[TOTAL_RECORD_FIELD] = total
	return paginator
}


