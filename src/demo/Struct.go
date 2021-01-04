package main

import (
	"fmt"
	"unsafe"
)

/*
结构体：
	是一种数据类型
	type person【】 struct{
		属性
	}

结构体比较： ==   !=

结构体地址
	结构体变量的地址 == 结构体首个元素的地址

函数不能返回局部变量的地址值，因为局部变量保存在栈帧上，函数调用结束后，栈帧释放。局部变量的地址不再受到系统保护。

*/
//结构体定义(通常放在全局位置)  **不能加var
//结构体是有地址的！！！！！作为形参只有地址传过去的时候才可以进行修改
type person struct {
	name string
	sex  string
	age  int
}

//结构体初始化
func structCreat() {
	//1.初始化
	man := person{"wangkedong", "man", 22}
	//2.部分初始化,为未初始化的成员变量，去该类型的默认值
	woman := person{name: "wanglirong", sex: "woman"}
	fmt.Println(man, woman)
	//3.简单初始化
	var student person
	student.sex = "man" //对变量进行赋值
	fmt.Println("修改前：", man)
	//alterPerson(man)---->值传递，不会改变
	alterPerson(&man)
	fmt.Println("修改后", man)
	fmt.Println(unsafe.Sizeof(man)) //获取大小，相当于sizeof()  不论什么类型的指针，在64位操作系统下都是8字节
	fmt.Println("结构体地址：", &man)
	fmt.Println("结构体首元素地址：", &man.name)
}

//初始化二
func structCreat2(man_name *string) {
	*man_name = "袖带"
}

//结构体做参数
//几乎不用，内存消耗太大了，效率低,利用指针传递
func alterPerson(man *person) {
	man.name = "修改"
}

func main() {
	man := new(person)
	structCreat2(&man.name)
	fmt.Println(man)
}
