package main

import "fmt"

func main() {
	//声明一个长度为3类型为int的数组
	var array1 [3]int
	fmt.Println(array1)
	//使用数组字面量初始化一个数组
	var array2 [3]int = [3]int{1, 2, 3}
	fmt.Println(array2)
	var array3 [3]int = [3]int{1, 2}
	fmt.Println(array3)
	// 使用...省略数组长度,由初始化元素个数决定
	array4 := [...]int{1, 2, 3, 4}
	fmt.Println(array4)
	// 定义一个长度为6,第一个元素为1,第六个元素为3的数组
	array5 := [...]int{0: 1, 5: 3}
	fmt.Println(array5)

	//数组作为参数传入,传入的是数组的副本
	array6 := [...]int{0, 0, 0}
	fmt.Printf("array update before values is:%v,address :%p\n", array6, &array6)
	// 传入数组
	update(array6)
	fmt.Printf("array update after values is:%v,address :%p\n", array6, &array6)
	// 传入数组的引用
	update2(&array6)
	fmt.Printf("array update after values is:%v,address :%p\n", array6, &array6)
}

func update(array [3]int) {
	fmt.Printf("arryay address : %p\n", &array)
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array)
}

func update2(array *[3]int) {
	fmt.Printf("arryay address : %p\n", &array)
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array)
}
