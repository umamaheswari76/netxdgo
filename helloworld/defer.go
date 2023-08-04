//defer statements executes last
//non-defer statements executes first and defer statemnet executes after tht

package main

import "fmt"

func main(){
	// fmt.Println("I am first")
	// defer fmt.Println("Predict me1")
	// defer fmt.Println("Predict me2")
	// fmt.Println("I am the end")

	 fun1()
//defer statemnet exectes at last...
//if there is more han one defer statement, then it will execute in the form  of stack

	//output
	//I am first
// I am the end
// Predict me1
}

func fun1(){
	defer fun2()
	 fmt.Println("fun1")
	
}

func fun2(){
	defer fun3()
	 fmt.Println("fun2")
	
}

func fun3(){
	 fmt.Println("fun3")
	
}