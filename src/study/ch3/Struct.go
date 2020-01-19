package main

import (
	"fmt"
)

//定义一个员工结构体
type Employee struct {
	Id       int
	Name     string
	Address  string
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

type Point struct { x,y int}

func main() {
	var mac Employee
	mac.Name = "mac"
	name := &mac.Name
	*name = *name + "-2"
	fmt.Printf("mac address :%p,mac = %+v\n",&mac,mac)
	//调用 无法修改Employee的值
	fmt.Println("=====================================")
	updateEmployee(mac)
	fmt.Printf("mac address :%p,mac = %+v\n",&mac,mac)
	//调用 可以修改Employee的值
	fmt.Println("=====================================")
	updateEmployee2(&mac)
	fmt.Printf("mac address :%p,mac = %+v\n",&mac,mac)

	// 结构体比较
	p := Point{1,2}
	q := Point{2,1}
	fmt.Println(p.x == q.x && p.y == q.y)
	fmt.Println(q == q)
}

func updateEmployee(e Employee) {
	fmt.Printf("update before: mac address :%p,mac = %+v\n",&e,e)
	e.Address = "更新地址"
	fmt.Printf("update after: mac address :%p,mac = %+v\n",&e,e)
}

func updateEmployee2(e *Employee) {
	fmt.Printf("update before: mac address :%p,mac = %+v\n",e,e)
	e.Address = "更新地址2"
	fmt.Printf("update after: mac address :%p,mac = %+v\n",e,e)
}
