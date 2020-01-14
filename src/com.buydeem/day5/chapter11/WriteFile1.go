package main

import (
	"log"
	"os"
)

// 写文件
func main() {
	file, err := os.OpenFile("1.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("hello 中国")
	n, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", n)
}
