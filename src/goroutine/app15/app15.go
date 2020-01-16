package main

import (
	"fmt"
	"sync/atomic"
)

var (
	seq int64
)

func main() {
	for i := 0; i < 10; i++ {
		go GenId()
	}
	fmt.Println(GenId())
}

func GenId() int64 {
	//atomic.AddInt64(&seq, 1)
	//return seq
	return atomic.AddInt64(&seq, 1)
}
