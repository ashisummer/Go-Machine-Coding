package main

import (
	"fmt"
	"sync"
)

// print 1 to 100 with 10 routines , workerpool
func main() {

	workers := 3

	var wg sync.WaitGroup

	ch := make(chan int)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for c := range ch {
				fmt.Println("worker", i, "speaking", c)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()

}
