package main

import (
	"fmt"
	"log"
	"os"
)

// 获取文件信息
func main() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fileName:", fileInfo.Name())
	fmt.Println("fileSize:", fileInfo.Size())
	fmt.Println("permissions:", fileInfo.Mode())
	fmt.Println("last modified:", fileInfo.ModTime())
	fmt.Println("is directory:", fileInfo.IsDir())
	fmt.Printf("system interface type: %T\n", fileInfo.Sys())
	fmt.Printf("system info:%+v\n", fileInfo.Sys())
}
