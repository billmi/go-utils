package pagenation

/**
	翻页与汇总SQL生成器
	@author Bill
 */

const(
	_pkField = "id"
	countField = "count"
	ListRow = 15
)



func QueryBuild(querySql string,currPage int,usePage bool) string{
	var sql string
	if usePage == true{
		sql = querySql + PagenationParse(currPage,ListRow)
	}
	return sql
}

func QueryTotalBuild(tableName string,otherCondi string,alias string) string {
	totalQuery := "SELECT count(" + _pkField + ") AS " + countField + " FROM " + tableName + " "
	if alias != ""{
		totalQuery += alias + " "
	}
	if otherCondi != ""{
		totalQuery += otherCondi
	}
	return totalQuery
}