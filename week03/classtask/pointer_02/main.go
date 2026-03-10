package main

import "fmt"

type Point struct {
	x, y int
}

func scalePoint(p *Point, factor float64) {
	p.x = int(float64(p.x) * factor)
	p.y = int(float64(p.y) * factor)
}

func main() {
	p := &Point{x: 1, y: 2}
	fmt.Println(*p)
	scalePoint(p, 2.0)
	fmt.Println(*p)
}
