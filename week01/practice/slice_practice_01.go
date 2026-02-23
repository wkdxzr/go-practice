package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s)

	s3to7 := s[2:7]
	fmt.Println(s3to7)

	s = append(s, 11, 12, 13)
	fmt.Println(s)

	s = append(s[:4], s[5:]...)
	fmt.Println(s)

	for i := range s {
		s[i] *= 2
	}
	fmt.Println(s)
	fmt.Println(cap(s))
}
