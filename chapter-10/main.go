package main

import "fmt"

type Square struct {
	xCoordinate float64
}

func SquareArea(Square Square) float64 {
	return Square.xCoordinate * Square.xCoordinate
}

// In-struct/in-object methods defined outside
func (Square *Square) SquareVolume() float64 {
	return Square.xCoordinate * Square.xCoordinate * Square.xCoordinate
}

type Rectangle struct {
	// Inherited from the `Square` structure
	Square

	yCoordinate float64
	zCoordinate float64
}

func (Rectangle *Rectangle) RectangleArea() float64 {
	return Rectangle.xCoordinate * Rectangle.yCoordinate
}

func (Rectangle *Rectangle) RectangleVolume() float64 {
	return Rectangle.xCoordinate * Rectangle.yCoordinate * Rectangle.zCoordinate
}

func main() {
	sqr := Square{10}

	fmt.Printf("Area of the square: %v\n", SquareArea(sqr))
	fmt.Printf("Volume of the square: %v\n", sqr.SquareVolume())
	fmt.Println()

	rect := Rectangle{}
	rect.xCoordinate = 10
	rect.yCoordinate = 20
	rect.zCoordinate = 30

	fmt.Printf("Area of the rectangle: %v\n", rect.RectangleArea())
	fmt.Printf("Volume of the rectangle: %v\n", rect.RectangleVolume())

}
