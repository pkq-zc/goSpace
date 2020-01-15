package main

import (
	"fmt"
	"time"
)

func main() {
	// 1代表缓冲区大小,如果缓存区大小为0,通道中的数据未被消费,将会导致阻塞
	ch := make(chan string, 1)
	ch <- "hello"
	go f1(ch)
	time.Sleep(1 * 1e9)
}

func f1(ch chan string) {
	fmt.Println(<-ch)
}
