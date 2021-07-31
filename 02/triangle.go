package main

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
