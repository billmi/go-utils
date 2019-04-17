package pagenation

import (
	"fmt"
	"reflect"
	"strings"
	"github.com/go-xorm/xorm"
	"strconv"
)

type DaoBase struct {
	Datasource *xorm.Engine
}

/**
	这里使用xorm引擎
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
 */
func (DaoBase *DaoBase) GetDatasource() *xorm.Engine {
	return DaoBase.Datasource
}

/**
	这里使用xorm引擎
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
 */
func (DaoBase *DaoBase) SetDatasource(datasource *xorm.Engine) {
	DaoBase.Datasource = datasource
}

/**
	分页
	注意 ： 连接需要传入datasource
 */
func (DaoBase *DaoBase) GetPageLists(po interface{}, table string, pk string, condition string, order string, page int, listRow int) map[string]interface{} {
	var _fields = "*"
	var _page = 0
	var _order = ""
	var orderCmd = " ORDER BY "
	var _pk = "`id`"
	var count = 0
	if pk == "" {
		_pk = pk
	}
	var querySql = fmt.Sprintf("SELECT %s FROM %s WHERE 1 ", _fields, table)
	var countSql = fmt.Sprintf("SELECT count(%s) AS count FROM %s WHERE 1 ", pk, table)
	var condi = " "
	if condition != "" {
		condi = condition
	}
	if order != "" {
		_order = orderCmd + order + " "
	} else {
		_order += orderCmd + _pk + " " + " DESC "
	}

	if page >= 0 {
		_page = page
	}
	querySql = querySql + condi + _order
	countSql = countSql + condi
	var handler = DaoBase.GetDatasource()
	countRes, _ := handler.QueryString(countSql)
	handler.SQL(QueryBuild(querySql, _page, true)).Find(po)
	var resData = make(map[string]interface{}, 0)
	count, _ = strconv.Atoi(countRes[0]["count"])
	if count > 0 {
		resData["list"] = po
	} else {
		resData["list"] = make([]int, 0)
	}
	var pageInfo = CommaPaginator(_page, listRow, int64(count))
	resData["total_page"] = pageInfo["total_page"]
	resData["curr_page"] = pageInfo["curr_page"]
	resData["page_rows"] = pageInfo["page_rows"]
	resData["total_record"] = count
	return resData
}

/**
	不分页
	注意 ： 连接需要传入datasource
 */
func (DaoBase *DaoBase) GetLists(po interface{}, table string, pk string, condition string, order string) map[string]interface{} {
	var _fields = "*"
	var _order = ""
	var orderCmd = " ORDER BY "
	var _pk = "`id`"
	if pk == "" {
		_pk = pk
	}
	var querySql = fmt.Sprintf("SELECT %s FROM %s WHERE 1 ", _fields, table)
	var condi = " "
	if condition != "" {
		condi = condition
	}
	if order != "" {
		_order = orderCmd + order + " "
	} else {
		_order += orderCmd + _pk + " " + " DESC "
	}
	querySql = querySql + condi + _order
	var handler = DaoBase.GetDatasource()
	handler.SQL(querySql).Find(po)
	var resData = make(map[string]interface{}, 0)
	resData["list"] = po
	return resData
}

/**
	SqlBuildConditon : 测试用例
		var condi = make(map[string]map[string]interface{})
		var inCodi = make(map[string]interface{})
		var like = make(map[string]interface{})
		like["name"] = "Bill"
		inCodi["type"] = 1
		inCodi["title"] = "yang"
		condi["AND"] = inCodi
		condi["LIKE"] = like
		fmt.Print(ConditionBuild(condi))
	author ： Bill
 */
func (DaoBase *DaoBase) ConditionBuild(condi map[string]map[string]interface{}) string {
	if len(condi) == 0 {
		return ""
	}
	//flag部分
	var (
		_andFlag  = "AND"
		_likeFlag = "LIKE"
		_gtFlag   = "GT" //比较
		_ltFlag   = "LT" //比较
	)

	var _condi = ""
	var str = " AND ( `%s` = '%v' )"
	var intCondi = " AND ( `%s` = %v )"
	var _gtCondi = " AND ( `%s` > %v )"
	var _ltCondi = " AND ( `%s` < %v )"
	var _currRel = ""
	for _rela, v := range condi {
		_currRel = strings.ToUpper(_rela)
		for _field, _v := range v {
			fmt.Print(reflect.TypeOf(_v).String())
			switch _currRel {
			case _andFlag:
				if reflect.TypeOf(_v).String() == "string" {
					_condi += fmt.Sprintf(str, _field, _v)
				} else {
					_condi += fmt.Sprintf(intCondi, _field, _v)
				}
			case _likeFlag:
				_condi += " AND ( `" + _field + "` LIKE '%" + _v.(string) + "%' )"
			case _gtFlag:
				_condi += fmt.Sprintf(_gtCondi, _field, _v)
			case _ltFlag:
				_condi += fmt.Sprintf(_ltCondi, _field, _v)
			}
		}
	}
	return _condi
}
