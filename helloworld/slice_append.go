package main

import "fmt"

func main(){

	//append() function appends new elemets at the end of a given slice

	//appending to  slice that doesn't have enugh capacity to accomodate new elements
	slice1 := [] string{"Cpp","C","Java"}
	slice2 := append(slice1, "Python", "Ruby", "go")
	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))

	slice2[0]="xyz"
	fmt.Println("\nslize1 = ",slice1)
	fmt.Println("slize2 = ",slice2)
	
	//output
// slice1 = [Cpp C Java], len = 3, cap = 3
// slice2 = [Cpp C Java Python Ruby go], len = 6, cap = 6

// slize1 =  [Cpp C Java]
// slize2 =  [xyz C Java Python Ruby go]

}
