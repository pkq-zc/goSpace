package main

import "fmt"

func main() {
	array := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	//截取所有数组元素做为slice
	s1 := array[:]
	//打印s1
	fmt.Printf("s1 address is:%p,values is :%v\n", &s1, s1)
	//更新
	updateSlice(s1)
	//打印s1
	fmt.Printf("s1 address is:%p,values is :%v\n", &s1, s1)
	//打印数组
	fmt.Printf("array is:%v\n", array)
}

func updateSlice(s []string) {
	//打印地址和值内容
	fmt.Printf("update s before address is:%p,values is :%v\n", &s, s)
	//将第一个元素更新为大写的A
	s[0] = "A"
	//再次打印结果
	fmt.Printf("update s after address is:%p,values is :%v\n", &s, s)
}
