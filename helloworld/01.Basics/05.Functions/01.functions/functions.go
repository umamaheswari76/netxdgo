package main

import (
	"fmt"
	//"math"
)

//go functions return multiple data
func price(prevprice, currentprice float32) (float32, float32){
	change := currentprice - prevprice
	percentchange := (change / prevprice) * 100
	return change, percentchange
}

func main(){
	prevprice := float32(50.0)
	currentprice := float32(100.0)

	// _ used for skipping values
	change, percentchange := price(prevprice, currentprice)

	if percentchange<0{
		fmt.Printf("Price is decereased by $%f which is of %f%% of the prev price\n ",change, percentchange)
	}else{
		fmt.Printf("Price is increased by $%f which is of %f%% of the prev price\n",change,percentchange )
	}
}