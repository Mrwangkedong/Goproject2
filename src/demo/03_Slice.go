package main

import (
	"fmt"
)

//切片
/*
为什么要用切片：
	1、数组的容量固定，不能自动拓展
 	2、值传递，数组作为函数参数时，将整个数组值拷贝一份给形参（可怕）
	在Go语言当中，几乎在所有的场景中去替换数组使用
切片的本质：
	不是一个数组的指针，是一个数据结构体，用来操作数组内部元素
如何得到切片
	切片可以由数组来，也可以有切片来
切片做函数引用
	传引用（传地址）
*/

func creatSlice() {
	//1、自动推导类型
	arr := [8]int{1, 2, 3, 4, 5, 6, 8}
	_ = arr

	slice := make([]int, 5, 6) //make(Type,len,cap)
	_ = slice

	slice2 := make([]int, 5) //常用：make(Type,len),cap随着长度变化而变化
	//make只能创建slice、map、channel(通道)，并且返回一个有初始值（非0）的对象
	fmt.Println(slice2)        //[0 0 0 0 0]
	slice2 = append(slice2, 2) //[0 0 0 0 0 2]
	slice2 = append(slice2, 3) //[0 0 0 0 0 2 3]

	fmt.Println(slice2)
}

//slice De append 操作
func sliceTest() {

	data := []string{"red", "", "black", "", "pink"}
	afterdata := noEmpty(data)
	fmt.Println(afterdata)
}

//去除切片中“”
func noEmpty(data []string) []string {
	afterdata := make([]string, 0)
	for _, str := range data {
		if str != "" {
			afterdata = append(afterdata, str)

		}
	}
	return afterdata
}

//slice的copy函数，会从des的0处开始进行覆盖。
func sliceCopy() {
	data := []string{"red", "", "black", "", "pink"}
	data2 := make([]string, 5, 10)
	data2[0] = "oook"
	copy(data2, data[2:3])
	fmt.Println(data2)
	fmt.Println(cap(data2), len(data2))
}

//slice的remove
func sliceRemove(data []int) {
	//去除3，也就是index=2的数
	index := 2
	copy(data[index:], data[index+1:]) //[1 2 4 5 5]
	afterData := data[:len(data)-1]    //[1 2 4 5]
	fmt.Println(afterData)
	sliceCopy()
}

func main() {

	/*
		arr := [8]int {1,2,3,4,5,6,8}  //若是【】不指定大小，则默认规定为切片

		s := arr[0:4:7]  //且片名称【low:high:max】  max>=high-low

		fmt.Println(s)    //[1,2,3,4]
		fmt.Println("cap(s) ",cap(s))     //cap表示切片的容量，指low到arr最后一个的长度，加入arr实际里面有4个，但是arr[8]，按照8作为最后一个
								//若在切片定义是进行规定，按规定的Max来
		fmt.Println("len(s) ",len(s))		//len表示切片的长度，指low到s最后一个的长度

		s2 := s[2:5]	   //虽然s的长度为4,表示（1,2,3,4）但是，他的容量为7，他的真实存在是【1,2,3,4,5,6,8】
		fmt.Println(s2)		//[3,4,5]
		fmt.Println("cap(s2) ",cap(s2))
		fmt.Println("len(s2) ",len(s2))*/

	//creatSlice()

	//sliceTest()

	//sliceCopy()

	data := []int{1, 2, 3, 4, 5}
	sliceRemove(data)
	fmt.Println(data)
}
