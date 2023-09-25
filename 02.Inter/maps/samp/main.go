package main

import "fmt"

//import "fmt"

// func main(){
// 	var m map[int]string
// 	var m1 map[int]string
// 	m1 = make(map[int]string)
// 	m = make(map[int]string)
// 	m[0] = "a"
// 	m[1] = "b"
// 	m1 = m
// 	m1[0]="xyz"

// 	for key,val := range m{
// 		fmt.Println(key," ",val)
// 	}
	
// 	temp, ok := m[1]
// 	fmt.Println("val: ",temp,ok)
// 	fmt.Println(m)
//  }

// type People struct{
// 	Name string
// 	Likes []string
// }

// var p []People = []People{}

// func main(){
// 	var people []*People
// 	likes := make(map[string]*People)



// 	for _,p := range people{
// 		for _, l := range p.Likes{
// 			likes[l] = p
// 		}
// 	}

// }


// type samp struct{
// 	name string
// }

// func main(){
// 	t := samp{name: "uma"}
// 	fmt.Println(t)
// }




func main(){

temp := struct{
	 string
}{ "uma"}
fmt.Println(temp.name)
}