package main

import "fmt"

func main() {
	// 把函数赋给f
	f := sayHello
	fmt.Printf("f type is :%T\n",f)
	f()

	// 加法
	add := Add
	result := Operation(add, 1, 2)
	fmt.Printf("result : %d\n",result)
	// 减法
	sub := Sub
	result = Operation(sub,1,2)
	fmt.Printf("result : %d\n",result)

	//乘法
	mul := func(x,y int) int {return x*y}
	result = Operation(mul, 1, 2)
	fmt.Printf("result : %d\n",result)


	//计数器使用，获取一个初始值为100的计数器
	c1 := counter(100)
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())
	//计数器使用，获取一个初始值为0的计数器
	c2 := counter(0)
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())


	// 可变参数
	total := Sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	values := []int{1,2,3,4,5,6,7,8,9}
	Sum(values...)
	fmt.Println(total)

	 total = Sum()
	 fmt.Println(total)

	 //defer
	 fmt.Println(f6())

	 f9()

	 fmt.Printf("f10() = %d\n",f10())
	 fmt.Printf("f11() = %d\n",f11())
}

func sayHello() {
	fmt.Println("Hello")
}

// op 行参是一个函数，可以传入一个动作
func Operation(op func(x, y int) int, x, y int) int {
	return op(x,y)
}

// 加法
func Add(x, y int) int {
	return x+y
}
// 减法
func Sub(x, y int) int {
	return x-y
}

// 计数器
func counter(initValue int) func() int {
	return func() int {
		initValue++
		return initValue
	}
}

//可变参数
func Sum(values ...int) int {
	total := 0
	for _, value := range values {
		total+= value
	}
	return total
}

func f6() string {
	defer f7("defer execute")
	return "return over"
}

func f7(s string) {
	fmt.Println(s)
}

// 执行顺序倒序
func f8() {
	defer fmt.Println(4)
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
}

// 被defer的参数在defer时确定
func f9() {
	i := 0
	defer fmt.Printf("i address :%p,value : %d\n",&i,i)
	i++
}

func f10() int {
	i := 0
	defer func() {i++}()
	return i
}

func f11() (i int) {
	defer func() {i++}()
	return i
}