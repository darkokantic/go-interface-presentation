package main

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
