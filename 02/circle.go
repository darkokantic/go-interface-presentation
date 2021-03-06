package main

import (
	"math"
)

type Circle struct {
	radius float64
}

func createCircle() Circle {
	return Circle{
		radius: readFloat(),
	}
}

func (c Circle) getArea() float64 {

	return math.Pi * c.radius * c.radius

}
