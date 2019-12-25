package main

import (
	"fmt"
	"strconv"
)

// 变量类型和运算
func main() {
	// boole类型
	success := true
	fmt.Printf("success is :%t\n",success)
	//整型
	//与平台相关的 int uint uintptr
	//与平台无关的 int8(-128~127) int16 int32 int64 uint8 uint16 uint32 uint64

	// 只有类型相同的变量才能进行二元运算结合,不支持隐式转换,也不支持运算符重载
	str1 := "age is :"
	age := 10
	//fmt.Println(str1+age)
	fmt.Println(str1 + strconv.Itoa(age))
}
