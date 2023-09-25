package main

import "fmt"

//"fmt"
//"reflect"

// func main(){
// 	var personAge = map[string]int{
// 		"Ram" : 25,
// 		"Raj" : 26,
// 		"Sarah" : 29,
// 	}

// 	personAge["sam"]=34

// 	for name,age := range personAge{
// 		fmt.Println(name,age)
// 	}
// }

// a. Sarah 29
// sam 34
// Ram 25
// Raj 26

// b. error

// c. Sarah 29
// Ram 25
// Raj 26

// 2.
// func main() {
// 	m := make(map[string]int)
// 	m["k1"] = 7
// 	m["k2"] = 8
// 	fmt.Println("map:", m["k1"])
// }

// a. map: 7
// b. map: "k1":7
// c. 7
// d. "k1": 7

//3.
// func main() {
//     elements := make(map[string]string)
//     elements["O"] = "Oxygen"
//     elements["Ca"] = "Calcium"
//     elements["C"] = "Carbon"
//     fmt.Println(elements)
// }

// a. map[O:Oxygen Ca:Calcium C:Carbon ]
// b. nothing will be printed
// c. maps in go are unordered
// d. O:Oxygen Ca:Calcium C:Carbon

// 4.
//  func main() {
//  	m := make(map[string]int)
//  	m["k1"] = 7
// 	m1 := m
// 	m2 := &m
// 	m1["k1"]= 8
// 	fmt.Println("map:",m1)
// 	(*m2)["k1"]= 9
// 	fmt.Println("map:",m)
// 	//fmt.Println(m1)
//  }

// a. map: map[k1:8]
// 	map: map[k1:9]
// b. map: map[k1:8]
// 	map: map[k1:7]
// c. map: map[k1:7]
// 	map: map[k1:7]
// d. map: map[k1:7]
// 	map: map[k1:9]

// 5.
// type order struct {
// 	ordId      int
// 	customerId int
// }

// func createQuery(q interface{}) {
// 	t := reflect.TypeOf(q)
// 	v := reflect.ValueOf(q)
// 	fmt.Println("Type ", t)
// 	fmt.Println("Value ", v)
// }
// func main() {
// 	o := order{
// 		ordId:      456,
// 		customerId: 56,
// 	}
// 	createQuery(o)

// }

// a. Type  order
// Value  {0 0}
// b. Type  main.order
// Value  {456 56}
// c. Type  order
// Value  {456 56}

// 6.
func main() {
    results := []string{}

    sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
    //fmt.Println(sammy)
    //sammy := map[string]int{"name": 1, "animal": 2, "color": 3, "location": 4}
    for res := range sammy {
    results = append(results, res)
    }
    fmt.Printf("%q", results)
}

// a. error
// b. not print anything
// c. map[animal:shark color:blue location:ocean name:Sammy]
// d. ["Sammy" "shark" "blue" "ocean"]

// 7.
// How to access the name and type values from the following program?
// func  main(){
// 	website := map[string]map[string]string {
// 		"Google": map[string]string {
// 			"name":"Google",
// 			"type":"Search",
// 		},
// 		"YouTube": map[string]string {
// 			"name":"YouTube",
// 			"type":"video",
// 		},
// 	}
// }

// a.  fmt.Println(website["name"])
// 	fmt.Println(website["type"])

// b.	fmt.Println(website."Google".["name"])
// 	fmt.Println(website."Google".["type"])

// c.	fmt.Println(["Google"]["name"])
// 	fmt.Println(["Google"]["type"])

// d.	fmt.Println(website["Google"]["name"])
// 	fmt.Println(website["Google"]["type"])

// 8.
// What this program will do?
// func main() {
//     m := map[int]map[string]string{
//         1: {
//             "a": "Apple",
//             "b": "Banana",
//             "c": "Coconut",
//         },
//         2: {
//             "a": "Tea",
//             "b": "Coffee",
//             "c": "Milk",
//         },
//         3: {
//             "a": "Italian Food",
//             "b": "Indian Food",
//             "c": "Chinese Food",
//         },
//     }
// }

// a. nothing
// b. throws error
// c. creating a map of int as keys and map as values
// d. nested maps are not allowed in go

// 9.
// type Service interface{
// 	SayHi()
// }

// type MyService struct{}
// func (s MyService) SayHi() {
// 	fmt.Println("Hi")
// }

// type SecondService struct{}
// func (s SecondService) SayHi() {
// 	fmt.Println("Hello From the 2nd Service")
// }

// func main() {
// 	interfaceMap := make(map[string]Service)
// 	interfaceMap["SERVICE-ID-1"] = MyService{}
// 	interfaceMap["SERVICE-ID-2"] = SecondService{}
// 	interfaceMap["SERVICE-ID-2"].SayHi()
// }

// a. error
// b. Hello From the 2nd Service
// c. can't use struct insde map in go
// d. Hi
