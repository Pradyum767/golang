package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(10 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 1
		for {
			select {
			case <-ticker.C:
				fmt.Printf("Ticker times up for another %v iteration\n", i)
			case <-timer.C:
				fmt.Println("Timer times up")
				return
			}
			i++
		}
	}()
	fmt.Println("Started timer", time.Now())
	wg.Wait()
}
