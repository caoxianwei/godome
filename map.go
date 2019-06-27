package main

import (
	"fmt"
	"sort"
)

func main()  {
	m :=map[int]string{1:"a",2:"b",3:"c",4:"d"}

	s := make([]int, len(m))
	i:=0
	for k,_ := range m{
		s[i] =k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)

	for i :=0; i<3;i++{
		defer func() { //闭包
			fmt.Println(i)
		}()
	}
}