package main

import "fmt"

func main() {
	var nums [5]int
	//打印数组
	printArray(nums)
	//与其他语言不同的是数组是值类型
	f1(nums)
	printArray(nums)
	f2(&nums)
	printArray(nums)
	//使用make创建
	ints := make([]int, 5)
	fmt.Printf("len(ints) = %d,cap(ints) = %d\n",len(ints),cap(ints))
	ints = make([]int,5,10)
	fmt.Printf("len(ints) = %d,cap(ints) = %d\n",len(ints),cap(ints))
}

func printArray(nums [5]int) {
	fmt.Println("=================")
	for i :=0;i<len(nums);i++{
		fmt.Printf("nums[%d] is : %d\n",i,nums[i])
	}
}

func f1(nums [5]int) {
	nums[0] = 100
}

func f2(nums *[5]int) {
	nums[0] = 100
}
