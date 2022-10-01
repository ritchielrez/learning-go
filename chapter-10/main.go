package main

import "fmt"

type Shape interface {
	volume() float64
}

type Square struct {
	xCoordinate float64
}

func area(Square Square) float64 {
	return Square.xCoordinate * Square.xCoordinate
}

// In-struct/in-object methods defined outside
func (Square *Square) volume() float64 {
	return Square.xCoordinate * Square.xCoordinate * Square.xCoordinate
}

type Rectangle struct {
	// Inherited from the `Square` structure
	Square

	yCoordinate float64
	zCoordinate float64
}

func (Rectangle *Rectangle) area() float64 {
	return Rectangle.xCoordinate * Rectangle.yCoordinate
}

func (Rectangle *Rectangle) volume() float64 {
	return Rectangle.xCoordinate * Rectangle.yCoordinate * Rectangle.zCoordinate
}

func totalVolume(shapes ...Shape) float64 {
	var totalVol float64

	for _, shape := range shapes {
		totalVol += shape.volume()
	}
	return totalVol
}

func main() {
	sqr := Square{10}

	fmt.Printf("Area of the square: %v\n", area(sqr))
	fmt.Printf("Volume of the square: %v\n", sqr.volume())
	fmt.Println()

	rect := Rectangle{}
	rect.xCoordinate = 10
	rect.yCoordinate = 20
	rect.zCoordinate = 30

	fmt.Printf("Area of the rectangle: %v\n", rect.area())
	fmt.Printf("Volume of the rectangle: %v\n", rect.volume())
	fmt.Println()

	fmt.Printf("Total volume of both shapes is: %v\n", totalVolume(&sqr, &rect))
}
