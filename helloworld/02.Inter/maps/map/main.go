//map unordered key value pairs
//make - used to create the data structure
package main

import "fmt"

func main(){
	//var m map[string]int
	//fmt.Println(m)

	m := make(map[string] int)
	m["salary"] = 2000
	m["age"] = 20
	m["employeeId"] = 1
	m["xyz"] =123
	m["abc"] = 234

	if m == nil{
		fmt.Println("m is nil")
	}else{
		fmt.Println(m)
	}
	//output
	//map[abc:234 age:20 employeeId:1 salary:2000 xyz:123]


	//the following statement will result in a runtime error
	//m["key"] = 100
} 