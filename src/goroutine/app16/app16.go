package main

import (
	"fmt"
	"sync"
)

var (
	// 逻辑中使用的某个变量
	count int
	// 与变量对应使用的互斥锁
	countGuard sync.Mutex
)

func main() {
	setCount(1)
	fmt.Println(GetCount())
}

func GetCount() int {
	//锁定
	countGuard.Lock()
	//在函数退出时解除锁定
	defer countGuard.Unlock()
	return count
}

func setCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}
