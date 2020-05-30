package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	var b bool = false
	fmt.Println(b)

	str := "hello world"
	fmt.Printf("str len() = %d\n",len(str))
	fmt.Printf("str size() = %d\n",utf8.RuneCountInString(str))
	str = "hello 中国"
	fmt.Printf("str len() = %d\n",len(str))
	fmt.Printf("str size() = %d\n",utf8.RuneCountInString(str))

	lStartIndex := strings.Index(str, "l")
	fmt.Printf("l start index = %d\n",lStartIndex)
	lLastIndex := strings.LastIndex(str, "l")
	fmt.Printf("l last index = %d\n",lLastIndex)
	index := strings.Index(str, "国")
	fmt.Printf("国 start index = %d\n",index)
	index2 := strings.IndexRune(str, '国')
	fmt.Printf("国 start index2 = %d\n",index2)

	r1 := strings.Replace("oink oink oink", "ink", "s", 2)
	fmt.Println(r1)
	r1 = strings.Replace("oink oink oink","i","s",-1)
	fmt.Println(r1)


	fmt.Println(time.Now().Format("2006-01-02 03:04:05"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
