package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/billmi/xorm-helper"
)

/**
	datetime : 2019-08-30 18:18:18
	author : Bill
 */

func main() {

	var XormHelper = xormhelper.XormHelper{}

	//condition Build
	var (
		_condi = make(map[string]map[string]interface{}, 0)
	)

	//Like Condi Set
	_condi["LIKE"] = map[string]interface{}{
		"title": "Bill",
	}

	//And Condi Set
	_condi["AND"] = map[string]interface{}{
		"title": "Bill",
	}

	//GT Set
	_condi["GT"] = map[string]interface{}{
		"device_screen_width": 300,
	}

	//GT Set
	_condi["LT"] = map[string]interface{}{
		"device_screen_width": 300,
	}

	//IN Set
	_condi["IN"] = map[string]interface{}{
		"id": "1,2,3",
	}

	//NULL Set  (only field)
	_condi["NULL"] = map[string]interface{}{
		"device_type_title": "",
	}

	//OR Set Or is special , condition you can build by yourself
	_condi["OR"] = map[string]interface{}{
		"Or_1": "id = 1",
		"Or_2": "device_screen_width > 0",
	}

	fmt.Print("\r\n ========== Conditon Build \r\n")
	fmt.Print(XormHelper.ConditionBuild(_condi))
	fmt.Print("\r\n ========== Conditon Build \r\n")

	//join build , -- join
	var joinDemo = [][]string{
		{"LEFT", "b b", "b.id = a.id"},
		{"INNER", "c c", "c.id = b.id"},
		{"RIGHT", "c c", "c.id = b.id"},
	}
	fmt.Print("\r\n ========== Join Conditon Build \r\n")
	fmt.Print(XormHelper.ConditionJoin(joinDemo))
}
