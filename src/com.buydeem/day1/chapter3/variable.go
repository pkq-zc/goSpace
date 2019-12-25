package main

import (
	"com.buydeem/day1/chapter3/customer"
	"fmt"
)

//定义名称
var name string = "小白"
//定义年龄 可以推算出类型
var age = 18
//定义多个
var high,width = 180,60
var score,address = 90,"深圳市"
//定义在一起
var(
	one = 1
	pi = 3.14
)
// 初始化不赋值,使用默认值
var str string
var number int

// 变量
func main() {
	fmt.Println(customer.CanSeeInOtherPackage)
	//fmt.Println(customer.notCanSeeInOtherPackage) 无法在其他包中访问
	customer.SayHello()
	var tempNews customer.News = "this is news content"
	fmt.Println(tempNews)

	fmt.Printf("str = %s\n",str)
	fmt.Printf("number = %d\n",number)

	//函数体内推荐 a:= 1 这种写法
	job := "go dev"
	fmt.Println(job)

	//值类型 主要包括 int float bool string等基础类型,同时包括结构体和数组等,值类型变量指向内存中的值
	x := 1
	// 使用 = 时, 将x的值拷贝给了y ,x值修改并不会影响y值
	y := x
	fmt.Printf("x = %d, y = %d\n",x,y)
	x = 2
	fmt.Printf("x = %d, y = %d\n",x,y)

	// 引用类型 主要包括slice(切片),map(集合),channel等等 引用类型指定变量的值的内存地址值,而这个地址就是叫做指针
	user  := map[string]string{
		"name":"小黑",
		"age":"18",
	}

	user2 := user
	fmt.Println(user)
	fmt.Println(user2)
	// 修改user2的值,会影响user的值
	user2["address"] = "深圳市"
	fmt.Println(user)
	fmt.Println(user2)

}

