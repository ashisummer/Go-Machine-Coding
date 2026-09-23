package main

import (
	"fmt"
	"sync"
)

//print odd even numbers from 1 to 10 using two goroutines and channels

func main() {
	odd := make(chan bool)
	even := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	// Odd goroutine
	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i += 2 {
			<-odd
			fmt.Println(i)
			even <- true
		}
	}()

	// Even goroutine
	go func() {
		defer wg.Done()

		for i := 2; i <= 10; i += 2 {
			<-even
			fmt.Println(i)

			// Don't send after printing 10
			if i < 10 {
				odd <- true
			}
		}
	}()

	// Start with odd
	odd <- true

	wg.Wait()
}
