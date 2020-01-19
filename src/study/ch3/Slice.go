package main

import "fmt"

func main() {
	//初始化一个长度为4,容量为10的slice
	s1 := make([]int, 4, 10)
	fmt.Printf("len(s1) = %d,cap(s1) = %d,s1 is :%v\n", len(s1), cap(s1), s1)
	//初始化一个slice,返回指针类型的slice
	s2 := new([]int)
	fmt.Printf("len(s2) = %d,cap(s2) = %d,s2 is :%v\n", len(*s2), cap(*s2), s2)
	//声明一个 slice
	s3 := []int{0, 1, 2, 3}
	fmt.Printf("len(s3) = %d,cap(s3) = %d,s3 is :%v\n", len(s3), cap(s3), s3)
	// 从数组中生成一个slice
	array := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	s4 := array[2:8]
	fmt.Printf("len(s4) = %d,cap(s4) = %d,s4 is :%v\n", len(s4), cap(s4), s4)

	demo1()
}

func demo1() {
	array := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	s1 := array[2:5]
	fmt.Printf("len(s1) = %d,cap(s1) = %d,s1 is :%v\n", len(s1), cap(s1), s1)
	s2 := array[5:9]
	fmt.Printf("len(s2) = %d,cap(s2) = %d,s2 is :%v\n", len(s2), cap(s2), s2)
}
