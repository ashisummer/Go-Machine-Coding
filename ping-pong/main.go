package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ping := make(chan bool)
	pong := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Time's up!")
				return
			case <-ping:
				fmt.Println("ping")
				time.Sleep(200 * time.Millisecond)

				select {
				case <-ctx.Done():
					fmt.Println("Time's up!")
					return
				case pong <- true:
				}
			}
		}

	}()

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Time's up!")
				return
			case <-pong:
				fmt.Println("pong")
				time.Sleep(200 * time.Millisecond)
				select {
				case <-ctx.Done():
					fmt.Println("Time's up!")
					return
				case ping <- true:
				}
			}
		}

	}()

	select {
	case <-ctx.Done():
		fmt.Println("Time's up!")
	case ping <- true:
	}

	wg.Wait()
}
