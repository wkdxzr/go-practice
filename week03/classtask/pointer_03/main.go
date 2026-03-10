package main

import "fmt"

func doubleValues(nums *[5]int) {
	for i := 0; i < len(nums); i++ {
		nums[i] *= 2
	}
}

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	doubleValues(&nums)
	fmt.Println(nums)
}
