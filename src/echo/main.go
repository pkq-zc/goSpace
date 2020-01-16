package main

import "os"

func main() {
	// 创建一个程序结束的通道
	exitChan := make(chan int)
	// 将服务器并发运行
	go server("127.0.0.1:8080", exitChan)
	// 通道堵塞等待接收返回值
	code := <-exitChan
	// 标记程序返回值并退出
	os.Exit(code)
}
