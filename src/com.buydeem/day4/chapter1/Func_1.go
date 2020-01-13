package main

import (
	"fmt"
	"io"
	"log"
)

// 1.普通带有名字的函数/匿名函数/方法
func main() {
	a, b, c := mult_returnval(5, 10)
	fmt.Printf("a = %d,b = %d,c = %d\n", a, b, c)
	a2, b2, c2 := mult_returnval2(5, 10)
	fmt.Printf("a2 = %d,b2 = %d,c2 = %d\n", a2, b2, c2)

	var value int
	Multiply(5, 10, &value)
	fmt.Printf("value is %d\n", value)

	min := min(1, 5, 9, 2, 0)
	fmt.Printf("min value is: %d\n", min)

	func1("Go")
}

func mult_returnval(x, y int) (int, int, int) {
	return x + y, x * y, x - y
}

func mult_returnval2(x, y int) (a, b, c int) {
	a = x + y
	b = x * y
	c = x - y
	return
}

// 改变外部变量
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

// 传递变长参数
func min(num ...int) int {
	if len(num) == 0 {
		return 0
	}
	min := num[0]
	for _, value := range num {
		if min > value {
			min = value
		}
	}
	return min
}

//defer
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d,%v", s, n, err)
	}()
	return 7, io.EOF
}
