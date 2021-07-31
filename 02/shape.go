package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Shape interface {
	getArea() float64
}

func printArea(s Shape) {

	name := reflect.TypeOf(s).String()
	shapeName := strings.Split(name, ".")
	fmt.Println("area of", shapeName[1], s.getArea())

}
