package main

import (
	"fmt"
	"os"
)

//目录打开
//与文件的打开类似   OpenFile(name string,flag int,perm FileMode)(*File, error)    perm通常为os.ModeDir
func openCatalog() {
	catalog, err := os.OpenFile("G://GolangFileText", os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("打开目录出现错误：", err)
	}
	files, err := catalog.Readdir(-1) //<=0，返回All，>=0，返回指定数目

	for _, item := range files {
		fmt.Println(item.Name(), item.IsDir()) //IsDir()判断是不是目录
	}

}

func main() {
	openCatalog()
}
