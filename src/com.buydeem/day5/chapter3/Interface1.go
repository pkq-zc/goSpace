package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	high, width float32
}

func (re *Rectangle) Area() float32 {
	return re.width * re.high
}

func main() {
	sq := &Square{side: 5}
	re := &Rectangle{5, 4}
	shapers := [...]Shaper{sq, re}
	for _, v := range shapers {
		fmt.Println(v.Area())
	}

	s1 := shapers[0]

	if t, ok := s1.(*Square); ok {
		fmt.Printf("The type of s1 is: %T\n", t)
	}

	if t, ok := s1.(*Rectangle); ok {
		fmt.Printf("The type of s1 is: %T\n", t)
	}
}
