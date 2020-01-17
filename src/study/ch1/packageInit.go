package main

import "fmt"

var a = b + c //3.最后初始化a
var b = f()   //2.通过调用f初始化b
var c = 1     //1.首先初始化c

func main() {
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
}

func f() int {
	return c + 1
}
