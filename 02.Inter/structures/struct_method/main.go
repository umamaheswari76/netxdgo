package main

import (
	"fmt"
	"strconv"
)

// import "fmt"

//refer this link
//https://www.bogotobogo.com/GoLang/GoLang_Structs.php#:~:
//text=A%20struct%20is%20just%20a,the%20value%20of%20any%20field.

//struct methods in go are declared as,
//func (p Person) hello(){}
//here we have defined a method called (p Point) hello(){}

//It is called receiver  in go terms
//func (receiver) func_name(parameters) return_type{code}

type Person struct{
	name string
	city string
	gender string
	age int
}

func (p Person) hello()string{
	
	return "hello: "+p.name+" from "+p.city+" age is "+strconv.Itoa(p.age)
}

func (p Person) hellobyval(){
	p.age = 43
}

func (p *Person) hellobyptr(){
	p.age = 43
}


func main(){
	p1 := Person{
		name: "sita",
		city: "chennai",
		gender: "female",
		age: 30,
	}

	p2 := Person{
		name:   "laks",
		city:   "chennai",
		gender: "female",
		age:    20,
	}
	fmt.Println(p1.name)
	fmt.Println(p1)
	p1.hellobyval()
	fmt.Println(p1.hello())
	fmt.Println("changed through ref")
	p1.hellobyptr()
	fmt.Println(p1.hello())

	fmt.Println(p2.name)
	fmt.Println(p2)
	fmt.Println(p2.hello())
}