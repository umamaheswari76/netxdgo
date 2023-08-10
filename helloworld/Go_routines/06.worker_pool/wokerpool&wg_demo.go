// this program shows the combination of wait group in worker pool

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct{
	id int
	randomno int
}

type Results struct{
	job job
	sumofdigits int
}

var jobs = make(chan Job,10)
var results = make(chan Result, 10)

func digits(number int) int{
	sum:=0
	no := number
	for no!=0{
		digit := no%10
		sum+=digit
		no/=10
	}
	//time.Sleep(2*time.Second)
	return sum
}

func worker(wg *sync WaitGroup){
	for job:=range jobs{
		output := Result(job, digits(job.randomno))
		results<-output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int){
	var wg sync.WaitGroup
	for i:=0;<noOfWorkers;i++{
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int){
	for i:=0;i<noOfJobs;i++{
		randomno := rand.Intn(999)
		job:=Job{i,rndomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool){
	for result := range results{
		fmt.Println("Job id %d, input random no %d, sum of digits %d\n",reslt.job.id, result.job.randomno, result.)
	}
	done <-true
}

func main(){
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go reslt(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(),"seconds")
}
