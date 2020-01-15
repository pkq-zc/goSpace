package main

import (
	"fmt"
	"time"
)

// 一段时间之后
func main() {
	// 声明一个退出用的通道
	exit := make(chan int)

	// 打印开始
	fmt.Println("start")

	//过1秒后,调用匿名函数
	time.AfterFunc(time.Second, func() {
		//1秒后打印结果
		fmt.Println("one second after")
		//通知main()的goroutine已经结果
		exit <- 0
	})

	//等待结束
	<-exit
}
