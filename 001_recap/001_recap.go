package main

import (
	"fmt"
	"math"
)

// create a type square
// create a type circle
// attach a method to each that calculates area and returns it
// create a type shape which defines an interface as anything which has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

type square struct{
	length float64
	width float64
}
type circle struct {
	radius float64
}
func (s square)area() float64{
	return s.length * s.width
}
func (c circle)area() float64 {
	return c.radius * c.radius * math.Pi
}
type shape interface {
	area() float64
}
func info(sh shape){
	fmt.Println("The area is:", sh.area())
}
func main() {
	sq := square{
		50,
		20,
	}
	ci := circle{
		42,
	}
	info(sq)
	info(ci)

}