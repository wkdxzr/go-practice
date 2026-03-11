package main

import (
	"fmt"
	"sync"
	"time"
)

// 生产者
func producer(n int, numChan chan int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(numChan)
	}()

	for i := 0; i < n; i++ {
		numChan <- i
		time.Sleep(time.Second)
	}
}

func consumers(numChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numChan {
		fmt.Println(num)
	}
}

func main() {
	n := 10
	var wg sync.WaitGroup
	numChan := make(chan int, n)

	wg.Add(1)
	go producer(n, numChan, &wg)
	wg.Add(1)
	go consumers(numChan, &wg)

	wg.Wait()
}
