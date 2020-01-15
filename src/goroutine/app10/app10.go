package main

import "fmt"

func main() {
	// 创建一个3个元素缓存大小的整数通道
	ch := make(chan int, 3)

	//查看当前通道大小
	fmt.Println(len(ch))

	// 发送3个元素
	ch <- 3
	ch <- 2
	ch <- 1

	//查看当前通道大小
	fmt.Println(len(ch))
}
