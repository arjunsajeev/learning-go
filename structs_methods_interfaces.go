package main

import "math"

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Perimeter() float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
