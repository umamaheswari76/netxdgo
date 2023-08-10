package main

import "fmt"

func myfunc(channel chan string){
	for i:=0;i<50;i++{
		channel<- "Go lang is awesome "
	}
	close(channel)
}
func main(){
	c := make(chan string,100)
	go myfunc(c)
	//myfunc id acting as the data publisher
	//fmt.Println("Capacity of the channel is", cap(c))
	//why sleep time not inclded - here the main func(res,ok) act as receiver, so there is no delay
	counter := 0
	for{
		res, ok := <-c //acts as receiver
		counter++
		//if channel is open ok will be true else false
		if !ok{
			fmt.Println("Channel is closed",ok)
			break
		}
		fmt.Println(counter)
		fmt.Println("Channel is open and data is", res, ok)
	}

}