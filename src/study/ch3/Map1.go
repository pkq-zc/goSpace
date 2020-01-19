package main

import "fmt"

func main() {
	//使用make创建一个key为string,值类型也为string的map
	m1 := make(map[string]string)
	fmt.Println(m1)
	//使用字面量创建一个带初始化的map
	user := map[string]string{
		"name": "tom",
		"age":  "18",
	}
	fmt.Println(user)

	delete(user, "sex")

	ages := map[string]int{
		"tom":  18,
		"jack": 19,
	}

	//删除
	delete(ages, "mac")
	//获取 不存在会返回0值
	fmt.Println(ages["mac"])
	//赋值
	ages["tom"] = 19
	//使用快捷方式赋值
	ages["jack"] += 2
	fmt.Println(ages)
	// 判断key是否存在map
	if _, ok := ages["tom"]; ok {
		fmt.Println("tom exist")
	}
	if _, ok := ages["mac"]; ok {
		fmt.Println("mac exist")
	}

	// m1为nil,将导致异常
	var m2 map[string]string
	//通过key获取元素
	fmt.Println(m2["mac"])
	//将导致异常
	m2["mac"] = "test"
}
