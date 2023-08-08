package main

import "fmt"


type Node struct{
	Data int
	Prev *Node
	Next *Node
}

var Head *Node = nil
var Tail *Node = nil

func Create(element int) *node{
	n := new(Node)
	n.Data = element
	n.Prev = nil
	n.Next = nil
	return n
}

func Add_begin(element int){
	if Head==nil{
		Front = Create(element)
		Tail = Front
	}else{
		temp := Create(element)
		temp.Next = Front
		Front.Prev = temp.Next
		Front = temp
	}
}

func Add_end(element int){
	 

}

func Add_pos(element, pos int){

}

func Del_begin(){

}

func Del_end(){

}

func Del_pos(pos int){

}

func Print(){

}

func main(){

}