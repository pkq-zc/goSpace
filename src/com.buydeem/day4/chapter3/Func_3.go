package main

import "fmt"

func main() {
	var fPlus = func(x, y int) int { return x + y }
	fmt.Println(fPlus(1, 2))

	var fv = func() { fmt.Println("hello world") }
	fv()
	fmt.Println(f())

	jpg := suffix(".jpg")
	gif := suffix(".gif")
	fileName := "sun"
	fmt.Printf("jpg(%s) = %s\n", fileName, jpg(fileName))
	fmt.Printf("gif(%s) = %s\n", fileName, gif(fileName))
	fileName = "book"
	fmt.Printf("jpg(%s) = %s\n", fileName, jpg(fileName))
	fmt.Printf("giff(%s) = %s\n", fileName, gif(fileName))

	// fibonacci
	fb := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fb())
	}
}

func f() (res int) {
	defer func() {
		res++
	}()
	return 1
}

func suffix(suffix string) func(name string) string {
	return func(name string) string {
		return name + suffix
	}
}

func fibonacci() func() int {
	x := -1
	y := 1
	return func() int {
		x, y = y, x+y
		return y
	}
}
