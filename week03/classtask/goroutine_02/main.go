package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// 统计文件中特定单词出现的次数，并把结果发送到通道
func wordSum(filePath string, word string, countChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		count += strings.Count(line, word)
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("扫描", filePath, "错误:", err)
		return
	}

	countChan <- count
}

func main() {
	var word string
	fmt.Println("请输入你想要统计的特定的单词:")
	fmt.Scanln(&word)

	paths := []string{
		"week03/classtask/goroutine_02/test01.txt",
		"week03/classtask/goroutine_02/test02.txt",
		"week03/classtask/goroutine_02/test03.txt",
	}
	var wg sync.WaitGroup
	countChan := make(chan int, len(paths))

	for _, v := range paths {
		wg.Add(1)
		go wordSum(v, word, countChan, &wg)
	}

	go func() {
		wg.Wait()
		close(countChan)
	}()

	sum := 0
	for v := range countChan {
		sum += v
	}
	fmt.Println(word, "出现总次数:", sum)
}
