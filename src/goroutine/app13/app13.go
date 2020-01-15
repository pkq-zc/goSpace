package main

import (
	"fmt"
	"time"
)

// 定点计时器
func main() {
	// 创建一个打点器,每500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)
	//创建一个计时器,2秒后触发
	timer := time.NewTimer(time.Second * 2)

	//声明计数变量
	var i int
	for {
		select {
		case <-timer.C:
			fmt.Println("stop")
			goto StopHere
		case <-ticker.C:
			i++
			fmt.Println("tick", i)
		}
	}
StopHere:
	fmt.Println("done")
}
