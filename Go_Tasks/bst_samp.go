package main

import "fmt"
type tree struct{
	left*tree
	val int
	right*tree
}

var root*tree
var enter=1
var foun=1
var level=1
func add(num int){
	add:=&tree{nil,num,nil}
        if(root==nil){
			root=add
		}else{
			var temp*tree=root
			for true{
				if(temp.val>=num){
					if(temp.left==nil){
						temp.left=add
						break
					}else{
						temp=temp.left
					}
				}else{
					if(temp.right==nil){
						temp.right=add
						break
					}else{
						temp=temp.right
					}
				}
			}
		}
}
func traversalpreorder(temp*tree){
        if temp==nil { 
			return
		}
       fmt.Println(temp.val)
	   traversalpreorder(temp.left)
	   traversalpreorder(temp.right)
}
func traversalpostorder(temp*tree){
	if temp==nil { 
		return
	}
	traversalpostorder(temp.left)
	traversalpostorder(temp.right)
    fmt.Println(temp.val)
}
func traversalinorder(temp*tree){
	if temp==nil{ 
		return
	}
	traversalinorder(temp.left)
   fmt.Println(temp.val)
    traversalinorder(temp.right)
 
}
func search(temp*tree,digit int){
	if(temp==nil){
		return 
	}
	if temp.val>root.val && enter==1 {
		enter =0
		level=2
	}
	if(temp.val==digit) {
		foun=0
		fmt.Println("Element Found",level)
		return 
	}
	
	    level++
        search(temp.left,digit)
		search(temp.right,digit)
}
func main(){
	var a=0
	for true{
		foun=1
		enter=1
		level=1
		a=0
		var n int
		fmt.Println("1.Add an element \n2.find an element\n3.InorderTraversal\n4.PostOrderTraversal\n5.PreOrderTraversal\n6.Stop")
         fmt.Scanln(&n)
		switch(n){
		case 1:{
			var digit int
			fmt.Println("Enter the number ")
            fmt.Scanln(&digit)
			add(digit)
		}
	    case 2:{
		    var digit int
		    fmt.Println("Enter a element to find")
		    fmt.Scanln(&digit)
			search(root,digit)
			if(foun==1) {
				
				fmt.Println("Element not found")
				
			}
			
	    }
	    case 3:{
		    traversalinorder(root)
     	}
        case 4:{
         	traversalpostorder(root)
        }
        case 5:{
	        traversalpreorder(root)
        }
	    case 6:{
			a=1 
			break
		}
		}
		if a==1 {
			break
		}
	}
	
}