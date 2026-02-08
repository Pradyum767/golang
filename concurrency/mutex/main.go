package main

import (
	"fmt"
	"sync"
	"time"
)

// Using sync.Mutex to ensure that only one goroutine can access the critical section of code that updates the counter value at a time
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func updateData() {
	counter := &Counter{}
	var wg sync.WaitGroup
	for range 1000 {

		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Counter value: ", counter.value)
}

// using sync.RWMutex to allow multiple readers to access the data concurrently while ensuring that only one writer can modify the data at a time
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func (s *SafeMap) Get(key string) string {
	s.mu.RLock() // Multiple readers can hold this at once
	defer s.mu.RUnlock()
	return s.data[key]
}

func (s *SafeMap) Set(key, value string) {
	s.mu.Lock() // Only one writer can hold this
	defer s.mu.Unlock()
	s.data[key] = value
}

func readAndWriteMapData() {
	// Initialize the struct and the internal map
	store := &SafeMap{
		data: make(map[string]string),
	}

	var wg sync.WaitGroup

	// 1. Start a "Writer" goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Writer: Setting session_1...")
		store.Set("session_1", "active")
	}()

	// 2. Start multiple "Reader" goroutines
	// These can run concurrently with each other
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Give the writer a tiny head start for the sake of the example
			time.Sleep(10 * time.Millisecond)

			val := store.Get("session_1")
			fmt.Printf("Reader %d: Got value: %s\n", id, val)
		}(i)
	}

	wg.Wait()
	fmt.Println("All operations completed safely.")
}

func main() {
	//check behaviur of mutex when multiple goroutines are trying to update the same data concurrently and ensure that the final value is correct
	for range 10 {
		updateData()
	}

	//check the behaviour of RWMutex when multiple goroutines are trying to read and write data concurrently and ensure that the readers can access the data while the writer is updating
	readAndWriteMapData()
}
