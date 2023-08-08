package main

 import "fmt"

 type Student struct{
	Rollnumber int
	Name string
 }

 
func main(){
	s := Student{11,"jack"}
	ps := &s
	fmt.Println(ps)
	fmt.Println((*ps).Name)
	fmt.Println(ps.Name) //same as above : no need for explicit dereference

	ps.Rollnumber =31
	fmt.Println(ps)
 }