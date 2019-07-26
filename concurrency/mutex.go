package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

    var mtx sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
            mtx.Lock() // this portion is now locked by a mutex, preventing race conditions
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched() // will allow another runtime to run
			v++
			counter = v
            mtx.Unlock() // unlocks the portion for read/write
			wg.Done() // done incrementing
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait() // don't exit the program until everything is done!
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
