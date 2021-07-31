package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Square struct {
	lenght float64
}

func createSquare() Square {
	return Square{
		lenght: readFloat(),
	}
}

func (s Square) getArea() float64 {

	return s.lenght * s.lenght

}

func (s Square) printArea() {

	name := reflect.TypeOf(s).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("area of", shapeName[1], s.getArea())

}
