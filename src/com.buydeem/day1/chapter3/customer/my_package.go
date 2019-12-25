package customer

import "fmt"

// 使用开头字母来标识变量是否能在外包使用
var notCanSeeInOtherPackage = "在其他包中不能访问"
var CanSeeInOtherPackage = "在其他包中可见"

// 使用开头字母来标识方法是否能在外包使用
func SayHello()  {
	fmt.Println("Hello")
}

// 使用开头字母来标识类型是否能在外包使用
type News string
