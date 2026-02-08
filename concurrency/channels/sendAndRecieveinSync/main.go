package main

import (
	"fmt"
	"strconv"
	"sync"
)

func printHello(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		fmt.Print("Hello ")
		ch <- i
		<-ch
	}
	close(ch)
}

func printWorld(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println("World " + strconv.Itoa(i))
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go printHello(ch, &wg) //initiating go routine to print Hello and send data to channel and receive data from channel
	go printWorld(ch, &wg) //initiating go routine to print World and receive data from channel and send data to channel
	wg.Wait()
}
