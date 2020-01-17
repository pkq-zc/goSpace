package main

import "fmt"

func main() {
	//使用内置函数new创建变量
	i := new(int)
	//打印i的类型和值
	fmt.Printf("i type is:%T,value is:%d\n", i, *i)
}
