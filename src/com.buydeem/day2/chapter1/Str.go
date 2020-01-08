package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	var ch1 int ='\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	// 获取十进制值
	fmt.Printf("%d - %d - %d \n",ch1,ch2,ch3)
	// 获取unicode值
	fmt.Printf("%c - %c - %c \n",ch1,ch2,ch3)
	// 获取十六进制值
	fmt.Printf("%X - %X - %X \n",ch1,ch2,ch3)
	// 获取unicode格式
	fmt.Printf("%U - %U - %U \n",ch1,ch2,ch3)
	// 判断是否是字母
	fmt.Printf("是否是字母：%t \n",unicode.IsLetter(rune(ch1)))
	// 判断是否是数字
	fmt.Printf("是否是数字：%t \n",unicode.IsDigit(rune(ch2)))
	// 是否是空白符号
	fmt.Printf("是否是空白符：%t \n",unicode.IsSpace(rune(ch3)))

	str := "hello world"
	fmt.Println(str)

	str = "你好china"
	fmt.Printf("字符串长度(len)：%d \n",len(str))
	fmt.Printf("字符串长度(utf8.RuneCountInString):%d \n",utf8.RuneCountInString(str))
	// 匹配前缀
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ",str,"你好")
	fmt.Printf("%t\n",strings.HasPrefix(str,"你好"))
	// 包含
	fmt.Printf("Does tht string \"%s\" contains \"%s\"? ",str,"china")
	fmt.Printf("%t\n",strings.Contains(str,"china"))
	// 判断字符串在字符中的位置
	fmt.Printf("%d\n",strings.LastIndex(str, "china"))
	fmt.Printf("%d\n",strings.IndexRune(str,'好'))
	// 重复输出 n 次
	str1 := strings.Repeat("hello", 3)
	fmt.Printf("%s\n",str1)
	// 替换 n 代表替换多少次 如果n<0 代表替换所有
	str2 := "hello 中国,中国 I love you"
	str3 := strings.Replace(str2,"中国","china",2)
	fmt.Printf("%s\n",str3)
	// 计算出现次数
	count := strings.Count("My name is : zengChao. What's your name ?", "name")
	fmt.Println(count)
	//大小写
	lowerStr := strings.ToLower("HELLO WORLD")
	fmt.Println(lowerStr)
	upperStr := strings.ToUpper(lowerStr)
	fmt.Println(upperStr)
	//去掉开头和结尾字符串中的空格
	str4 := " hello world     "
	fmt.Println(strings.TrimSpace(str4))
	str5 := "-hello word-"
	fmt.Println(strings.TrimLeft(str5,"-"))
	fmt.Println(strings.TrimRight(str5,"-"))
	// 分割字符串 使用空格切分字符串
	fields := strings.Fields("hello word ")
	for i := range fields {
		fmt.Println(fields[i])
	}
	splits := strings.Split("hello world | hello china", "|")
	for i := range splits {
		fmt.Println(splits[i])
	}
	//字符串拼接
	strs := []string{"中国","美国","日本"}
	fmt.Println(strings.Join(strs,";"))
	//其他类型 =>字符串
	str_1 := strconv.Itoa(1)
	fmt.Printf("%T\n",str_1)
	float32Str := strconv.FormatFloat(3.1415926, 'E', -1, 64)
	fmt.Println(float32Str)
	// 字符串 => 其他类型
	num, _ := strconv.Atoi("120")
	fmt.Printf("%T\n",num)
	double, _ := strconv.ParseFloat("3.1415926", 64)
	fmt.Printf("%T %f\n",double,double)
}
