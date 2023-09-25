package main

import "fmt"

type Vertex struct{
	x, y float64
}

func (v *Vertex) Scale(){
	v.x++
	v.y++
}

func ScaleFunc(v *Vertex){
	v.x = v.x + 2
	v.y = v.y + 2
}

func main(){
	p := &Vertex{100,200}
	p.Scale()
	fmt.Println(p)

	p1 := Vertex{10,20}
	ScaleFunc(&p1)
	fmt.Println(p1)
}

//note:
// functions with a pointer argument must take a pointer:

// var v Vertex
// ScaleFunc(v, 5)  // Compile error!
// ScaleFunc(&v, 5) // OK
// while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

// var v Vertex
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK
// if the receiver method is in pointer type(v *vector),
//then while calling the method will change the values
// and if the receiver method is in value type(v vertex),
//then while calling the method won't change the values

//but it is better to pass the valid type to the receiver
//method, because it may cause error while using interfaces
// refer interface_methods folder