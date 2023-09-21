//range is an operator

package main

import "fmt"

func main(){
	primeNumbers := []int{2,3,5,11,13,17,19,23,29}
	//here index and number are var
	sum:=0
	for index, number := range primeNumbers{
		fmt.Printf("PrimeNumber(%d) = %d\n", index+1, number)
	}

//output
// 	PrimeNumber(1) = 2
// PrimeNumber(2) = 3
// PrimeNumber(3) = 5
// PrimeNumber(4) = 11
// PrimeNumber(5) = 13
// PrimeNumber(6) = 17
// PrimeNumber(7) = 19
// PrimeNumber(8) = 23
// PrimeNumber(9) = 29


//using blank identifier

	for _, number := range primeNumbers{
		sum += x
	}
	fmt.Printf("Total sum = %d",sum)
}