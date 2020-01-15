package app1

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")
	go logWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func logWait() {
	fmt.Println("Beginning logWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("End of logWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("End of shortWait()")
}
