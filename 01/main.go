package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

type triangle struct {
	base   float64
	height float64
}

type square struct {
	lenght float64
}

type circle struct {
	radius float64
}

func main() {

	s := square{
		lenght: 10,
	}

	t := triangle{
		base:   5,
		height: 10,
	}

	c := circle{
		radius: 5,
	}

	s.printArea()
	t.printArea()
	c.printArea()

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) printArea() {
	name := reflect.TypeOf(t).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("area of", shapeName[1], " is", t.getArea())
}

func (s square) getArea() float64 {
	return s.lenght * s.lenght
}

func (s square) printArea() {
	name := reflect.TypeOf(s).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("area of", shapeName[1], s.getArea())
}

func (c circle) getArea() float64 {

	return math.Pi * c.radius * c.radius

}

func (c circle) printArea() {
	name := reflect.TypeOf(c).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("area of", shapeName[1], c.getArea())
}
