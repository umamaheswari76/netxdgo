package main

import (
	"fmt"
	"sync"
)

var (
	SharedResorce int
	mutex         sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			adding(id)
		}(i)
	//  wg.Wait() 
	}
	wg.Wait() 
	fmt.Println("final val ", SharedResorce)
}

func adding(id int) {
	// mutex.Lock()
	// defer mutex.Unlock()
	// fmt.Println("trying to increment ", id)
	SharedResorce++
	// fmt.Println("incremented shared resorce ", id, " ", SharedResorce)
}

// wg.Add(1) -> adds a goroutine
// wg.Done() -> decrements the goroutine
// mutex.Lock() -> it locks the shared resources from other go routines

//1. wg.Wait() after the for loop will concurrently run the goroutines, wait untill all the go routines completes its execution
//2. wg.Wait() adding inside the for loop, after the go routine,will sequntially run the goroutines

// The presence of fmt.Println statements after the mutex.Lock() code 
// can affect the scheduling and timing of goroutines. 

// The output from multiple goroutines may be interleaved in 
// unexpected ways, making it harder to observe the issues.

// Because of this, even if the mutex is present or not, the final  val printed is 100,
// means the print statements provided some timings inbetween the execution of goroutines,
// so the shared resources didnt affected

// if there is no print statement and also if there is no mutex, then there is no timings
// for the goroutines to run, so there may be a chance of overlapping of goroutines,
// this may results in unexpected final val.(try this with the code nad check)


