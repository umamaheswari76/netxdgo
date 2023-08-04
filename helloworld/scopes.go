// local
// global
// lexical or closure --> (scope in fnctoins inside functions)

//

//CLOSURE
//can assign anonmous func to var
//closure can accces the outer functions data
//inner fun have the enclosure

//client - server eg
//capability of spinning of wenserver

package main

import "fmt"

func main(){
	x:=sum(10,20)  // here x is of func type
	// x:=sum(10,20) ()  // here x is of int type
	y:=x()
	fmt.Println(y)
}

func sum(a, b int) func()int{   //2nd fun is anonymouys fun with int type return type
	c:=100
	return func() int{
		return a+b+c
	}

}