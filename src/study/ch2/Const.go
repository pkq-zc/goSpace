package main

import "fmt"

//声明一个pi常量
const pi = 3.1415926

//声明一个name常量,类型可以自己推导出来
const name string = "tom"

//多个常量,还可以这样简写方便
const (
	author         = "莫言"
	country string = "中国"
	//内置函数len()计算常量的长度也是常量
	l = len(country)
)

const (
	c1 = iota + 5
	c2
	c3
	c4
	c5
)

func main() {
	fmt.Printf("c1 = %d,c2 = %d,c3 = %d,c4 = %d,c5 = %d\n", c1, c2, c3, c4, c5)
	var f1 float32 = 1.2
	var f2 float64 = 13.3
	fmt.Println(f1 + pi)
	fmt.Println(f2 + pi)

}
