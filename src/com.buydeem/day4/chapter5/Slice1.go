package main

import "fmt"

// 引用类型 可以通过下标索引
// 长度可变 通过len()获取长度 通过cap()获取容量 长度永远不会超过容量
// 底层通过数组实现,如果多个切片使用同一数组,数据可以共享
// 切片容量会自动扩充
func main() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(s1) = %d,cap(s1) = %d,s1 is :%v\n", len(s1), cap(s1), s1)
	arrays := [10]int{1, 2, 3, 4, 5}
	s2 := arrays[0:4]
	fmt.Printf("len(s2) = %d,cap(s2) = %d,s2 is :%v\n", len(s2), cap(s2), s2)
	s3 := arrays[0:8]
	fmt.Printf("len(s3) = %d,cap(s3) = %d,s3 is :%v\n", len(s3), cap(s3), s3)
	//改变arrays的值
	arrays[0] = 10
	fmt.Printf("s2 = %v\n", s2)
	fmt.Printf("s3 = %v\n", s3)
	// 改变值
	s4 := []string{"china", "japan", "usa"}
	fmt.Printf("s4 address :%p\n", s4)
	fmt.Printf("update before s4:%v\n", s4)
	update(s4, 0, "中国")
	fmt.Printf("update after s4:%v\n", s4)

	//使用make创建slice 5代表长度,10代表底层数组容量
	s5 := make([]string, 5, 10)
	fmt.Printf("len(s5) = %d,cap(s5) = %d\n", len(s5), cap(s5))
	//
	s6 := make([]byte, 5)
	fmt.Printf("len(s6) = %d,cap(s6) = %d", len(s6), cap(s6))
	// 使用for -range
	s7 := []string{"a", "b", "c", "d"}
	for index, value := range s7 {
		if index%2 == 0 {
			s7[index] = s7[index] + "0"
		} else {
			// value  只是一个拷贝值
			value = value + "1"
		}
	}
	fmt.Printf("s7 = %v\n", s7)

	//
	s8 := []int{1, 3, 5, 0, 10, 8, 7}
	max, min := findMaxAndMin(s8)
	fmt.Printf("max = %d, min = %d\n", max, min)
	// 切片复制
	s9 := []int{1, 2, 3}
	s10 := make([]int, 5)
	copy(s10, s9)
	fmt.Println(s10)
}

func update(s []string, index int, value string) {
	s[index] = value
	fmt.Printf("s address :%p\n", s)
}

// 寻找最大值和最小值
func findMaxAndMin(s []int) (max, min int) {
	if len(s) == 0 {
		return 0, 0
	}
	max = s[0]
	min = s[0]
	for _, value := range s {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	return max, min
}
