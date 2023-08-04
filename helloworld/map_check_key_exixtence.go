//map_check_key_exixtence.go

package main

import "fmt"

func main(){
	var personAge = map[string]int{
		"Ram" : 25,
		"Raj" : 26,
		"Sarah" : 29,
	}

	personAge["sam"]=34
	x,isExisting := personAge["kiran"];
	fmt.Println(x)
	fmt.Println(isExisting)

	for name,age := range personAge{
		fmt.Println(name,age)
	}
}
