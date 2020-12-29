package pagenation

import "fmt"

/**
	翻页sql生成器
	@author Bill
*/

const (
	_pkField   = "id"
	countField = "count"
	ListRow    = 20
)


/**
	当前sql + LIMIT 0,20 [分页sql]
 */
func QueryBuild(querySql string, currPage int, usePage bool) string {
	var sql string
	if usePage == true {
		sql = querySql + PagenationParse(currPage, ListRow)
	}
	return sql
}

/**
	汇总sql这里直接使用sql
 */
func QueryTotalBuild(tableName string, otherCondi string, alias string) string {
	totalQuery := "SELECT count(" + _pkField + ") AS " + countField + " FROM " + tableName + " "
	if alias != "" {
		totalQuery += alias + " "
	}
	if otherCondi != "" {
		totalQuery += otherCondi
	}
	return totalQuery
}

/**
页面计算(用于sql解析) 比如 1 页面 20 行 => LIMIT 0,20
比如取第五页 => 80,20
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
		pageInfo["starRow"] = PagenationStart(page, limitRow)
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


