package main

import (
	"fmt"
	"reflect"
	"strings"
)

type square struct {
	lenght float64
}

func createSquare() square {
	return square{
		lenght: readFloat(),
	}
}

func (s square) getArea() float64 {

	return s.lenght * s.lenght

}

func (s square) printArea() {

	name := reflect.TypeOf(s).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("area of", shapeName[1], s.getArea())

}
