package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
	"github.com/billmi/xorm-helper"
)

/**
	datetime : 2019-09-30 18:18:18
	author : Bill
 */

// fields map struct
type Po1 struct {
	DeviceType         int    `json:"device_type"`
	DeviceTypeName     string `json:"device_type_name"`
	DeviceTypeTitle    string `json:"device_type_title"`
	DeviceScreenWidth  int    `json:"device_screen_width"`
	DeviceScreenHeight int    `json:"device_screen_height"`
}

// fields map struct
type Po2 struct {
	DeviceType         int    `json:"device_type"`
	DeviceTypeName     string `json:"device_type_name"`
	DeviceTypeTitle    string `json:"device_type_title"`
	DeviceScreenWidth  int    `json:"device_screen_width"`
	DeviceScreenHeight int    `json:"device_screen_height"`
	Title              string `json:"title"` //test1 .title
}

func main() {
	engine, err := xorm.NewEngine("mysql",
		"root"+":"+"root"+"@tcp("+"127.0.0.1"+":"+"3306"+")/"+"bill"+"?charset=utf8",
	)
	if err != nil {
		fmt.Print(err.Error())
	}
	var (
		XormHelper = xormhelper.XormHelper{}

		//Use Data po struct, po <=> fields
		po = []*Po2{}

		//if not join  can be set empty string  => ""
		po1Join = [][]string{
			{"INNER", "test1 b", "b.id = a.r_id"},
		}

		//if no condition can be set empty string  => ""  ||  use condition-build.go
		condition = ""

		page = 1 //default

		listRow = 15 //default

		group     = "" // id
		order     = "a.id DESC"
		alias     = "a" // alisa a
		pk        = "a.id"
		tableName = "test"  // tablename => test

		//field Select , if you have custom ,you can set the Po by yourself
		fields = "a.*,b.title"
	)

	//Set xorm engine
	XormHelper.SetDatasource(engine)

	pageListInfo := XormHelper.GetPageLists(&po, tableName, fields, pk, alias, XormHelper.ConditionJoin(po1Join), condition, order, group, page, listRow)

	//Can Get PageInfo
	fmt.Printf("\r\n total_page : %d \r\n", pageListInfo["total_page"])
	fmt.Printf("\r\n page : %d \r\n", pageListInfo["page"])
	fmt.Printf("\r\n rows : %d \r\n", pageListInfo["rows"])
	fmt.Printf("\r\n total_record : %d \r\n", pageListInfo["total_record"])

	//Data in po
	for _, row := range po {
		fmt.Printf("\r\n relation_title : %s \r\n", row.Title)
	}
}
