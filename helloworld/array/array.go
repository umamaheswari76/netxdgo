package main

import "fmt"

func main() {
	arr := [5]int{100,200,300}
	for ind, val := range arr{
		fmt.Println(ind," ",val)
	}
}



//01.

// func main() {
// 	arr := [5]int{100,200,300,400}
// 	arr[4] = 500
// 	arr[5] = 600
// 	for i:=0;i<len(arr);i++{
// 		fmt.Println(arr[i])
// 	}
// }
