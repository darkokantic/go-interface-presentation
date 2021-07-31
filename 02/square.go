package main

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
