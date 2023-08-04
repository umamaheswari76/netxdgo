package main

import (
	"fmt"
	"errors"
)

// func factorial(num int) (int, error){
// 	if(num<0){
// 		err:=errors.New("Number should not be negative")
// 		return 0,err
// 	}
	
// 	res := 1
// 	for temp:=num;temp>=1;temp--{
// 		res*=temp
// 	}
// 	return res,nil
// }

func factorial(num int) (int, error){
	if num<0{
		err := errors.New("Number should not be negative")
		return 0, err
	}
	
	if num==0{
		return 1,nil
	}
	temp,err := num*factorial(num-1)

	return temp,nil
	
}


func main(){
	var inp int
	fmt.Scanln(&inp)
	fact, err := factorial(inp)
	
	if err!=nil{
		fmt.Println("ERROR\n")
	}else{
		fmt.Printf("Factorial of %d is: %d ", inp, fact)
	}


}



