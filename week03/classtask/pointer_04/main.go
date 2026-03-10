package main

import "fmt"

func divide(a, b *int, result *float64) error {
	if *b == 0 {
		return fmt.Errorf("除数不能为0")
	}
	*result = float64(*a) / float64(*b)
	return nil
}

func main() {
	/*a1, b1 := 1, 0
	var result1 float64
	err1 := divide(&a1, &b1, &result1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(result1)*/

	a2, b2 := 5, 2
	var result2 float64
	err2 := divide(&a2, &b2, &result2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(result2)
}
