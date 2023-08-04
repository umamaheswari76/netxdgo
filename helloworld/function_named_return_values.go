package main

import (
	"fmt"
	"math"
)


//go functions return mltiple data
//here the return types are named
func price(prevprice, currentprice float64) (change, percentchange float64){
	change = currentprice - prevprice
	percentchange = (change / prevprice) * 100
	return change, percentchange
}


func main(){
	prevprice := 0.0
	currentprice :=100000.0


	change, percentchange := price(prevprice, currentprice)

	if change<0{
		fmt.Printf("Price is decereased by $ %.2f which of %.2f%% of the prev price\n ",math.Abs(change), math.Abs(percentchange) )
	}else{
		fmt.Printf("Price is increased by $ %.2f which is %.2f%% of the prev price\n",change, percentchange )
	}
}