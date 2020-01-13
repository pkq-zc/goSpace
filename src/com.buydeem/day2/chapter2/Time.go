package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%02d/%02d/%02d\n",t.Year(),t.Month(),t.Day())
	fmt.Println(t)
	t = time.Now().UTC()
	fmt.Println(t)
	// 天数+1
	day := 60 * 60 * 24 * 1 * 1e9
	t = t.Add(time.Duration(day))
	fmt.Println(t)

}
