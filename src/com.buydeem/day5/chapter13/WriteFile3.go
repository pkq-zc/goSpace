package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 缓存写
func main() {
	file, err := os.OpenFile("3.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建buffer write
	writer := bufio.NewWriter(file)

	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i) + "\n"
		writer.WriteString(s)
	}

	// 检查缓存中的字节数
	buffered := writer.Buffered()
	fmt.Println(buffered)
	// 查看还有多少缓存可用
	available := writer.Available()
	fmt.Println(available)
	//把内存中的数据写到硬盘
	writer.Flush()
}
