package main

import "fmt"

const Pi float64 = 3.1415926
const Key  = "Hello world"
// 定义多个常量
const (
	a = iota
	b
	c
)

// 色彩枚举
type Color int
const (
	RED Color = iota
	BLACK
	YELLOW
)
func main() {
	fmt.Printf("Pi = %f\n",Pi)
	fmt.Printf("Key = %s\n",Key)
	fmt.Printf("a = %d\n",a)
	fmt.Printf("b = %d\n",b)
	fmt.Printf("c = %d\n",c)
	fmt.Printf("RED = %d\n",RED)
	fmt.Printf("BLACK = %d\n",BLACK)
	fmt.Printf("YELLOW = %d\n",YELLOW)
}
