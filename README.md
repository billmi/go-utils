# go-utils
go 语言工具库
工具库大部分不依赖第三方包,按目录文件分包管理

### 文档

[![GitHub stars](https://img.shields.io/github/stars/billmi/go-utils)](https://github.com/billmi/go-utils/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/billmi/go-utils)](https://github.com/billmi/go-utils/network)
[![GitHub license](https://img.shields.io/github/license/billmi/go-utils)](https://github.com/billmi/go-utils/blob/master/LICENSE)

---------------------------------------------------------------------------------------------------------------------------------

|  目录包名   |                           功能简要                           |         备注         |
| :---------: | :----------------------------------------------------------: | :------------------: |
|     aes     |                       主要是aes加解密 【CBC】                |                      |
|    array    | 数组操作,如 : 合并数组,删除数组元素,<br />数组[]string => []int <=> []int => []string相互转换,<br />验证字符串或数字是否在数组内 |                      |
|  camelcase  |                   驼峰转小写 AaBb => aa_bb                   |         收集         |
|    catch    |         用于抓取错误记录打印(panic信息,不可度量的)          |      [还需完善]      |
|    check    |            实现是否为手机号码,是否是身份证号验证             |                      |
|   combine   | 笛卡尔组合<br />cart := [][]string{{"a1", "a2"},{"b1", "b2"}}<br />结果为 2 * 2 = 4个 |                      |
|   convert   | 实现一些字符转换功能,如编/解码二进制,<br />byte转16进制字符串,16进制字符串转byte等 |                      |
|   ddebug    |                     打印当前数据调试使用                     |                      |
|    djson    |                     实现json字符串互转                      |                      |
|    dmd5     |                           md5加密                            |                      |
|    drand    |                      随机数/字符串生成                       |                      |
|   dstring   |    反转字符串,字符串 <=> 字节互转,字符串是否为中文             |                      |
|    dtime    |    处理时间获取比如获取时间戳,时间日期,顺带处理时区的问题    |                      |
|    emoji    |                 Emoji表情解码,Emoji表情转换                  |                      |
|   errors    |                    错误处理,抛出打印处理                     | 使用需要参考测试文件 |
|    file     |   文件存在,目录创建,获取目录所有文件夹,获取目录所有文件等    |                      |
| floatformat |        浮点数与浮点字符串转换,尤其支持保留位填充方式         |                      |
|   hbase64   |                       baseDecodeString                       |                      |
|   hcache    |                           内存缓存                           |       收集整理       |
|  helastic   |                          es基本操作                          |       收集整理       |
|     hip     |                      返回远程客户端的ip                      |                      |
|  hreflect   |                          获取函数名                          |                      |
|    hsort    | 这里主要是实现sort string接口排序,其他类型可参照文档实现       |                      |
|    hsrsa    |                      rsa公私密钥加解密                       |                      |
|  httpsend   |                     Get,Post请求发送处理                     |                      |
| pagenation  |                           分页计算                           |                      |
| prototools  |              protobuf转换map[string]interface{}              |                      |
|    respo    |                     返回数据封装转字符串                     |                      |
|   savelog   | 写文件日志,这里可以用其他包代替,此包只是用于调试使用[不依赖第三方包] |                      |
| serializing |                    字符串序列化与反序列化                    |                      |
|    sign     |                         数据填充****                         |                      |
|    task     |                           定时任务                           | 仔细看示例[收集整理] |
|    trans    |          Struct2Map,结构体转map[string]interface{}           |                      |

###### 从工作,生活,学习中来,有时也有收集欲望,自己用也可以大家用! 如果收录了在这里找到了引用未加入您的地址也可以联系我,希望一起学习,一起进步.
