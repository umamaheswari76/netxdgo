package main

import "fmt"

func main(){
	//deleting a key from a map using builtin delete() function

	var fileExtensions = map[string] int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
	}

	fmt.Println(fileExtensions)
	delete(fileExtensions, "three")
	fmt.Println(fileExtensions)
}