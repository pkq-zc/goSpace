package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// 创建一个新的错误
	notFoundError := errors.New("Not Found Error")
	if e, ok := notFoundError.(*os.PathError); ok {
		fmt.Println(e)
	} else {
		fmt.Printf("error : %v\n", notFoundError)
	}
}
