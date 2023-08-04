package main

import "fmt"

func main(){
	var personAge = map[string]int{
		"Ram" : 25,
		"Raj" : 26,
		"Sarah" : 29,
	}

	personAge["sam"]=34

	for name,age := range personAge{
		fmt.Println(name,age)
	}
}