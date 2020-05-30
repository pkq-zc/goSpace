package main

import "fmt"

func main() {
	s, j, k := f1(1, 2)
	fmt.Printf("s = %d,j = %d,k = %d\n",s,j,k)
	s, j, k = f2(1, 2)
	fmt.Printf("s = %d,j = %d,k = %d\n",s,j,k)


}

func f1(a int, b int) (s, j, k int) {
	return a+b,a-b,a*b
}

func f2(a, b int) (int, int, int) {
	return a+b,a-b,a*b
}

