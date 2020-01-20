package main

import (
	"fmt"
	"os"
)

func main() {
	s := comma("1234567890")
	fmt.Println(s)

	str, _ := f5()
	fmt.Println(str)

	//错误
	file, err := os.OpenFile("test1.txt", os.O_RDONLY, 0111)
	if err != nil {
		fmt.Println(err)
		fmt.Println("打开文件失败！")
	}
	defer file.Close()
}

// 行参为s 返回字符串
func f1(s string) string {
	return "f1"
}

// 无返回值
func f2(s string) {

}

// 返回r
func f3(s string) (r string) {
	r = "hello world"
	return
}

// t同时返回多个值
func f4() (r1, r2 string) {
	r1 = "hello"
	r2 = "world"
	return
}

func f5() (string, int) {
	return "happy new year",2020
}

func comma(strNumber string) string {
	n := len(strNumber)
	if n <= 3 {
		return strNumber
	}
	return comma(strNumber[0:n-3]) + "," + strNumber[n-3:]
}
