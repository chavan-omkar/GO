package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {

	sqaureArea := square{
		sideLength: 10.0,
	}

	triangleArea := triangle{
		height: 10.0,
		base:   5.0,
	}

	printArea(sqaureArea)
	printArea(triangleArea)

}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
