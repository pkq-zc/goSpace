package main

import (
	"log"
	"os"
)

// 创建文件
func main() {
	file, err := os.Create("test.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file)
}
