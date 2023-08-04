/*

function return tyoes
go routines
error handling
*/

package main

import "fmt"

func main(){
	fmt.Println("Wellcome to golang")
	var one,two,three = 1, "two", 3.0
	//fmt.Printf("one - %d, tow - %s, three - %f\n", one, two, three)
	fmt.Printf("one - %T, tow - %T, three - %T \n", one, two, three)

	//variable declaraton 1 - dec and assig seperately 
	var a int
	a=1

	//dec 2 - dec and assig same line
	var b int = 2

	//dec 3 - without declarng data type
	var c = 3

	//dec 4 - using shorthand ((without using var)
	d := 4

	//dec 5 - modified dec 1

	var(
		firstName, lastName string
		age int
		salary float64
		isConfirmed bool
	)

	firstName = "Jhon"
	lastName = "s"
	age = 12
	salary = 130000.3
	isConfirmed = true

	fmt.Printf("firstname: %s, lastname: %s, age: %d, salary: %f, isconfrimed: %t \n", firstName,lastName,age,salary, isConfirmed)

}