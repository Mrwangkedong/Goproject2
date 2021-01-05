package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//创建文件
func fileCreat(file_outlook string) {
	f, err := os.Create(file_outlook + "first.txt") //返回一个文件指针，返回错误信息
	fmt.Println(f)
	fmt.Println(err)
	defer f.Close()
}

//打开文件
func fileOPen(file_outlook string) {
	f, err := os.Open(file_outlook + "first.txt") //**只读**  方式打开
	//打开文件
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	defer f.Close()
	//写入文件
	_, err = f.WriteString("加油！") //error： Access is denied.can't write
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

}

//写入文件
func fileWrite(file_outlook string) {
	f, err := os.OpenFile(file_outlook+"first.txt", os.O_RDWR, 6) //os.OpenFile(name string, flag int, perm FileMode)----文件打开路径，打开文件权限，一般传6
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	//写入文件内容
	//按照字符写
	_, err = f.WriteString("1234") //直接就覆盖了  utf-8的格式
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	//按照位置写  seek() 修改文件的读写位置  参1：偏移量，正，向文件尾偏，负，向文件头偏
	//  参2：偏移的起始位置：io.SeekStart   io.SeekCurrent   io.SeekEnd
	//	返回值：ret从文件起始位置到当前文件位置的偏移量
	off, _ := f.Seek(3, io.SeekStart)
	fmt.Println("当前偏移:", off)
	//按照字节写，通常搭配seek
	n, _ := f.WriteAt([]byte("1111"), off) //func (f *File) WriteAt(b []byte, off int64) (n int, err error)
	fmt.Println("WriteAt写入的len：", n)
	//按照字节写

	buf := []byte(("okok"))
	_, _ = f.Write(buf)

}

//文件读入
func fileRead(file_outlook string) {
	f, err := os.OpenFile(file_outlook+"first.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	defer f.Close()
	//按行来读  1、创建一个reader
	//创建一个带有缓冲区的reader
	reader := bufio.NewReader(f) //1、创建reader
	for {
		buf, err := reader.ReadBytes('\n') //2、读数据  ReadBytes[split_item]  读取1-->返回49
		if err != nil && err == io.EOF {   //3.判断文件结尾标志(需要单独去读)--io.EOF
			fmt.Println("ERROR: ", err)
			return
		}
		fmt.Println(string(buf))
	}

}

//练习--将一个文件中的内容读到另一个文件中
func fileMove(fromFile string, toFile string) {
	//1.读入文件
	freader, err := os.OpenFile(fromFile, os.O_RDWR, 6)
	if err != nil {
		fmt.Println("打开读入文件错误：: ", err)
	}
	defer freader.Close()
	//2.创建写入文件
	fwriter, err := os.Create(toFile)
	if err != nil {
		fmt.Println("创建x写入文件错误：: ", err)
	}
	defer fwriter.Close()
	//3.读并，给予写入文件权限
	fwriter, err = os.OpenFile(toFile, os.O_RDWR, 6)
	if err != nil {
		fmt.Println("读取写入文件错误：: ", err)
	}
	//4.利用循环进行文件读取
	reader := bufio.NewReader(freader) //1、创建reader
	off := 0                           //创建光标位置，用于下次写入
	for {
		//按行进行读取
		buf, err := reader.ReadBytes('\n') //2、读数据  ReadBytes[split_item]  读取1-->返回49
		if err != nil && err == io.EOF {   //3.判断文件结尾标志(需要单独去读)--io.EOF
			fmt.Println("ERROR: ", err)
			return
		}
		//按行进行写入
		n, _ := fwriter.WriteAt(buf, int64(off))
		off += n
	}
}

//练习--将一个文件中的内容读到另一个文件中
func fileMove2(fromFile string, toFile string) {
	freader, err := os.OpenFile(fromFile, os.O_RDWR, 6)
	if err != nil {
		fmt.Println("打开读入文件错误：: ", err)
	}
	defer freader.Close()
	//2.创建写入文件
	fwriter, err := os.Create(toFile)
	if err != nil {
		fmt.Println("创建x写入文件错误：: ", err)
	}
	defer fwriter.Close()
	//3.创建buffer
	buf := make([]byte, 1024*4)
	for {
		n, err := freader.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读完了、、、、")
			return
		}
		_, _ = fwriter.Write(buf[:n])
	}
}

func main() {
	//file_outlook := "G://GolangFileText/"
	//fileCreat(file_outlook)
	//fileOPen(file_outlook)
	//fileWrite(file_outlook)
	//fileRead(file_outlook)

	fromFile := "G://GolangFileText/讲座.jpg"
	toFile := "G://GolangFileText/test.jpg"
	//fileMove(fromFile,toFile)           //move1 相比较 move2少了1kb。。。。。
	fileMove2(fromFile, toFile)
}
