package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "I Love中国"
	// 获取字符的长度
	fmt.Printf("len(str) = %d\n", len(str))
	// 获取第1个字节所代表的字符
	fmt.Printf("str[0] = %c\n", str[0])
	// str[len(str)] 超过下标将导致异常
	// 通过s[i:j]方式获取子串
	fmt.Printf("str[0:7] = %s\n", str[0:6])
	// 不能通过s[i] = 'a' 这种方式改变字符串的内容
	// 使用utf8.RuneCountInString()计算字符个数
	fmt.Printf("utf8.RuneCountInString(str) = %d\n", utf8.RuneCountInString(str))
	for index, char := range str {
		fmt.Printf("index = %d, char = %c\n", index, char)
	}

	fmt.Printf("comma(\"1234567890\") = %s\n", comma("1234567890"))
}

// 接收一个整数字符串参数,从右开始,每三位数就插入一个逗号
func comma(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}
	return comma(s[:l-3]) + "," + comma(s[l-3:])
}
