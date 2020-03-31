# kgo
k`s golang helper/library/utils
golang 常用函数库/工具集,仅测试支持64位linux

### 文档

[![GitHub stars](https://img.shields.io/github/stars/billmi/go-utils)](https://github.com/billmi/go-utils/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/billmi/go-utils)](https://github.com/billmi/go-utils/network)

### 函数接收器
- *KFile* 为文件操作,如
```go
chk := KFile.IsExist(filename)
```
- *KStr* 为字符串操作,如
```go
res := KStr.Trim(" hello world ")
```
- *KNum* 为数值操作,如
```go
res := KNum.NumberFormat(123.4567890, 3, ".", "")
```
- *KArr* 为数组(切片/字典)操作,如
```go
mp := map[string]string{
    "a": "aa",
    "b": "bb",
}
chk := KArr.InArray("bb", mp)
```
- *KTime* 为时间操作,如
```go
res, err := KTime.Str2Timestamp("2019-07-11 10:11:23")
```
- *KConv* 为类型转换操作,如
```go
res := KConv.ToStr(false)
```
- *KOS* 为系统和网络操作,如
```go
res, err := KOS.LocalIP()
```
- *KEncr* 为加密操作,如
```go
res, err := KEncr.PasswordHash([]byte("123456"))
```
- *KDbug* 为调试操作,如
```go
KDbug.DumpStacks()
```





```







