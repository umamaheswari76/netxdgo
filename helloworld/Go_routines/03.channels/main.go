package main


import "fmt"
func routine(ch chan int){
	fmt.Println(100+ <-ch)

}

//main go routine sending the data to the other go routine
//data - channel - main
//in below code there are two go routines, one is main go routine and the other is go routine

func main(){
	channel_demo := make(chan int) //here the channel  is duplex
	//channel is accepting int
	//channel_demo <- 234  ---- deadlock(before sending the data, we must have the receiver)
	
	//this go routine cts as a receiver
	go routine(channel_demo) //calling and wait for receiving the data
	channel_demo <- 234 
		//sending the data to the channel
	
		fmt.Println("value of channel ", channel_demo)
		fmt.Println("Tyoe of channel %T", channel_demo)

}