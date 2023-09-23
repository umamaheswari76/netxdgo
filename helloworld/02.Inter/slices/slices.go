package main

import "fmt"

func main(){
	//slices have length and size
	//it is a special DS that can be expanded and shrink 
	//slice is a dynamically-sized, segmented view of an underlying array
	//this segment can be entire array or a subset of an array
	//the subbset of an array can be indicated by the start and end index
	//slices provide a dynamic window of the underlying array
	
	//max size - cap
	//current size - len
	var a = [5]string{"alp","bet","cat","dog","egg"}

	//creating a slice
	var s []string = a[1:4]
	// fmt.Println("Array a = ", a)
	// fmt.Println("Slice s = ",s)
	slice1 := a[1:4]
	slice2 := a[:3]
	slice3 := a[2:]
	slice := a[:]

	fmt.Println("slice1 = ")

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}