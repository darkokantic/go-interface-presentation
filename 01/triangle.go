package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Triangle struct {
	base   float64
	height float64
}

func createTriangle() Triangle {
	return Triangle{
		base:   readFloat(),
		height: readFloat(),
	}
}

func (t Triangle) getArea() float64 {

	return 0.5 * t.base * t.height

}

func (t Triangle) printArea() {

	name := reflect.TypeOf(t).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("area of", shapeName[1], " is", t.getArea())

}
