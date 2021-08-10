package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
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

func (c Circle) printArea() {

	name := reflect.TypeOf(c).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("The area of", shapeName[1], " is ", c.getArea())

}
