package main

import "fmt"

type Node struct{
	Data int
	Prev *Node
	Next *Node
}

var Head *Node = nil
var Tail *Node = nil
var count  = 0

func Create(element int) *Node{
	n := new(Node)
	count++
	n.Data = element
	n.Prev = nil
	n.Next = nil
	return n
}

func Add(element int){
	fmt.Println("Choose any one option(number), Add element at\n	1.Begin\n	2.End\n   3.Position")
	var inp int
	fmt.Scanln(&inp)

	if inp==1{
		Add_begin(element)
	}else if inp==2{
		Add_end(element)
	}else if inp==3{
		fmt.Println("Enter the position to add element ")
		var pos int
		fmt.Scanln(&pos)
		Add_at(element, pos)
	}
}

func Add_begin(element int){
	if Head==nil{
		Head = Create(element)
		Tail = Head
	}else{
		temp := Create(element)
		temp.Next = Head
		Head.Prev = temp.Next
		Head = temp
		temp = nil
	}
}

func Add_end(element int){
	 if Head==nil{
		Head = Create(element)
		Tail = Head
	}else{
	    temp := Create(element)
	    Tail.Next = temp
	    temp.Prev = Tail
	    Tail = temp
	    temp = nil
	}
}

func Add_at(element, pos int){
    temp := new(Node)
    it := Head
    
    //3   2
    for i:=1;i<pos-1;i++{
        it = it.Next
    }
    if it==nil{
        fmt.Println("Enter valid position to add element for option 3 ")
    }else{
        temp.Next = it.Next
        temp.Prev = it
        it.Next = temp
        temp.Next.Prev = temp
    }

}

func Del(){
	fmt.Println("Choose any one option(number), Delete element at\n	1.Begin\n	2.End\n	  3.Position ")
	var inp int
	fmt.Scanln(&inp)

	if inp==1{
		Del_begin()
	}else if inp==2{
		Del_end()
	}else if inp==3{
		fmt.Println("Enter the position to delete element for option 3 ")
		var pos int
		fmt.Scanln(&pos)
		Del_pos(pos)
	}
}

func Del_begin(){
    temp := Head
	Head = temp.Next
	Head.Prev = nil
	temp = nil
}

func Del_end(){
    temp := Tail
    temp.Prev = Tail
    Tail.Next = nil
    temp = nil

}

func Del_pos(pos int){
    it := Head
    for i:=1;i<pos-1;i++{
        it = it.Next
    }
    
    if it==nil{
        fmt.Println("Enter valid position to delete element")
    }else{
        count--
        temp := it.Next
        it.Next = temp.Next
        temp.Next.Prev = it
        temp = nil
    }
}

func Print(){
    for i:=Head; i!=nil; {
		temp := i
		fmt.Println(temp.Data)
		i = temp.Next
	}
}

func main(){
    Add(1)
	Add(2)
	// Add(3)
	// Add(4)
	// Add(5)
	// Add(6)
	// Add(7)
	// Add(8)
	// Add(9)
	// Del()
	Print()


    

}