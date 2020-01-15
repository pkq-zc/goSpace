package main

import "fmt"

// 关闭通道
func main() {
	// 向关闭的通道发送数据会异常
	//closeChannel1()
	// 向关闭的通道获取数据不会异常
	closeChannel2()
}

func closeChannel1() {
	// 创建通道
	ch := make(chan int)
	//关闭通道
	close(ch)
	//打印通道的指针和容量
	fmt.Printf("ptr:%p cap:%d len:%d\n", ch, cap(ch), len(ch))
	//给关闭的通道发送数据 会导致异常
	ch <- 1
}

//从已经关闭的通道接收数据或者正在接收数据时，将会接收到通道类型的零值，然后
//停止阻塞并返回。
func closeChannel2() {
	//创建通道
	ch := make(chan int, 4)
	//给通道中放入数据
	ch <- 0
	ch <- 1
	// 关闭通道
	close(ch)
	// 遍历通道内的数据
	for i := 0; i < cap(ch)+1; i++ {
		// 从通道中取出数据
		data, ok := <-ch
		// 打印取出数据的状态
		fmt.Println(data, ok)
	}
}
