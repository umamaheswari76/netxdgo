package main

import "fmt"

type Node struct{
	Data int
	Next *Node
}

var Front *Node = nil
var Rear *Node = nil

func Create(element int) *Node{
	// n := Node{element,nil}
	// temp_ptr = &n
	n := new(Node)
	n.Data = element
	n.Next = nil
	return n
}

func Add(element int){
	
	if Front==nil{
		temp := Create(element)
		Front = temp
		Rear = temp
		// fmt.Println(front.data)
		// fmt.Println(rear)

	}else{
		// fmt.Println("not nil")
		temp := Create(element)
		Rear.Next = temp
		Rear = temp
		// fmt.Println(rear)

	}
}

func Print_queue(){
	for i:=Front; i!=nil; {
		temp := i
		fmt.Println(temp.Data)
		i = temp.Next
	}
}

func Delete() {
	temp := Front
	Front = temp.Next
	temp = nil

}

func main(){
	Add(1)
	Add(2)
	Add(3)
	Add(4)
	Add(5)
	Delete()
	Delete()
	Add(6)
	Print_queue()

}