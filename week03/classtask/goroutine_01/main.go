package main

import (
	"fmt"
	"sync"
)

// 写一个函数计算部分和，把结果发送到通道
func calPartSum(s []int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var parSum int
	for _, v := range s {
		parSum += v * v
	}
	resultChan <- parSum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numPartitions := 4

	partLength := len(numbers) / numPartitions
	resultChan := make(chan int, numPartitions)
	var wg sync.WaitGroup

	//循环numPartitions次计算
	for i := 0; i < numPartitions; i++ {
		start := i * partLength
		end := start + partLength
		if i == numPartitions-1 {
			end = len(numbers)
		}
		wg.Add(1)
		go calPartSum(numbers[start:end], resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	sum := 0
	for v := range resultChan {
		sum += v
	}
	fmt.Println(sum)
}
