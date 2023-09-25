package main

import (
	"fmt"
	"math"
)

//refer this link
//https://www.bogotobogo.com/GoLang/GoLang_Interfaces.php
type Shape interface{
	Area() float64
	Peri() float64
}
type Shape2 interface{
	Shape
	Additional()
}

type Circle struct{
	 Radius float64
}

type Rect struct{
	Length float64
	Breadth float64
}

func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Peri() float64{
	return 2*math.Pi*c.Radius
}

func (c Circle) Diameter() float64{
	return c.Radius/2.0
}

func (c Circle) Additional(){
	fmt.Println("This is an additional func")
}
func (r Rect) Area() float64{
	return r.Length*r.Breadth
}

func (r Rect) Peri() float64{
	return 2 * r.Length * r.Breadth
}

func (r Rect) Additional(){
	fmt.Println("This is an additional func")
}

func main(){
	var s =  Circle{10.0}
	fmt.Println(s.Area())
	var r Shape2 = Rect{10.0, 20.0}
	fmt.Println(r.Area())
	fmt.Println(r.Additional())
}

// db.createUser({	user: "netxd",pwd: "netxd",roles:[{role: "userAdmin" , db:"netxd"}]})

