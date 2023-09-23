// 3 ways of struct def
 package main

 import "fmt"

 type Person struct{
	first string
	last string
	age int
 }

 func main(){
	var p Person
	fmt.Println(p)
 }