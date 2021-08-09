package main

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
