package main

import "fmt"

func main() {
	// 创建一个数组,长度为4
	array := [4]int{}
	fmt.Printf("array :%v\n", array)
	// 从数组中窃取一部分做为切片
	s1 := array[0:2]
	fmt.Printf("len(s1) = %d, cap(s1) = %d, array = %v, s1 = %v\n", len(s1), cap(s1), array, s1)
	//添加一个元素
	s2 := append(s1, 100)
	fmt.Printf("len(s2) = %d, cap(s2) = %d, array = %v, s2 = %v\n", len(s2), cap(s2), array, s2)
	// 添加两个元素
	s3 := append(s2, 200, 300)
	fmt.Printf("len(s3) = %d, cap(s3) = %d, array = %v, s3 = %v\n", len(s3), cap(s3), array, s3)
}
