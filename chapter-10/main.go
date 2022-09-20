package main

import "fmt"

type Rectangle struct {
	xCoordinate float64
	yCoordinate float64
  zCoordinate float64
}

func RectangleArea(rectangle Rectangle) float64 {
	return rectangle.xCoordinate * rectangle.yCoordinate
}

func RectangleVolume(rectangle Rectangle) float64 {
  return rectangle.xCoordinate * rectangle.yCoordinate * rectangle.zCoordinate
}

func main() {
	rect1 := Rectangle{10, 20, 30}
	fmt.Printf("Area of rectangle 1: %v\n", RectangleArea(rect1))
	fmt.Printf("Volume of rectangle 1: %v\n", RectangleVolume(rect1))
}
