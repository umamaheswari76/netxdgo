// Golang program to show how
// to take input from the user
package main
  
import (
	"fmt"
	"strconv")
  
// main function
func main() {
  
    // Println function is used to
    // display output in the next line
    fmt.Println("Enter Your First Name: ")
  
    // var then variable name then variable type
    var first string
  
    // Taking input from user
    fmt.Scanln(&first)
    fmt.Println("Enter Second Last Name: ")
    var second string
    fmt.Scanln(&second)
  
    // Print function is used to
    // display output in the same line
    fmt.Println("Your Full Name is: ",first + " " + second)

	
}