package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
	fmt.Println(name, "done")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go increment(fmt.Sprintf("Goroutine %d", i), &wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
