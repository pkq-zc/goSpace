package main

import "fmt"

const (
	width  = 1920
	height = 1080
)

type pixel int

var screen [width][height]pixel

// 数组长度固定
// 长度不同的数组不是同一种类型
// 可以通过下标索引读取和修改数组中的值
func main() {
	var names [3]string
	for index, value := range names {
		fmt.Printf("index = %d,value = %s\n", index, value)
	}

	var numbers = [...]int{6, 8, 9, 1, 3}
	for index, value := range numbers {
		fmt.Printf("index = %d,value = %d\n", index, value)
	}

	var years = new([3]string)
	fmt.Printf("years type is :%T,names type is :%T\n", years, names)

	//更新数组的值
	names[0] = "tom"
	fmt.Printf("names address :%p\n", &names)
	update1(names)
	fmt.Println(names)
	update2(&names)
	fmt.Println(names)

	//
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			screen[x][y] = 0
		}
	}
}

func update1(names [3]string) {
	fmt.Printf("names address :%p\n", &names)
	names[0] = "mac"
}

func update2(names *[3]string) {
	fmt.Printf("names address :%p\n", names)
	names[0] = "mac"
}
