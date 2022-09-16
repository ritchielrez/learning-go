package main

import "fmt"

type Rectangle struct {
	xCoordinate float64
	yCoordinate float64
}

func RectangleArea(rectangle Rectangle) float64 {
	return rectangle.xCoordinate * rectangle.yCoordinate
}

func main() {
	rect1 := Rectangle{10, 20}
	fmt.Printf("Area of rectangle 1: %v\n", RectangleArea(rect1))

	rect2 := Rectangle{10, 20}
	fmt.Printf("Area of rectangle 1: %v\n", RectangleArea(rect2))
}
