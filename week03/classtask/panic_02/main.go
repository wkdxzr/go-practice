package main

import "fmt"

func accessArray(intArray [10]int, index int) int {
	if index >= len(intArray) || index < 0 {
		panic("索引超出数组范围")
	}
	return intArray[index]
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("程序恢复！捕获到 panic 错误:%v\n", r)
		}
	}()

	nums := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	value := accessArray(nums, 10)
	fmt.Println(value)
}
