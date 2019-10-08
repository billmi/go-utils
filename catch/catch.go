package catch

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"time"
)

/*
  用于抓取错误记录打印(panic信息，不可度量的)
  注意：每个 goroutine  开始部分都需要写上:defer mycatch.Dmp()
  类似于
	go func(){
		defer mycatch.Dmp()
	}()
*/

func Dmp() {
	errstr := ""
	if err := recover(); err != nil {
		errstr += (fmt.Sprintf("%v\r\n", err)) //输出panic信息
		errstr += ("--------------------------------------------\r\n")
	}

	errstr += (string(debug.Stack())) //输出堆栈信息
	//OnWriteErrToFile(errstr)
}

func OnWriteErrToFile(errstring string) {
	path := GetModelPath() + "/err"
	if !PathExists(path) {
		os.MkdirAll(path, os.ModePerm) //生成多级目录
	}

	now := time.Now()                                               //获取当前时间
	pid := os.Getpid()                                              //获取进程ID
	time_str := now.Format("2006-01-02")                            //设定时间格式
	fname := fmt.Sprintf("%s/panic_%s-%x.log", path, time_str, pid) //保存错误信息文件名:程序名-进程ID-当前时间（年月日时分秒）
	fmt.Println("panic to file ", fname)

	f, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString("=========================" + now.Format("2006-01-02 15:04:05 ========================= \r\n"))
	f.WriteString(errstring) //输出堆栈信息
	f.WriteString("=========================end=========================")
}

//获取目录地址
func GetModelPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path := filepath.Dir(file)
	path, _ = filepath.Abs(path)

	return path
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
