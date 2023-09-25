package main

import "fmt"

//Different number of parameters but of the same type
//Variadic Function – A Variadic Function is a function that
//accepts a variable number of arguments

func add(a ...int)int{
	sum := 0
	for _,v := range a{
		sum+=v
	}
	return sum
}

func main(){
	fmt.Println(add(1, 2))
    fmt.Println(add(1, 2, 3))
    fmt.Println(add(1, 2, 3, 4))
}