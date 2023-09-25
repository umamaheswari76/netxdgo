// main and go routines will run concurrently
// all go routines will be under main function
//if main execution is completed then the go routine will be killed even it is executing or not


package main

import (
	"fmt"
	"time")

func hello(){
	fmt.Println("I am a go routine")
}

func main(){
	fmt.Println("I am n the main")
	fmt.Println("Calling hello")
	go hello() //prefixing any method with go will br a go routine
	
	//make main to wait so that all other go routines will get executed
	//when main is completed, go routine will be killed
	time.Sleep(10*time.Second)// if it is not given then the main don't wait for go routine to execute and main will be completed
	//so we don't know about the execution condition of go routine
	//for knowing that we are making the main func to wait for some time to make the full execution of go routine
	fmt.Println("Main completed")
}

