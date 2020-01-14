package main

import (
	"log"
	"os"
)

// 重命名文件
func main() {
	oldFileName := "test.txt"
	newFileName := "test1.txt"
	err := os.Rename(oldFileName, newFileName)
	if err != nil {
		log.Fatal(err)
	}
}
