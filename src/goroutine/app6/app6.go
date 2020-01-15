package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpu := runtime.NumCPU()
	fmt.Printf("cpu core is:%d\n", cpu)
	// 通道发送数据 通道 <- 数据
	// 通道接收数据 数据 <- 通道

	// 构建通道
	ch := make(chan int)

	//开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		ch <- 0
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	//等待匿名函数结束
	<-ch
	fmt.Println("all done")
}
