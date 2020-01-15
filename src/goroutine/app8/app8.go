package main

import "fmt"

func main() {
	// 创建一个通道
	c := make(chan int)
	// 并发执行printer
	go printer(c)

	for i := 1; i <= 10; i++ {
		// 将数据通过channel投给printer
		c <- i
	}
	//通知printer 结束循环
	c <- 0
	// 等待printer 结束
	<-c
}

func printer(ch chan int) {
	for true {
		data := <-ch
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
	ch <- 0
}
