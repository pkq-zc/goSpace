package main

import (
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("2.txt", []byte("你好 world"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
