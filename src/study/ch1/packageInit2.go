package main

import (
	"fmt"
	"study/ch1/customerunit"
)

func init() {
	fmt.Println("packageInit2 init")
}

func main() {
	fmt.Printf("customerunit.B : %d\n", customerunit.B)
}
