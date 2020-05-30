package main

import (
	"fmt"
	"runtime"
)

const Pi = 3.1415926
var str = "hello world!"
func main() {
	fmt.Println(Pi)
	fmt.Println(str)
	goos := runtime.GOOS
	fmt.Printf("The operating system is: %s\n",goos)
}
