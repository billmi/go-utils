package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/billmi/xorm-helper"
	"fmt"
	"time"
	"log"
)


/**
	基于xorm
	增删改查操作demo
 */
func main() {

	engine, err := xorm.NewEngine("mysql",
		"root"+":"+"root"+"@tcp("+"127.0.0.1"+":"+"3306"+")/"+"bill"+"?charset=utf8",
	)
	if err != nil {
		fmt.Print(err.Error())
	}
	var XormHelper = xormhelper.XormHelper{}

	//Set xorm engine
	XormHelper.SetDatasource(engine)

	var (
		tableName  = "test"
		currTime   = int(time.Now().Unix())
		inSertData = map[string]interface{}{
			"device_type":          1,
			"device_type_name":     "test device",
			"device_type_title":    "test device",
			"r_id":                 1,
			"device_screen_width":  300,
			"device_screen_height": 500,
			"create_time":          currTime,
		}
		editData = map[string]interface{}{
			"device_type_name": "update device",
			"update_time":      currTime,
		}
	)
	//insert
	XormHelper.InsertRow(tableName, inSertData)

	//Edit
	condi := fmt.Sprintf("id = %d ", 1)
	XormHelper.EditRow(tableName, condi, editData)

	/*
		Search base on XormEngine
		The demo : Only Demo
		You can create the model by reverse.
	*/
	var _model = XXModel{}
	has, err := engine.Where(condi).Get(&_model)
	if err != nil {
		log.Fatal(err.Error())
	}
	if has {
		fmt.Print(_model)
	} else {
		fmt.Print("Not Find !")
	}
}
