package main

import "fmt"

// 与java中HashMap类似
// key不能是切片和数组,可是是string,int等,也可以是结构体,但需要提供Key()和Hash()方法
// 永远不要用new创建map
func main() {
	// 创建
	var m1 map[string]string
	fmt.Println(m1)

	var m2 = map[string]string{"name": "mac", "age": "18"}
	fmt.Println(m2)
	// 创建map,设置容量为5
	m3 := make(map[string]string, 5)
	fmt.Println(m3)
	// 获取值
	fmt.Printf("m2[\"%s\"] = %s\n", "name", m2["name"])
	age, ok := m2["age"]
	if ok {
		fmt.Printf("age is :%s\n", age)
	}
	if age, ok := m2["age"]; ok {
		fmt.Printf("age is :%s\n", age)
	}
	//删除
	delete(m2, "age")
	_, ok = m2["age"]
	fmt.Printf("m2[\"age\"] existe :%t\n", ok)
}
