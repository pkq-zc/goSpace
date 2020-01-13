package main

import (
	"fmt"
	"strconv"
)

func main() {
	if true {
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}

	val := 10
	if val > 20 {
		fmt.Println("val > 20")
	}else {
		fmt.Println(val)
	}

	if val2 := 10; val2 > 20 {
		fmt.Println("val2 > 20")
	}else {
		fmt.Println(val2)
	}

	var orig string = "10"
	fmt.Printf("The size of ints is:%d\n",strconv.IntSize)
	i, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n",orig)
		return
	}
	fmt.Printf("The integer is %d\n",i)
	i = i + 5
	var newS string
	newS = strconv.Itoa(i)
	fmt.Printf("The new string is: %s\n",newS)
}
