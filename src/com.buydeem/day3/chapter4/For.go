package main

import "fmt"

const (
	FIZZ      = 3
	BUZZ      = 5
	FIZZ_BUZZ = 15
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("The is the %d iteration\n", i)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i = %d,j = %d\n", i, j)
		}
	}

	// 99乘法表
	fmt.Printf("99乘法表\n")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", i, j, i*j)
		}
		fmt.Printf("\r\n")
	}

	// fizzBuzz
	FizzBuzz(100)

	//打印矩形
	rectangle_stars()

	// while
	maxNum := 10
	for maxNum > 0 {
		fmt.Printf("%d\n", maxNum)
		maxNum--
	}

	// range
	str := "hello 中国"
	for index, char := range str {
		fmt.Printf("index = %d,value = %c\n", index, char)
	}

}

func FizzBuzz(maxNum int) {
	for i := 1; i <= maxNum; i++ {
		switch {
		case i%FIZZ_BUZZ == 0:
			fmt.Printf("%d is FIZZ_BUZZ\n", i)
		case i%FIZZ == 0:
			fmt.Printf("%d is FIZZ\n", i)
		case i%BUZZ == 0:
			fmt.Printf("%d is BUZZ\n", i)
		}
	}
}

func rectangle_stars() {
	for high := 1; high <= 10; high++ {
		for width := 1; width <= 20; width++ {
			fmt.Print(" * ")
		}
		fmt.Print("\n")
	}
}
