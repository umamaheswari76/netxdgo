//error, nil are keywords....eithr one of them must be writtern

package main

import (
	"fmt"
	"errors"
	"math"
)
func price(prevprice float, currentprice float64) (float64, float64, error){

	if prevprice == 0{
		err := errors.New("Previous price cannot be zero")
		return 0,0,err
	}
	change := currentprice - prevprice
	percentchange := (change / prevprice) * 100
	return change, percentchange, nil
}

func main(){
	prevprice := 16.0
	currentprice :=20.0

	// _ used for skipping values
	change, percentchange, err := price(prevprice, currentprice)

	//satndard approach t handle error in go (like try catch in other lang)
	if err != nil{
		fmt.Println("Sorry! There was an error: ", err)
	}else {
		if change<0{
			fmt.Printf("Price is decereased by $ %.2f which of %.2f%% of the prev price\n ",math.Abs(change), math.Abs(percentchange) )
		}else{
			fmt.Printf("Price is increased by $ %.2f which is %.2f%% of the prev price\n",change, percentchange )
		}
		
	}
}