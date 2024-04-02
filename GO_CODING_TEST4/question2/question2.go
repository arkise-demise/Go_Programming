package main

import (
	"fmt"
)

func Incrementer()func() int{
	count :=0

	return func () int  {
		count++
		return count
	}
}
 func main(){
	increment :=Incrementer()

	fmt.Println("Incremented values:")
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())


 }