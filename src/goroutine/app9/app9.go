package main

import "fmt"

// 单向通道
func main() {
	//只能发送通道
	var ch1 chan<- int
	//只能接收通道
	var ch2 <-chan int

	fmt.Println(ch1)
	fmt.Println(ch2)

}
