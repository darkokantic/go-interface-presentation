package main

import (
	"math"
)

type circle struct {
	radius float64
}

func ceateCircle() circle {
	return circle{
		radius: readFloat(),
	}
}

func (c circle) getArea() float64 {

	return math.Pi * c.radius * c.radius

}
