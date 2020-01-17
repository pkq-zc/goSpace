package main

import "fmt"

// 类型声明
type Celsius float64

// 声明摄氏温度类型
type Fahrenheit float64

// 声明多个常量
const (
	//绝对零度
	AbsoluteZeroC Celsius = -273.15
	//结冰温度
	FreezingC Celsius = 0
	//沸腾温度
	BoilingC Celsius = 100
)

func main() {
	fmt.Printf("%g\n", BoilingC)
}

// 摄氏温度转华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// 华氏温度转摄氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
