package main

import (
	"fmt"
)

var sss = "123"


const (
	a = 'A'
	b = iota
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	dd := [20]int{19: 1}
	fmt.Println(dd)
	for i := 0; i < 3; i++ {
		v := i+1
		fmt.Println(v)
	}

	bb:=[10]int{}
	fmt.Println(bb)

	s1 := make([]int,3,10)
	fmt.Println(s1, len(s1), cap(s1))
}
