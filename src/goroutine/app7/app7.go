package main

import (
	"fmt"
	"time"
)

func main() {
	// 构建通道
	ch := make(chan int)
	// 开启并发函数
	go func() {
		for i := 3; i >= 0; i-- {
			//发送
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	// 变量接收通道数据
	for data := range ch {
		fmt.Println(data)
		if data == 0 {
			break
		}
	}
}
