package main

import (
	season "com.buydeem/chart3/ex1"
	"fmt"
)

func main() {
	var i int = 3
	fmt.Printf("i = %d, *i = %p\n", i, &i)
	var p *int = &i
	fmt.Printf("i = %d, *i = %p\n", *p, p)
	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	var str string
	if str == "" {
		fmt.Println("str is null")
	}
	var pStr *string
	if pStr == nil {
		fmt.Println("pStr is null")
	}

	if v := 10; v > 20 {
		fmt.Println("v > 20")
	} else {
		fmt.Println("v <= 20")
	}

	// switch
	operation := "+"
	switch operation {
	case "+":
		{
			fmt.Println("加号")
		}
	case "-":
		{
			fmt.Println("减号")
		}
	case "*":
		{
			fmt.Println("乘号")
		}
	case "/":
		{
			fmt.Println("除号")
		}
	default:
		{
			fmt.Println("未知")
		}
	}

	//switch
	num := 10
	switch {
	case num < 10:
		{
			fmt.Println("num < 10")
		}
	case num == 10:
		{
			fmt.Println("num == 10")
		}
	case num > 0:
		fmt.Println("num > 10")
	}

	fmt.Println(season.Season(1))

	// for
	for i := 0;i<5;i++{
		fmt.Printf("i = %d\n",i)
	}

	//
	j := 5
	for j > 0{
		fmt.Printf("j = %d\n",j)
		j--
	}

	//字符串 for
	str = "hello 中国"
	for i := 0;i< len(str);i++{
		fmt.Printf("i = %d, v = %c\n",i,str[i])
	}
	for i, v := range str {
		fmt.Printf("i = %d, v = %c\n",i,v)
	}
}