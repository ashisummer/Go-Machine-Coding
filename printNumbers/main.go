package main

//print 1 to 10 with routines and waitgroup

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("speaking", i)
		}()
	}
	wg.Wait()
}
