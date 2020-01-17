package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

const (
	width, height = 600, 320            //已像素表示画布的大小
	cells         = 100                 //网格单元的个数
	xyrange       = 30.0                //坐标轴的范围(-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange //x/y轴上每个单位长度的像素
	zscale        = height * 0.4        //z轴上每个单位长度的像素
	angle         = math.Pi / 6         //x/y周的角度 (= 30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	file, err := os.OpenFile("1.svg", os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(file, "<svg xmlns='http://www.w3.org/2000/svg'"+
		" style='stroke:grey;fill:white;stroke-width:0.7'"+
		" width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(file, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(file, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// 求出网格党员(i,j)的顶点坐标(x,y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	//计算曲面高度
	z := f(x, y)
	// 将x,y,z等角投射到二维svg绘图平面上,坐标是(sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}
