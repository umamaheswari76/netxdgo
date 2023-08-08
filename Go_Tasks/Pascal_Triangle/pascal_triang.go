package main

import "fmt"

func Pascal(num int){
	for row:=1;row<=num;row++{
		for space :=1;space<=num-row;space++{
			fmt.Printf(" ")
		}
		temp:=1
		for i:=1;i<=row;i++ {
			fmt.Printf("%d ",temp)
			temp = temp*(row-i)/i
		}
		fmt.Println()
	}
}

func main(){
	var num int
	fmt.Scanln(&num)
	Pascal(num)
}