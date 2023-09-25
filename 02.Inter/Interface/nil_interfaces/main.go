package main

import (
	"fmt"
	"reflect"
)

// If the concrete value inside the interface itself is nil,
// the method will be called with a nil receiver.

type Abser interface {
	Abs()
}

type Vertex struct {
	x, y int
}

func (v *Vertex) Abs() {
	//checks if the concrete(Vertex) type is nill(doesn't contains any values)
	if v == nil{
		fmt.Println("<nil>")
		return
	}
	v.x++
	v.y++
	fmt.Println(v)
}

func describe(a Abser) {
	fmt.Println(a, reflect.TypeOf(a))
}

func main() {
	var a Abser
	v := &Vertex{10, 20}
	a = v //this can be replaced as, a := &Vertex{10,20}
	a.Abs()
	describe(a)
	//here the above a interface has a value v of Vertex type.
	//this vertex type is called concrete type

	var v1 *Vertex
	a = v1
	a.Abs()
	//here the a interface, only has the concrete type,
	//it doesn't has the value
	//So, if the concrete value inside the interface itself is nil,
	//the method will be called with a nil receiver.
	describe(a)

	a.Abs()
	// A nil interface value holds neither value nor concrete type.
	// Calling a method on a nil interface is a run-time error
	// because there is no type inside the
	// interface tuple to indicate which concrete method to call.
	
}
