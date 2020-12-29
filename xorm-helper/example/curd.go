package main

import (
	"fmt"
	"github.com/billmi/xorm-helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go-utils/xorm-helper/example/models"
	"log"
	"time"
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
	//insert 这里会做常规类型判断转化 只支持 float32,64 || int32 ,64 || String 可以自己扩展
	XormHelper.InsertRow(tableName, inSertData)

	//Edit
	condi := fmt.Sprintf("id = %d ", 1)
	XormHelper.EditRow(tableName, condi, editData)

	/*
		Search base on XormEngine
		使用xorm生成的model结构体
		The demo : Only Demo
		You can create the model by reverse.
	*/
	var _model = models.Test{}
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
