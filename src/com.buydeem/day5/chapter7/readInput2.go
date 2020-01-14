package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")
	input, err := reader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
