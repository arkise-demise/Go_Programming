package main

import (
	"fmt"
	"time"
)


func nonBlockingRead(ch1, ch2 <-chan int) int {
	select {
	case val := <-ch1:
		return val
	case val := <-ch2:
		return val
 
	}
}


func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 2
	}()

	result := nonBlockingRead(ch1, ch2)
	fmt.Println("First value received:", result)
}

