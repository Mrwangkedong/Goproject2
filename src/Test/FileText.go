package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
1.读取目录，筛选出.txt文件
2.读取txt文件，进行逐行读取，并对每一行的中的关键词进行累加
3.利用map进行存取
*/
func readCatelog(src string, filesSlice []string) []string {
	catalog, err := os.OpenFile(src, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("打开目录出现错误：", err)
	}
	files, err := catalog.Readdir(-1) //<=0，返回All，>=0，返回指定数目

	for _, item := range files {
		if strings.HasSuffix(item.Name(), ".txt") {
			filesSlice = append(filesSlice, item.Name())
		}
	}

	return filesSlice
}

func readFile(fileName string, mapData map[string]int) {
	//1.打开文件
	f, err := os.OpenFile(fileName, os.O_RDWR, 6)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	defer f.Close()
	//2.利用循环逐行读取
	//2.0 创建一个带有缓冲区的reader
	reader := bufio.NewReader(f) //2.1 创建reader
	for {
		buf, err := reader.ReadBytes('\n') //2.2 读数据  ReadBytes[split_item]  读取1-->返回49
		if err != nil && err == io.EOF {   //2.3 判断文件结尾标志(需要单独去读)--io.EOF
			fmt.Println("ERROR: ", err)
			return
		}
		//2.4 将buf按空格进行分割
		words := strings.Split(string(buf), " ")
		//2.5将words数组进行遍历，加入在map中就加1，假如不在就加上
		for _, item := range words {
			if _, has := mapData[item]; !has {
				mapData[item] = 1
			} else {
				mapData[item]++
			}
		}
	}
}

//3.将mapData中的数据写入到.txt中
func writeToFile(mapData map[string]int) {
	f, err := os.Create("G://GolangFileText/words.txt")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	defer f.Close()
	for k, v := range mapData {
		_, _ = f.Write([]byte(k + "---" + strconv.Itoa((v)) + "\n"))
	}
}

func main() {
	filesSlice := make([]string, 0)
	mapData := make(map[string]int)
	filesSlice = readCatelog("G://GolangFileText", filesSlice)
	fmt.Println(filesSlice)
	for _, item := range filesSlice {
		readFile("G://GolangFileText/"+item, mapData)
	}
	for k, v := range mapData {
		fmt.Printf("%s,%d\n", k, v)
	}
	writeToFile(mapData)
}
