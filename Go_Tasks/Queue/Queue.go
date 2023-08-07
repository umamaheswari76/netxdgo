package main

import "fmt"

var front *node = nil
var rear *node = nil

type node struct{
	data int
	next *node
}

func create(element int) *node{
	// n := node{element,nil}
	// temp_ptr = &n
	n := new(node)
	n.data = element
	n.next = nil
	return n
}

func add(element int){
	
	if front==nil{
		temp := create(element)
		front = temp
		rear = temp
		// fmt.Println(front.data)
		// fmt.Println(rear)

	}else{
		// fmt.Println("not nil")
		temp := create(element)
		rear.next = temp
		rear = temp
		// fmt.Println(rear)

	}
}

func print_queue(){
	for i:=front; i!=nil; {
		temp := i
		fmt.Println(temp.data)
		i = temp.next
	}
}

func delete() {
	temp := front
	front = temp.next
	temp = nil

}

func main(){
	add(1)
	add(2)
	add(3)
	add(4)
	add(5)
	delete()
	delete()
	add(6)
	print_queue()

}