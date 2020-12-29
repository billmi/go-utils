package xormhelper

import (
	"fmt"
	"reflect"
	"strings"
	"github.com/go-xorm/xorm"
	"strconv"
	"sort"
)

/**
	Struct XormHelper
    author : Bill
 */
type XormHelper struct {
	Datasource *xorm.Engine
}

/**
	GetDatasource
 */
func (xormhelper *XormHelper) GetDatasource() *xorm.Engine {
	return xormhelper.Datasource
}

/**
	SetDatasource
 */
func (xormhelper *XormHelper) SetDatasource(datasource *xorm.Engine) {
	xormhelper.Datasource = datasource
}

/**
	PageList to po
 */
func (xormhelper *XormHelper) GetPageLists(po interface{}, table string, fields string, pk string, alias string, join string, condition string, order string, group string, page int, listRow int) map[string]interface{} {
	var (
		_fields  = "*"
		_page    = 0
		_order   = ""
		orderCmd = " ORDER BY "
		groupBy  = " GROUP BY "
		_group   = ""
		_pk      = "`id`"
		count    = 0
		_listRow = LIST_ROWS
		querySql string
		countSql string
		resData = make(map[string]interface{}, 0)
		pageInfo = make(map[string]int,0)
	)
	if pk != "" {
		_pk = pk
	}
	if fields != "" {
		_fields = fields
	}
	var condi = " "
	if condition != "" {
		condi = condition
	}
	if order != "" {
		_order = orderCmd + order + " "
	} else {
		_order += orderCmd + _pk + " " + " DESC "
	}
	if group != "" {
		_group = groupBy + group
	}
	if page >= 0 {
		_page = page
	}
	if listRow > 0 {
		_listRow = listRow
	}
	querySql = fmt.Sprintf("SELECT %s FROM %s "+" "+alias+" "+join+" WHERE 1 ", _fields, table)
	countSql = fmt.Sprintf("SELECT count(%s) AS count FROM %s "+" "+alias+" "+join+" WHERE 1 ", _pk, table)
	querySql = querySql + condi + _group + _order
	countSql = countSql + condi + _group
	var handler = xormhelper.GetDatasource()
	countRes, _ := handler.QueryString(countSql)
	handler.SQL(QueryBuild(querySql, _page, _listRow, true)).Find(po)
	if len(countRes) > 0 {
		count, _ = strconv.Atoi(countRes[0]["count"])
	}
	if count > 0 {
		resData["list"] = po
	} else {
		resData["list"] = make([]int, 0)
	}
	pageInfo = CommaPaginator(_page, listRow, int64(count))
	resData["total_page"] = pageInfo["total_page"]
	resData["curr_page"] = pageInfo["curr_page"]
	resData["page_rows"] = pageInfo["page_rows"]
	resData["total_record"] = count
	return resData
}

/**
	List to po
 */
func (xormhelper *XormHelper) GetLists(po interface{}, table string, fields string, pk string, alias string, join string, condition string, order string, group string) map[string]interface{} {
	var (
		_fields = "*"
		_order = ""
		orderCmd = " ORDER BY "
		groupBy = " GROUP BY "
		_group = ""
		_pk = "`id`"
		querySql string
		resData = make(map[string]interface{}, 0)
	)
	if pk != "" {
		_pk = pk
	}

	var condi = " "
	if condition != "" {
		condi = condition
	}
	if order != "" {
		_order = orderCmd + order + " "
	} else {
		_order += orderCmd + _pk + " " + " DESC "
	}
	if group != "" {
		_group = groupBy + group
	}
	querySql = fmt.Sprintf("SELECT %s FROM %s "+" "+alias+" "+join+" WHERE 1 ", _fields, table)
	querySql = querySql + condi + _group + _order
	var handler = xormhelper.GetDatasource()
	handler.SQL(querySql).Find(po)
	resData["list"] = po
	return resData
}

/**
	Test Demo
	var data = [][]string{
		{"INNER","b b","b.id = a.id"},
		{"INNER","c c","c.id = b.id"},
	}
	var base = &DaoBase{}
	fmt.Print(base.ConditionJoin(data))

	Join Builder
	author : Bill
 */
func (xormhelper *XormHelper) ConditionJoin(join [][]string) string {
	var _join = ""
	for _, row := range join {
		var in = strings.ToUpper(strings.TrimSpace(row[0]))
		var table = row[1]
		var condi = row[2]
		switch in {
		case _innerFlag:
			_join += fmt.Sprintf(innerTpl, table, condi)
		case _leftFlag:
			_join += fmt.Sprintf(leftTpl, table, condi)
		case _rightFlag:
			_join += fmt.Sprintf(rightTpl, table, condi)
		default:
			_join += fmt.Sprintf(innerTpl, table, condi)
		}
	}
	return _join
}

/**
	SqlBuildConditon : Test Demo
	var condi = make(map[string]map[string]interface{})
	var inCodi = make(map[string]interface{})
	var like = make(map[string]interface{})
	var null = make(map[string]interface{})        // null
	var or = make(map[string]interface{},0)

	orCondi["condi"]  = fmt.Sprintf("bc.status = %d",0)
	nullCondi["bc.cv_id"] = nil
	like["name"] = "Bill"
	inCodi["type"] = 1
	inCodi["title"] = "young"
	condi["AND"] = inCodi
	condi["LIKE"] = like
	condi["OR"] = or
	condi["NULL"] = null
	fmt.Print(ConditionBuild(condi))
	Condition Builder
	author ： Bill
 */
func (xormhelper *XormHelper) ConditionBuild(condi map[string]map[string]interface{}) string {
	if len(condi) == 0 {
		return ""
	}
	var (
		tips []string
	)
	for k := range condi {
		tips = append(tips, k)
	}
	sort.Strings(tips)
	for _, tip := range tips {
		_currRel = strings.ToUpper(tip)
		for _field, _v := range condi[tip] {
			switch _currRel {
			case _andFlag:
				if reflect.TypeOf(_v).String() == "string" {
					_condi += fmt.Sprintf(str, _field, _v)
				} else {
					_condi += fmt.Sprintf(intCondi, _field, _v)
				}
			case _likeFlag:
				_condi += ` AND ( ` + _field + ` LIKE  "%` + _v.(string) + `%" )`
			case _gtFlag:
				_condi += fmt.Sprintf(_gtCondi, _field, _v)
			case _ltFlag:
				_condi += fmt.Sprintf(_ltCondi, _field, _v)
			case _InFlag:
				_condi += fmt.Sprintf(_InCondi, _field, _v)
			case _nullFlag:
				_condi += fmt.Sprintf(_NullCondi, _field)
			case _orFlag:
				_condi += fmt.Sprintf(_orCondi, _v)
			case _betweenFlag:
				_condi += fmt.Sprintf(_betweenCondi, _field, _v)
			case _expFlag:
				_condi += fmt.Sprintf(_expCondi, _field, _v)
			}

		}
	}
	return _condi
}

/**
	find by po Struct
 */
func (xormhelper *XormHelper) GetByPo(po interface{}, table string, condi string) interface{} {
	var handler = xormhelper.GetDatasource()
	handler.SQL("SELECT * FROM " + table + " WHERE 1 " + condi).Find(po)
	return po
}

/**
	Editor Rows && Return effect rows
 */
func (xormhelper *XormHelper) EditRow(table string, condi string, params map[string]interface{}) int {
	var handler = xormhelper.GetDatasource()
	effRow, err := handler.Table(table).Where(condi).Update(params)
	if err != nil {
		return 0
	}
	//Because update too fast,
	if effRow == 0 && err == nil { // Succ && err == nil , The Res is 0 set to 1,
		effRow = 1
	}
	return int(effRow)
}

/**
	Insert row  && 	Return new InsertId
 */
func (xormhelper *XormHelper) InsertRow(table string, params map[string]interface{}) (int, bool) {
	if len(params) == 0 {
		return 0, false
	}
	var (
		sep     = ","
		valBuff = ""
		_fields = make([]string, 0)
		_vals   = make([]interface{}, 0)
	)
	for key, row := range params {
		_fields = append(_fields, key)
		_vals = append(_vals, row)
	}
	for _, row := range _vals {
		var val = ""
		if strings.Contains(reflect.TypeOf(row).String(), "string") {
			val = fmt.Sprintf("'%v'", row)
		} else {
			val = fmt.Sprintf("%v", row)
		}
		valBuff += val + sep
	}
	var handler = xormhelper.GetDatasource()
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", table, strings.Join(_fields, sep), strings.Trim(valBuff, sep))
	res, err := handler.Exec(sql)
	if err != nil {
		return 0, false
	}
	nId, err := res.LastInsertId()
	if err != nil {
		return 0, false
	}
	return int(nId), true
}

//Get Trans Stat
func (xormhelper *XormHelper) StartTransaction() *xorm.Session {
	var handler = xormhelper.GetDatasource()
	return handler.NewSession()
}
