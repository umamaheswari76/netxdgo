package main

import "fmt"

type Node struct{
	Data int
	Right *Node
	Left *Node
}

var Root *Node = nil

func Create(element int) *Node{
	temp := new(Node)
	temp.Data = element
	temp.Left = nil
	temp.Right = nil
	return temp
}

func Add(element int){
	if Root==nil{
		Root = Create(element)
	}else{
		if 
	}
}

func Delete(){

}


func main(){

}