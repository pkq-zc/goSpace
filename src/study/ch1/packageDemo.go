package main

import (
	"fmt"
	"study/ch1/customerunit"
)

func main() {
	// 定义1千米
	var distance customerunit.KM = 1
	cm := customerunit.KM2M(distance)
	fmt.Println(cm)
	//在包外无法访问base
	//fmt.Println(customerunit.base)
	//在包外无法调用kM2CM
	//customerunit.kM2CM(distance)
}
