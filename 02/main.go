package main

import (
	"fmt"
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

func main() {

	s := square{
		lenght: 10,
	}

	t := triangle{
		base:   5,
		height: 10,
	}

	s.printArea()
	t.printArea()

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
