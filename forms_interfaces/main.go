package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct{
	sideLength float64
}
type triangle struct{
	height float64
	base float64
}


func main() {
	sq := square{10}
	tr := triangle{height: 10, base: 10}

	// fmt.Println(sq.getArea())
	// fmt.Println(tr.getArea())

	printArea(tr)
	printArea(sq)
}



func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

