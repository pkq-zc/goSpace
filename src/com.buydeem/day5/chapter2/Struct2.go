package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamePoint struct {
	Point
	name string
}

// 覆盖子类的Abs
func (p *NamePoint) Abs() float64 {
	return p.Point.Abs() * 100
}

func main() {
	n := &NamePoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())
}
