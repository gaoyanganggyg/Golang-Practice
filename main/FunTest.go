package main

import (
	"math"
	"fmt"
)

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(scale int) *Point {
	p.X = float64(scale) * p.X
	p.Y = float64(scale) * p.Y
	return p
}

func main() {
	p := Point{X:10, Y:45}
	q := Point{X:24, Y:89}
	res1 := Distance(p, q)
	fmt.Printf("%.2f\n", res1)

	res2 := p.Distance(q)
	fmt.Printf("%.5f\n", res2)
}
