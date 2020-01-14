package main

import (
	"fmt"
	"reflect"
	"strings"
)

type struct1 struct {
	i1   int
	f1   float64
	str1 string
}

type Person struct {
	firstName string
	lastName  string
}

// 带标签的结构体
type Student struct {
	name string "student user name"
	age  int    "student age"
}

//  匿名字段和内嵌结构体
type innerStruct struct {
	innerNumber1 int
	innerNumber2 int
}

type outerStruct struct {
	outerNumber3 int
	outerNumber4 int
	string
	innerStruct
}

type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1
}

func (b B) write() string {
	return fmt.Sprint(b)
}

func main() {
	s1 := new(struct1)
	s1.i1 = 1
	s1.f1 = 2.0
	s1.str1 = "hello world"
	fmt.Printf("i1 = %d\n", s1.i1)
	fmt.Printf("f1 = %f\n", s1.f1)
	fmt.Printf("str1 = %s\n", s1.str1)

	//
	var p1 Person
	p1.firstName = "mac"
	p1.lastName = "jackson"
	upPerson(&p1)
	fmt.Println(p1)

	//
	student := Student{"tom", 18}
	studentType := reflect.TypeOf(student)
	nameField := studentType.Field(0)
	fmt.Printf("%v\n", nameField.Tag)

	//
	outer := new(outerStruct)
	outer.outerNumber3 = 3
	outer.outerNumber4 = 4
	outer.innerNumber1 = 1
	outer.innerNumber2 = 2
	outer.string = "hello world"
	fmt.Println(outer)

	//
	var b1 B
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B)
	b2.change()
	fmt.Println(b2.write())
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
