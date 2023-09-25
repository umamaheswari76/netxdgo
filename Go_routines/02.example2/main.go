package main


import (
	"fmt"
	"time"
)

//cant predict the execution og go routies
//when it ga the available slots in cpu then theywill execte
//the design pattern tht we willl write in our prog 
func main(){
	// for i:=0;i<5;i++{
	// 	//time.Sleep(1*time.Second)
	// 	go func(){
	// 		//time.Sleep(10*time.Second)
	// 		fmt.Println(i)
			
	// 	}()
	// 	time.Sleep(1*time.Second)
	// }

	for i:=0;i<20;i++{
		go func(x int){
			fmt.Println(x)
			
		}(i)
		time.Sleep(1*time.Second)
	}
	//time.Sleep(1*time.Second)
}