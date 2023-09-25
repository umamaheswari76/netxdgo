package main

import "fmt"

func in(){}

// func main(){
// 	func ()(){
// 		fmt.Println("hi")
// 	}() //anonymous function is called here
// }

// we can also assign anonymous function to a variable,
//and the function call is done by the variable itself
// func main(){
// 	val := func (c int){
// 		fmt.Println(c,"hi")
// 	}
// 	val(45)
// 	// fmt.Println(val)
// }

//we can also pass anonymous func as an argument 
//to a normal function
// func normal(a func()){
// 	a()
// }
// func main(){
// 	val := func(){
// 		fmt.Println("anony")
// 	}
// 	normal(val)
// }

//we can also return anonymous func from normal function
// func normal()(func()){
// 	val := func(){
// 		fmt.Println("hi")
// 	}
// 	return val
// }
// func main(){
// 	temp := normal()
// 	// fmt.Println(temp)
// 	temp()
// }

//returning anonymous functions with values
// func argu()(func(int, int) int){
// 	val := func(c,d int)(int){
// 		return c+d
// 	}
// 	return val

// 	//returning ike this also fine(straightly returning, instead of assigning to var)
// 	// return func(c,d int)(int){
// 	// 	return c+d
// 	// }
// }

// func main(){
// 	temp := argu()
// 	fmt.Println(temp(2,5))
// }


// A closure is a special type of anonymous 
//function that references variables declared outside of the function itself.
func main(){
	a := 0
	counter := func()int{
		a+=1
		return a
	}
	fmt.Println(counter())
	fmt.Println(counter())

	// val1 := func()int{
	// 	a++
	// 	return a
	// }
	// fmt.Println(val1())
}
//there is a slight problem as any other function which will be defined in the main has access to the global 
//variable GFG and it can change it without calling counter function. 
//Thus closure also provides another aspect which is data isolation.

