package main

import (
	"fmt"
	"strconv"
	"time"
)

var ch chan int

func main() {
	ch = make(chan int)

	//协程1
	go func() {
		for i := 1; i < 100; i++ {
			if i == 10 {
				//runtime.Gosched()
				<-ch
			}
			fmt.Println("routina 1:" + strconv.Itoa(i))
		}
	}()
	//协程1
	go func() {
		for i := 1; i < 200; i++ {
			println("routina 2:" + strconv.Itoa(i))
		}
	}()

	time.Sleep(time.Second)
}
