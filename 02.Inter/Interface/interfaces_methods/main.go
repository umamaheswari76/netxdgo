package main

import "fmt"

type Abser interface{
	Abs() 
}

type Vertex struct{
	x, y int
}

type Myint int

func (v *Vertex) Abs(){
	v.x++
	v.y++
	fmt.Println(v)
}

func (m Myint) Abs(){
	m++
	fmt.Println(m)
}

func main(){
	var a Abser

	//anyone is valid
	v := &Vertex{10,20} 
	//a = &v

	//both of these are invalid
	//v1 := Vertex{10,20} 
	//a = v1
	// because cannot use v1 (variable of type Vertex) as Abser value in assignment:
	// Vertex does not implement Abser (method Abs has pointer receiver)compilerInvalidIfaceAssign
	m := Myint(30)
	a = v
	a.Abs()
	//here the abs func is referring to v

	a = m
	a.Abs()
	//here the abs func is refering to m
	//like run time polymorphism

	// v.Abs()
	// m.Abs()
}

//Runtime Polymorphism in Go :
// In Runtime Polymorphism, the method is resolved at runtime.
// In Golang, runtime polymorphism is implemented using Interfaces.

// Interfaces in golang are quite different from other languages like Java where a class has to explicitly state that it implements an interface using the implements keyword.
// Go interfaces are implemented implicitly if a type provides the definition for all the methods declared in the interface.

// A variable of an interface type can contain value of any type which implements the interface. This is the property that helps in achieving polymorphism in the Go language

// Objects of different types are treated in a consistent way, as long as they stick to a single interface, which is the essence of polymorphism.


//refer this for runtime polymorphism on golang
//https://sagarsonwane230797.medium.com/understanding-polymorphism-in-go-d704944e9816#:~:text=Runtime%20Polymorphism%3A%20It%20is%20a,is%20achieved%20by%20Method%20Overriding.
