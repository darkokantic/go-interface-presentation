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

func createTriangle() triangle{
	return triangle{
		base:   readFloat(),
		height: readFloat(),
	}
}

func (t triangle) getArea() float64 {

	return 0.5 * t.base * t.height

}

func (t triangle) printArea() {

	name := reflect.TypeOf(t).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("area of", shapeName[1], " is", t.getArea())

}
