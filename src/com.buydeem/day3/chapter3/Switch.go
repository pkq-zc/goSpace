package main

import "fmt"

func main() {
	var num1 = 98
	//
	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

	//
	var i = 10
	switch {
	case i < 0:
		fmt.Println("i < 0")
	case i == 0:
		fmt.Println("i == 0")
	case i > 0:
		fmt.Println("i > 0")
	}

	//
	fmt.Printf("12月是:%s\n",Season(12))
	fmt.Printf("3月是:%s\n",Season(3))
	fmt.Printf("6月是:%s\n",Season(6))
	fmt.Printf("10月是:%s\n",Season(10))
}

func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return "冬季"
	case 3, 4, 5:
		return "春季"
	case 6, 7, 8:
		return "夏季"
	case 9, 10, 11:
		return "秋季"
	}
	return "未知季节"
}
