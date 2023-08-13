When we want things to be done concurrently in Golang, we use Goroutines.

Concurrency: 
        Dealing multiple jobs by some time slice(dealing with multiple things at once (does not need to be done at the same time) with some time schedule)

Parallelism:
         When things are being done at the same time it’s called parallelism.

Goroutines
       -> Goroutines are a way of doing tasks concurrently in golang means, go routines allows us to create and run multiple methods or functions concurrently in the same address space inexpensively. 
       -> Creating and destructing goroutines are very cheap compared to threads, and they are scheduled over OS threads.

Goroutines execution:
        -> Unlike normal functions, the control does not wait for goroutine to finish executing.
        -> The control immediately returns to the next line of the code after a Goroutine call
        -> The main function should be run for anyother goroutines to run(main func also a goroutine)
        -> if main terminates, then other goroutine will not run, they will be terminated

Just know about, preemptive(have time slice) and non-preemptive scheduling

Scheduling of Goroutines
        -> Goroutines are cooperatively scheduled. 
        -> In Cooperative scheduling, there is no concept of scheduler time slice, ie. the goroutines periodically provide the control when they are idle or logically blocked in order to run multiple goroutines
        ->The switch between Goroutines happen only at well defined points — when an explicit call is made to the Goruntime scheduler. 


Ref 
        -> https://riteeksrivastava.medium.com/a-complete-journey-with-goroutines-8472630c7f5c
        -> https://blog.nindalf.com/posts/how-goroutines-work/

channels - go with the order
        -> https://blog.nindalf.com/posts/how-goroutines-work/
        -> https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/
        -> https://www.meetgor.com/golang-go-routines/