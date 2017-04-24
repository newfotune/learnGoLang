package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 10     //same as w.Circle.Point.X = 10
	w.Y = 10     //same as w.Circle.Point.Y = 10
	w.Radius = 5 //same as w.Circle.Radius = 5
	w.Spokes = 20

	w.print()

	var c Circle
	c.X = 20
	c.Y = 20
	c.Radius = 30
	//c.print()
}

func (w Wheel) print() {
	fmt.Printf("X: %d, Y: %d, Radius: %d, Spokes: %d", w.X, w.Y, w.Radius, w.Spokes)
}
