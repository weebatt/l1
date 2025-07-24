package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint() *Point {
	return &Point{
		x: 0,
		y: 0,
	}
}

func (p *Point) Distance(op *Point) float64 {
	return math.Sqrt(math.Pow(p.x-op.x, 2) + math.Pow(p.y-op.y, 2))
}

func (p *Point) Set(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) Get() (float64, float64) {
	return p.x, p.y
}

func main() {
	point1 := NewPoint()
	point2 := NewPoint()

	var x, y float64

	_, _ = fmt.Scanln(&x, &y)
	point1.Set(x, y)

	_, _ = fmt.Scanln(&x, &y)
	point2.Set(x, y)

	fmt.Println(point1.Distance(point2))
}
