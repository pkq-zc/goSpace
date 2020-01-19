package main

import (
	"fmt"
	"time"
)

//定义一个员工结构体
type Employee struct {
	Id       int
	Name     string
	Address  string
	DoB      time.Time
	Position string
}

type Node struct {
	// 数据
	data int
	// 父节点
	parent *Node
	// 子节点
	children []*Node
}

func main() {
	var mac Employee
	mac.Name = "mac"
	name := &mac.Name
	*name = *name + "-2"
	fmt.Println(mac)
}
