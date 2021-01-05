package main

import (
	"fmt"
	"strings"
)

/*
map 不能使用cap
*/

func creatMap() {
	//1
	var m1 map[int]string //这样m1 == nil
	if m1 == nil {
		fmt.Println("m1是空指针")
	}
	//2
	m2 := map[int]string{}
	fmt.Println(len(m2))
	m2[2] = "wangkedong"
	fmt.Println(len(m2))
	//3
	m3 := make(map[int]string)
	m3[1] = "123"
	fmt.Println(len(m3))

	var m4 map[int]string = map[int]string{1: "123", 2: "456"}
	fmt.Println(m4)

	array := make([]int, 0)
	array = append(array, 1)
	array = append(array, 5)
	array = append(array, 3)
	m5 := make(map[int][]int)
	m5[1] = array
	fmt.Println(m5[1])
}

//map的删除，添加，遍历，判断是否存在
func useMap() {
	var m1 map[int]string = map[int]string{1: "123", 2: "456"}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//range返回的k，省略value
	for v := range m1 {
		fmt.Println(v) //1 ， 2
	}

	//判断map中的key是否存在
	//m1[5]不存在，则返回nil，也不报错
	fmt.Println(m1[5]) //m1[index]返回两个值，第一个是value，第二个是bool，代表可以是否存在

	if _, has := m1[9]; !has {
		fmt.Println("不存在")
	} else {
		fmt.Println("存在")
	}

	//删除map -- delete(map,key)
	fmt.Println("删除之前", m1)
	//delete(m1, 1)
	MapDelete(m1, 1)
	fmt.Println("删除之后", m1)
	//利用传参数进行删除

}

func MapDelete(mapData map[int]string, key int) {
	delete(mapData, key)
}

func MapTest() {
	str := "I love my work and I love my family too"
	m := make(map[string]int)
	_ = m
	kv := strings.Split(str, " ") //将字符串按照“ ”划分为数组  【I love my work and I love my family too】

	for _, key := range kv {
		if _, has := m[key]; !has {
			m[key] = 1
		} else {
			m[key]++
		}
	}
	fmt.Println(m)
}

func main() {
	//creatMap()

	//useMap()

	MapTest()
}
