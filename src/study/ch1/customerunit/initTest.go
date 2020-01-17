package customerunit

import "fmt"

var A = f()
var B int

func init() {
	fmt.Println("customerunit init start")
	fmt.Printf("A = %d,B = %d\n", A, B)
	B = 2
	fmt.Println("customerunit init over")
}

func f() int {
	fmt.Println("A init")
	return 1
}
