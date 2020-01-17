package main

import "fmt"

// 声明常量
const PI float64 = 3.1415926

// 声明类型
// 声明摄氏温度类型
type celsius float64

// 声明摄氏温度类型
type fahrenheit float64

func main() {
	// 声明变量
	var str string = "hello world"
	//调用方法
	print(str)
}

// 声明方法
func print(s string) {
	fmt.Println(s)
}
