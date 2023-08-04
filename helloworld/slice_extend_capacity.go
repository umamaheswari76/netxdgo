package main

import "fmt"

func main(){
	s := []int{10,20,30,40,50,60,70,80,90,100}
	fmt.Println("Orginal array")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
	

	s = s[:8]
	fmt.Println("After extending the len")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
	

	s = s[2:]
	fmt.Println("After dropping the first two elements")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
	
	// output
// Orginal array
// s = [10 20 30 40 50 60 70 80 90 100], len = 10, cap = 10
// After extending the len
// s = [10 20 30 40 50 60 70 80], len = 8, cap = 10
// After dropping the first two elements
// s = [30 40 50 60 70 80], len = 6, cap = 8

}