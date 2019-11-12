// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle
package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s circle) area() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}

func info(s shape) {
	fmt.Printf("Area of %T:\t%v\n", s, s.area())
}

func main() {
	smallSquare := square{length: 10}

	largeCircle := circle{radius: 28.5}

	info(smallSquare)
	info(largeCircle)
}
