package main

import (
	"fmt"

	"github.com/ashisummer/Go-Machine-Coding/cache-go/cache"
)

// implement a thread safe cache with a map and a mutex lock
func main() {

	cache := cache.NewCache()

	cache.Set("key1", 1)

	value, ok := cache.Get("key1")
	if ok {
		println("Value:", value)
	}

	cache.Set("key2", 2)

	value, ok = cache.Get("key2")
	if ok {
		println("Value:", value)
	}

	value, ok = cache.Get("key3")
	if ok {
		println("Value:", value)
	} else {
		println("Key not found")
	}

	fmt.Println("All data in cache:")
	data := cache.GetAll()

	fmt.Println("Data in cache:", data)

	// for _, d := range data {
	// 	println("Key:", d.Key, "Value:", d.Value)
	// }

}
