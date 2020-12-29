BaseFrom : [go-xorm](https://github.com/go-xorm/xorm)

## **Install**
`
go get -u -v github.com/billmi/xorm-helper
`

### 说明 : 

此项目初衷，不想把一些MYSQL简单操作复杂化,做为一个扩展工具使用:已简单实现联表等常用操作,一般很少用到的子查询或者复合查询等,后期会继续更新.

以人人可以快速上手使用的心态,进行项目更新~

| 简要                                            | 链接                                                         | 备注                              |
| ----------------------------------------------- | ------------------------------------------------------------ | --------------------------------- |
| join生成器                                      | [Join](https://github.com/billmi/xorm-helper/blob/master/example/join.go) | 您也可以使用xorm提供的Builder生成 |
| 实现xorm常规操作DB,使用map结构化,使用时参考demo | [Curd](https://github.com/billmi/xorm-helper/blob/master/example/curd.go) | 您也可以使用xorm提供的Builder生成 |
| condition生成器                                 | [Condition](https://github.com/billmi/xorm-helper/blob/master/example/condition-build.go) | 使用时,请查看Demo                 |
| 连表分页查询                                    | [Pagenation](https://github.com/billmi/xorm-helper/blob/master/example/pagenation-list.go) | 结果必须关联Pojo-Struct           |
| 不分页List                                      | [List-NoPagenation](https://github.com/billmi/xorm-helper/blob/master/example/list.go) | 结果必须关联Pojo-Struct           |

ps : test.sql用来测试

