package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sum(12, 13, ch)
	fmt.Println(<-ch)
}

func sum(x, y int, ch chan int) {
	time.Sleep(5 * 1e9)
	ch <- x + y
}
