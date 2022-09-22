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

// Let's define a method, which lives outside. But is a in-struct
// or-object method. Just defined outside.
func (rect *Rectangle) RectangleVolume(rectangle Rectangle) float64 {
	return rectangle.xCoordinate * rectangle.yCoordinate * rectangle.zCoordinate
}

// Inheriting from `Rectangle` struct/object
type Square struct {
  Rectangle // An anonymous field
}

func (sqr *Square) SquareArea(square Square) float64 {
  return square.xCoordinate * square.xCoordinate
}

func (sqr *Square) SquareVolume(square Square) float64 {
  return square.xCoordinate * square.xCoordinate * square.xCoordinate
}
func main() {
	rect1 := Rectangle{10, 20, 30}
	fmt.Printf("Area of rectangle 1: %v\n", RectangleArea(rect1))
	fmt.Printf("Volume of rectangle 1: %v\n", rect1.RectangleVolume(rect1))
  fmt.Println()

  sqr1 := Square{}
  sqr1.xCoordinate = 10
  fmt.Printf("Area of square 1: %v\n", sqr1.SquareArea(sqr1))
  fmt.Printf("Volume of square 1: %v\n", sqr1.SquareVolume(sqr1))
}
