package main

import "sync"

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	workers := 3

	var wg sync.WaitGroup

	wg.Add(3)

	for i := 0; i < workers; i++ {

		go func() {
			defer wg.Done()
			for j := range jobs {
				result := j * 2
				results <- result
			}
		}()
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}

	for i := 0; i < 10; i++ {
		result := <-results
		println(result)
	}

	close(jobs)

	wg.Wait()

	close(results)

}
