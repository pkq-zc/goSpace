package main

// 定义一个Person结构体
type Person struct {
	Name string
	age int
}

// 定义一个Student结构体
type Student struct {
	P Person
	class string
}

type Chinese struct {
	country string
	Person
}

func main() {
	//
	var student Student
	student.P.age = 18
	student.P.Name = "小明"
	student.class = "1班"
	//
	var chinese Chinese
	chinese.Name = "李小龙"
	chinese.age = 23
	chinese.country = "中国"
}
