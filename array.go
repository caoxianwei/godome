package main

import "fmt"

func main() {
	y := make([]int, 3,5)
	y = append(y, 5,6,7)

	fmt.Printf("%v",y)
}
