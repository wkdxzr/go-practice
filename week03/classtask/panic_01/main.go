package main

import "fmt"

func divide(a, b int) float64 {
	if b == 0 {
		panic("除数不能为0")
	}
	return float64(a) / float64(b)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("程序恢复！捕获到 panic 错误: %v\n", r)
		}
	}()

	result := divide(1, 0)
	fmt.Println(result)
}
