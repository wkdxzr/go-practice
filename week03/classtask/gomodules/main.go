package main

import (
	"fmt"

	"github.com/wkdxzr/go-practice/week03/classtask/gomodules/mathutil"
)

func main() {
	sum := mathutil.Add(1, 2)
	fmt.Println(sum)

	product := mathutil.Multiply(3, 4)
	fmt.Println(product)
}
