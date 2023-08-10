package main

import (
	"fmt"
	"time")

func fibonacci(n int, c chan int){
	x,y := 0,1
	for i:=0; i<n; i++{
		c <- x
		x,y = y,x+y
	}
	close(c)
}

func main(){
	start:=time.Now()
	c := make(chan int,10)//buffered type channel
	go fibonacci(cap(c), c)
	for i:=range c{
		fmt.Println(i)
	}
	end := time.Now().Sub(start)
	fmt.Println(end)
}