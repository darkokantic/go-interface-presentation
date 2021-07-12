package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
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

func (c circle) printArea() {

	name := reflect.TypeOf(c).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("The area of", shapeName[1], " is ", c.getArea())

}
