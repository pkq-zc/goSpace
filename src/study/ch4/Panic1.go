package main

import "fmt"

func main() {
	//p2(";")

	r := p3(1, 0)
	fmt.Println(r)
}

//
func p1(index int) {
	arrays := [3]int{1,2,3}
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()
	defer func() {
		fmt.Println("defer 3")
	}()
	fmt.Println(arrays[index])
	fmt.Println("method over")
}

func p2(operation string) string{
	switch operation {
	case "+":
		return "加法"
	case "-":
		return "减法"
	case "*":
		return "乘法"
	case "/":
		return "除法"
	default:
		panic(fmt.Sprintf("未知的操作%s\n",operation))
	}
}

func p3(x,y int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	r := x / y
	fmt.Println(r)
	return r
}