package chapter1

import "fmt"

const i  = 5

// 指针
func main() {
	var i1 = 5
	fmt.Printf("value is %d, point address %p\n",i1,&i1)

	var intP *int = &i1
	fmt.Printf("value is %d, point address %p\n",*intP,intP)

	str := "hello china"
	var p *string = &str
	*p = "hello world"
	fmt.Printf("str value is : %s\n",str)
	fmt.Printf("*p value is :%s\n",*p)

	// 不能获取常量的指针
	//ptr := &i

	var p1 *int = nil
	*p1 = 0
	fmt.Printf("%d",*p1)
}
