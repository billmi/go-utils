package file

import (
	"os"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

/**
    文件存在检测
    @author Bill
*/

func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

/**
    创建目录
    @author Bill
*/
func BuildDir(abs_dir string) error {
	return os.MkdirAll(path.Dir(abs_dir), os.ModePerm) //生成多级目录
}

/**
    删除文件或文件夹
*/
func DeleteFile(abs_dir string) error {
	return os.RemoveAll(abs_dir)
}

/**
    获取目录所有文件夹
*/
func GetPathDirs(abs_dir string) (re []string) {
	if PathExists(abs_dir) {
		files, _ := ioutil.ReadDir(abs_dir)
		for _, f := range files {
			if f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

/**
    获取目录所有文件
*/
func GetPathFiles(abs_dir string) (re []string) {
	if PathExists(abs_dir) {
		files, _ := ioutil.ReadDir(abs_dir)
		for _, f := range files {
			if !f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}


/*
	获取程序运行路径
*/
func GetCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}
