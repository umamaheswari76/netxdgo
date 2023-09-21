package main

import (
	"fmt"
	//"encoding/json"
	//"fmt"
	//	"reflect"
	//"reflect"
)

// type Person struct {
// 	first string
// 	last  string
// 	place string
// 	age   int
// }

// 1.
// func main(){
//    var p Person
//    fmt.Println(p)
// }

// a. nil,nil,nil,nil
// b.  , , , ,0
// c. nil,nil,nil,0

// //2.
// func main(){
// 	person1 := Person{first:"Akshay"}
// 	person1 := Person{"Akshay","kumar","chennai",12}
// 	fmt.Println(person1)
// }
// a. Akshay
// b. Akshay  0
// c. error

// //3.
// type Salary struct {
// 	Basic, HRA float64
// }

// type Employee struct {
// 	FirstName, LastName, Email string
// 	Age                        int
// 	MonthlySalary              []Salary
// }

// func main() {
// 	e := Employee{
// 		FirstName: "Mark",
// 		LastName:  "Jones",
// 		Email:     "mark@gmail.com",
// 		Age:       25,

// 		//assign two vales for MonthlySalary field
// 	}
// }

// a.MonthlySalary : {{15000.00,5000.00,},{16000.00, 5000.00,},}

// b.MonthlySalary : {
// 	{
// 		Basic: 15000.00,
// 		HRA: 5000.00,
// 	},
// 	{
// 		Basic: 16000.00,
// 		HRA:   5000.00,
// 	},}

// c.MonthlySalary : {
// 	Salary{
// 		Basic: 15000.00,
// 		HRA:5000.00,
// 	},
// 	Salary{
// 		Basic: 16000.00,
// 		HRA:   5000.00,
// 	},
// }
// D. error

// // 4.
// type Employee struct {
// 	FirstName string `json:"firstname"`
// 	LastName  string `json:"lastname"`
// 	City      string `json:"city"`
// }

// func main() {
// 	json_string := `
//     {
//         "firstname": "Rocky",
//         "lastname": "Sting",
//         "city": "London"
//     }`

// 	emp1 := new(Employee)
// 	json.Unmarshal([]byte(json_string), emp1)
// 	fmt.Println(emp1)
// }

// a.&{firstname lastname city}
// b.<nil>
// c.&{Rocky Sting London}
// d.{}

//5.
// func main(){
// 	emp2 := new(Employee)
//     emp2.FirstName = "Ramesh"
//     emp2.LastName = "Soni"
//     emp2.City = "Mumbai"
//     jsonStr, _ := json.Marshal(emp2)
//     fmt.Printf("%s\n", jsonStr)
// }

// a. {"FirstName":"Ramesh","LastName":"Soni","City":"Mumbai"}
// b. error
// c. {"firstname":"Ramesh","lastname":"Soni","city":"Mumbai"}
// d. {}

// 6.
// type Employee struct {
// 	Name string
// 	Age  int
// }

// func (obj *Employee) Info() {
// 	if obj.Name == "" {
// 		obj.Name = "John Doe"
// 	}
// 	if obj.Age == 0 {
// 		obj.Age = 25
// 	}
// }

// func main() {
// 	emp1 := Employee{Name: "Mr. Fred"}
// 	emp1.Info()
// 	fmt.Println(emp1)
// }

// a. {Mr. Fred 0}
// b. {John Doe 25}
// c. {Mr. Fred 25}
// d. { 25}

//7.
// type rectangle struct {
// 	length  float64
// 	breadth float64
// 	color   string
// }

// func main() {
// 	var rect1 = rectangle{10, 20, "Green"}
// 	fmt.Println(reflect.TypeOf(rect1))
// 	fmt.Println(reflect.ValueOf(rect1).Kind())
// }

// a.  rectangle
// 	interface
// b.  main.rectangle
// 	int
// c.  main.rectangle
// 	struct
// d.  main.rectangle
// 	interface

//8.
// func main(){
// 	rect3 := new(rectangle)
// 	fmt.Println(reflect.TypeOf(rect3))         // *main.rectangle
// 	fmt.Println(reflect.ValueOf(rect3).Kind())
// }
// a.  *main.rectangle
// 	interface
// b.  &main.rectangle
// 	struct
// c.  &main.rectangle
// 	ptr
// d.  *main.rectangle
// 	ptr

//9.
type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {
	var rect1 = rectangle{10, 20, "Green"}
	rect2 := rectangle{length: 10, breadth: 20, color: "Red"}

	if rect1 == rect2 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

// a. True
// b. error
// c. nil
// d. False

// 10.
// func main() {
// 	r1 := rectangle{10, 20, "Green"}
// 	r2 := r1
// 	r2.color = "Pink"
// 	r3 := &r1
// 	r3.color = "Red"
// 	fmt.Println(r1)
// }

// a. {10 20 Green}
// b. {10 20 Red}
// c. &{10 20 Red}
// d. {10 20 Pink}
