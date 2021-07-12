package main

import (
	"fmt"
)

func main() {
	fmt.Println("Calculate surface area of circle, square or triangle.")
	fmt.Println("------------------------------------------------------")
	fmt.Println("Choose option: 1, 2 or 3")
	fmt.Println("1 - Circle")
	fmt.Println("2 - Square")
	fmt.Println("3 - Triangle")
	fmt.Println("Option: ")

	option := readInt()
	switch option {
	case 1:
		fmt.Println("What is the radius of the circle")
		c := ceateCircle()
		c.printArea()

	case 2:
		fmt.Println("What is the lenght of the square?")
		s := createSquare()
		s.printArea()
	case 3:
		fmt.Println("What are the base and height of the triangle?")
		t := createTriangle()
		t.printArea()
	default:
		fmt.Println("Incorect option, please enter 1, 2 or 3!")

	}

}

func readInt() int {
	var i int
	fmt.Println("Enter an integer value: ")

	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println(err)
	}
	return i
}

func readFloat() float64 {

	var f float64

	fmt.Println("Enter a float value: ")
	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		fmt.Println(err)
	}

	return f
}
