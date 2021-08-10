package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Square struct {
	length float64
}

func createSquare() Square {
	return Square{
		length: readFloat(),
	}
}

func (s Square) getArea() float64 {

	return s.length * s.length

}

func (s Square) printArea() {

	name := reflect.TypeOf(s).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("area of", shapeName[1], s.getArea())

}
