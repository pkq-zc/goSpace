package main

import "fmt"

func main() {
	// new
	var num *int
	i := 0
	num = &i
	fmt.Println(*num)

	var num2 = new(int)
	fmt.Println(*num2)
	fmt.Println(num2)

	//make
	ints := make([]int,5)
	fmt.Println(len(ints))
	for index, value := range ints {
		fmt.Printf("index is [%d],value is [%d]\n",index,value)
	}
}
