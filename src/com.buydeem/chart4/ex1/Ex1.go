package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(f1())
	fmt.Println("==============")
	i := 0
	f2(&i)
	fmt.Println(i)

	fmt.Println(fibonacci(5))

	printNum(10)

	fmt.Println("==================")
	f := fibonacci2()
	for i := 0;i < 10;i++{
		fmt.Println(f())
	}

	jpgSuffix := MakeAddSuffix(".jpg")
	fmt.Println(jpgSuffix("apple"))
	fmt.Println(jpgSuffix("banana.jpg"))

	txtSuffix := MakeAddSuffix(".txt")
	fmt.Println(txtSuffix("history"))
	fmt.Println(txtSuffix("math.txt"))
}

func f1() int {
	i := 0
	fmt.Printf("before i = %d,&i = %p\n",i,&i)
	defer func() {
		i++
		fmt.Printf("defer i = %d,&i = %p\n",i,&i)
	}()
	return i
}

func f2(i *int) {
	fmt.Printf("before i = %d,&i = %p\n",*i,i)
	defer func() {
		*i++
		fmt.Printf("defer i = %d,&i = %p\n",*i,i)
	}()
}

func fibonacci(i int) int {
	if i <= 1 {
		return 1
	}else {
		return fibonacci(i -1) + fibonacci(i - 2)
	}
}

func printNum(num int) int {
	if num <= 0 {
		return num
	}else {
		fmt.Println(num)
		return printNum(num-1)
	}
}

//闭包的使用案例
func fibonacci2() func() int {
	a := 0
	b := 1
	return func() int {
		a,b = b,a+b
		return b
	}
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name,suffix) {
			return name + suffix
		}
		return name
	}
}