package main

import (
	"fmt"
	"time"
)

func fib(n int){
	start := time.Now()
	val1 := 0
	val2 := 1
	fmt.Println(val1)
	fmt.Println(val2)
	for i:=2;i<n;i++ {
		temp := val1+val2
		val1 = val2
		val2 = temp
		fmt.Println(temp)
	}
	end := time.Now().Sub(start)
	fmt.Println(end)
}
 
func main(){
	n := 1000
	fib(n)
}