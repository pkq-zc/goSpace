package main

import "fmt"

//声明
func main() {
	// 同时给i,j,k赋值
	i, j, k := 1, 2, 3
	fmt.Printf("i=%d,j=%d,k=%d\n", i, j, k)
	// 可以同时给不同类型的变量赋值
	number, str := 0, "hello world"
	fmt.Printf("number = %d,str = %s\n", number, str)
	// 交换i和j的值,不需要借助第三个临时变量
	i, j = j, i
	fmt.Printf("i=%d,j=%d\n", i, j)
}
