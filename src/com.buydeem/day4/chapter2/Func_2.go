package main

import (
	"fmt"
)

func main() {
	//1
	for i := 1; i <= 10; i++ {
		fmt.Printf("fibonacci(%d) is : %d\n", i, fibonacci(i))
	}
	//2
	print(10)
	//3
	res := factorial(10)
	fmt.Printf("res = %d\n", res)
}

// 使用递归打印斐波拉契数
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

// 使用递归打印num - 1的值
func print(num int) {
	fmt.Println(num)
	if num == 1 {
		return
	} else {
		print(num - 1)
	}
}

// 计算阶乘
func factorial(n int) (res int) {
	if n == 0 {
		res = 1
	} else {
		res = n * factorial(n-1)
	}
	return
}
